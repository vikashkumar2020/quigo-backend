package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/controllers/rides"
)

func RegisterRiderRoutes(router *gin.RouterGroup) {
	// rider ride routes
	router.POST("/ride", rides.CreateRide()) // request for a new ride from rider
	router.GET("/rides/:id", rides.GetRiderRideDetails()) // get all availble rides
	router.POST("/rides/:id/payment", rides.Payment()) // pay for ride

}

func RegisterDriverRoutes(router *gin.RouterGroup) {
	// driver ride routes
	router.GET("/rides", rides.GetRides()) // get all availble rides
	router.GET("/rides/:id", rides.GetDriverRideDetails()) // get ride details
	router.PUT("/rides/:id/accept", rides.AcceptRide()) // accept ride
	router.PUT("/rides/:id/start", rides.Start()) // start ride	
	router.GET("/rides/:id/complete", rides.CompleteRide()) // complete ride
	router.GET("/rides/:id/cancel", rides.CancelRide()) 	// cancel ride
}