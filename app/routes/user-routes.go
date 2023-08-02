package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/controllers/user"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/me", user.GetMe())
	// router.GET("/rides",user.GetRides())
	// router.GET("/wallet",user.GetWallet())
	// router.GET("/rides/:id",user.GetRide())
	// router.GET("/wallet/transactions",user.GetTransactions())
}
