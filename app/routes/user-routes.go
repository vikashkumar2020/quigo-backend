package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/controllers/user"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/me",user.GetMe())
}