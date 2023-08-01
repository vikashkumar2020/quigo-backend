package rides

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/models"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
)

func CreateRide() gin.HandlerFunc{
	// create ride
	return func(ctx *gin.Context) {

		var payload *models.RideRequest

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(400, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user := ctx.MustGet("currentUser").(models.User)

		ride := models.Rides{
			RiderEmail: user.Email,
			From: payload.From,
			To: payload.To,
			Price: payload.Amount,
			RideStatus: "requested",
			PaymentStatus: "pending",
			RiderAddress: user.Address,
			RiderPrivateKey: user.PrivateKey,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		db := pgdatabase.GetDBInstance().GetDB()
		result := db.Create(&ride)

		if result.Error != nil {
			ctx.JSON(400, gin.H{"status": "error", "message": "Something bad happened"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "create ride",
			"ride_details": ride,
		})
	}
}

func GetRiderRideDetails() gin.HandlerFunc{
	return func(c *gin.Context) {

		id := c.Param("id")

		db := pgdatabase.GetDBInstance().GetDB()
		var ride models.Rides

		result := db.Where("id = ?", id).First(&ride)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Ride not found"})
			return
		}

		if ride.RideStatus == "requested" {
			c.JSON(200, gin.H{"status": "pending", "message": "Ride is still pending"})
			return
		}

		var driverDetails models.User
		result = db.Where("email = ?", ride.DriverEmail).First(&driverDetails)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Driver not found"})
			return
		}


		riderRideDetails := models.RiderRideDetails{}
		riderRideDetails.From = ride.From
		riderRideDetails.To = ride.To
		riderRideDetails.Price = ride.Price
		riderRideDetails.RideStatus = ride.RideStatus
		riderRideDetails.PaymentStatus = ride.PaymentStatus
		riderRideDetails.DriverName = driverDetails.Name
		riderRideDetails.DriverNumer = driverDetails.Phone

		c.JSON(200, gin.H{
			"message": "ride details",
			"ride_details": riderRideDetails,
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