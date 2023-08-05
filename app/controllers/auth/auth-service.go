package auth

import (
	// "fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"github.com/vikashkumar2020/quigo-backend/app/models"
	services "github.com/vikashkumar2020/quigo-backend/app/services"
	"github.com/vikashkumar2020/quigo-backend/config"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
	utils "github.com/vikashkumar2020/quigo-backend/utils"
)

// Register User as rider or driver
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload *models.SignUpInput

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		if payload.Password != payload.PasswordConfirm {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
			return
		}

		hashedPassword, err := utils.HashPassword(payload.Password)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			return
		}

		now := time.Now()
		newUser := models.User{
			Name:       payload.Name,
			Email:      strings.ToLower(payload.Email),
			Phone:      payload.Phone,
			Password:   hashedPassword,
			Role:       payload.Role,
			Verified:   false,
			Address:    "",
			PrivateKey: "",
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		db := pgdatabase.GetDBInstance().GetDB()
		result := db.Create(&newUser)

		if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email or Phone already exists"})
			return
		} else if result.Error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
			return
		}

		code := randstr.String(20)

		verification_code := utils.Encode(code)
		// Update User in Database
		newUser.VerificationCode = verification_code
		db.Save(newUser)

		var firstName = newUser.Name

		if strings.Contains(firstName, " ") {
			firstName = strings.Split(firstName, " ")[0]
		}

		emailData := utils.EmailData{
			URL:       "http://localhost:3000/auth" + "/verifysuccess/" + code,
			FirstName: firstName,
			Subject:   "Your account verification code",
		}

		utils.SendEmail(&newUser, &emailData, "verificationCode.html")

		message := "We sent an email with a verification code to " + newUser.Email
		ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
	}
}

// VerifyEmail User as rider or driver
func VerifyEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Params.ByName("verificationCode")
		verification_code := utils.Encode(code)

		var updatedUser models.User
		db := pgdatabase.GetDBInstance().GetDB()
		result := db.First(&updatedUser, "verification_code = ?", verification_code)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid verification code or user doesn't exists"})
			return
		}

		if updatedUser.Verified {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already verified"})
			return
		}

		address, privateKey := services.GetUnusedPrivateKey(db, updatedUser.Email)

		updatedUser.Verified = true
		updatedUser.Address = address
		updatedUser.PrivateKey = privateKey
		db.Save(&updatedUser)

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email verified successfully","data": updatedUser.Email})
	}
}

// Login User as rider or driver
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload *models.SignInInput

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User
		db := pgdatabase.GetDBInstance().GetDB()
		result := db.First(&user, "email = ?", strings.ToLower(payload.Email))
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
			return
		}

		if !user.Verified {
			ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Please verify your email"})
			return
		}

		if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Password"})
			return
		}

		config := config.NewJWTConfig()

		// Generate Tokens
		token, err := utils.GenerateToken(30*24*time.Hour, user.ID, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.SetCookie("token", token, 60*24*30, "/", "localhost", false, true)

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
	}
}

// ForgotPassword User as rider or driver
func ForgotPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload *models.ForgotPasswordInput

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		message := "You will receive a reset email if user with that email exist"

		var user models.User
		db := pgdatabase.GetDBInstance().GetDB()
		result := db.First(&user, "email = ?", strings.ToLower(payload.Email))
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
			return
		}

		if !user.Verified {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Account not verified"})
			return
		}

		// Generate Verification Code
		resetToken := randstr.String(20)

		passwordResetToken := utils.Encode(resetToken)
		user.PasswordResetToken = passwordResetToken
		user.PasswordResetAt = time.Now().Add(time.Minute * 15)
		db.Save(&user)

		var firstName = user.Name

		if strings.Contains(firstName, " ") {
			firstName = strings.Split(firstName, " ")[0]
		}

		// ? Send Email
		emailData := utils.EmailData{
			URL:       "http://localhost:3000/auth" + "/resetpassword/" + resetToken,
			FirstName: firstName,
			Subject:   "Your password reset token (valid for 10min)",
		}

		utils.SendEmail(&user, &emailData, "resetPassword.html")

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	}
}

// ResetPassword User as rider or driver
func ResetPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload *models.ResetPasswordInput
		resetToken := ctx.Params.ByName("resetToken")

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		if payload.Password != payload.PasswordConfirm {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
			return
		}

		hashedPassword, _ := utils.HashPassword(payload.Password)

		passwordResetToken := utils.Encode(resetToken)

		var updatedUser models.User
		db := pgdatabase.GetDBInstance().GetDB()
		result := db.First(&updatedUser, "password_reset_token = ? AND password_reset_at > ?", passwordResetToken, time.Now())
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The reset token is invalid or has expired"})
			return
		}

		updatedUser.Password = hashedPassword
		updatedUser.PasswordResetToken = ""
		db.Save(&updatedUser)

		ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password data updated successfully"})

	}
}

// Logout User as rider or driver
func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully Logged out"})
	}
}
