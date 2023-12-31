package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/controllers/auth"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", auth.Login())
	router.POST("/register", auth.Register())
	router.POST("/forgotpassword", auth.ForgotPassword())
	router.PATCH("/resetpassword/:resetToken", auth.ResetPassword())
	router.GET("/logout", auth.Logout())
	router.GET("/verifysuccess/:verificationCode", auth.VerifyEmail())
}
