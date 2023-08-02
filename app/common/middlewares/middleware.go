package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/models"
	"github.com/vikashkumar2020/quigo-backend/config"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
	"github.com/vikashkumar2020/quigo-backend/utils"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config := config.NewJWTConfig()
		sub, err := utils.ValidateToken(token, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User
		db := pgdatabase.GetDBInstance().GetDB()
		result := db.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Set("role", user.Role)
		ctx.Next()
	}
}

func DriverCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		if role != "driver" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not a driver and not authorized to access this resource"})
			return
		}
		ctx.Next()
	}
}

func RiderCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		if role != "rider" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not a rider and not authorized to access this resource"})
			return
		}
		ctx.Next()
	}
}

func RegisterUserMiddleware(router *gin.RouterGroup) {
	router.Use(DeserializeUser())
}

func RegisterDriverMiddleware(router *gin.RouterGroup) {
	router.Use(DeserializeUser())
	router.Use(DriverCheck())

}

func RegisterRiderMiddleware(router *gin.RouterGroup) {
	router.Use(DeserializeUser())
	router.Use(RiderCheck())
}
