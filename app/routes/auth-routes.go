package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/controllers/auth"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login())
	router.POST("/register", controllers.Register())
	router.POST("/forgot-password", controllers.ForgotPassword())
}