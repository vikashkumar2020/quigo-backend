package rides

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/models"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
)

func GetRides() gin.HandlerFunc {
	return func(c *gin.Context) {

		var availableRides []models.RideDetail
		db := pgdatabase.GetDBInstance().GetDB()

		if err := db.Model(&models.Rides{}).
			Select("id, ride_status, price, origin, destination, payment_status").
			Where("ride_status = ?", "requested").
			Find(&availableRides).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve pending rides"})
			return
		}

		c.JSON(200, gin.H{
			"message": "all the pending rides",
			"rides":   availableRides,
		})
	}
}

func GetDriverRideDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		db := pgdatabase.GetDBInstance().GetDB()

		var ride models.Rides

		result := db.Where("id = ?", id).First(&ride)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Ride not found"})
			return
		}

		if ride.RideStatus != "requested" {
			c.JSON(400, gin.H{"status": "error", "message": "Cannot View details of this ride"})
			return
		}

		var riderDetails models.User
		result = db.Where("email = ?", ride.RiderEmail).First(&riderDetails)
		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Rider not found"})
			return
		}

		driverRideDeatils := models.DriverRideDetails{
			RiderName:     riderDetails.Name,
			RiderNumer:    riderDetails.Phone,
			Origin:        ride.Origin,
			Destination:   ride.Destination,
			Price:         ride.Price,
			RideStatus:    ride.RideStatus,
			Duration: 	ride.Duration,
			Distance: 	ride.Distance,
			PaymentStatus: ride.PaymentStatus,
		}

		c.JSON(200, gin.H{
			"message":      "get ride details",
			"ride_details": driverRideDeatils,
		})
	}
}

func AcceptRide() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		db := pgdatabase.GetDBInstance().GetDB()

		user := c.MustGet("currentUser").(models.User)

		var ride models.Rides

		result := db.Where("id = ?", id).First(&ride)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Ride not found"})
			return
		}

		if ride.RideStatus != "requested" {
			c.JSON(400, gin.H{"status": "error", "message": "Cannot accept this ride"})
			return
		}

		ride.RideStatus = "accepted"
		ride.UpdatedAt = time.Now()
		ride.DriverEmail = user.Email
		ride.DriverAddress = user.Address
		ride.DriverPrivateKey = user.PrivateKey

		db.Save(&ride)
		c.JSON(200, gin.H{
			"message":      "accepted the ride",
			"ride_details": ride,
		})
	}
}

func Start() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		db := pgdatabase.GetDBInstance().GetDB()

		var ride models.Rides

		result := db.Where("id = ?", id).First(&ride)

		if result.Error != nil {

			c.JSON(400, gin.H{"status": "error", "message": "Ride not found"})
			return
		}

		if ride.RideStatus != "accepted" {
			c.JSON(400, gin.H{"status": "error", "message": "Cannot start this ride"})
			return
		}

		ride.RideStatus = "started"
		ride.UpdatedAt = time.Now()
		ride.Departure = time.Now()

		db.Save(&ride)

		c.JSON(200, gin.H{
			"message":      "started ride",
			"ride_details": ride,
		})
	}
}

func CompleteRide() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")

		db := pgdatabase.GetDBInstance().GetDB()

		var ride models.Rides

		result := db.Where("id = ?", id).First(&ride)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Ride not found"})
			return
		}

		if ride.RideStatus != "started" {
			c.JSON(400, gin.H{"status": "error", "message": "Cannot complete this ride"})
			return
		}

		if ride.PaymentStatus == "pending" {
			c.JSON(400, gin.H{"status": "error", "message": "Cannot complete this ride, Payment in pending"})
			return
		}

		ride.RideStatus = "completed"
		ride.UpdatedAt = time.Now()
		ride.Arrival = time.Now()

		db.Save(&ride)

		c.JSON(200, gin.H{
			"message":      "start ride",
			"ride_details": ride,
		})
	}
}

func CancelRide() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		db := pgdatabase.GetDBInstance().GetDB()

		var ride models.Rides

		result := db.Where("id = ?", id).First(&ride)

		if result.Error != nil {

			c.JSON(400, gin.H{"status": "error", "message": "Ride not found"})
			return
		}

		if ride.RideStatus != "requested" {
			c.JSON(400, gin.H{"status": "error", "message": "Cannot cancel this ride"})
			return
		}

		if ride.RideStatus == "completed" {
			c.JSON(400, gin.H{"status": "error", "message": "Cannot cancel this ride, already completed"})
			return
		}

		ride.RideStatus = "cancelled"
		ride.UpdatedAt = time.Now()

		c.JSON(200, gin.H{
			"message":      "cancel ride",
			"ride_details": ride,
		})
	}
}
