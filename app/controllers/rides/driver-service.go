package rides

import (
	"github.com/gin-gonic/gin"
)

func GetRides() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get rides",
		})
	}
}

func GetDriverRideDetails() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get ride details",
		})
	}
}

func AcceptRide() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "accept ride",
		})
	}
}

func Start() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "start ride",
		})
	}
}

func CompleteRide() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "complete ride",
		})
	}
}

func CancelRide() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "cancel ride",
		})
	}
}