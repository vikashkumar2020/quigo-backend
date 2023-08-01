package rides

import "github.com/gin-gonic/gin"

func CreateRide() gin.HandlerFunc{
	// create ride
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create ride",
		})
	}
}

func GetRiderRideDetails() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get ride details",
		})
	}
}

func Payment() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "payment",
		})
	}
}