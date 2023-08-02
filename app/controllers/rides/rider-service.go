package rides

import (
	"math/big"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/models"
	"github.com/vikashkumar2020/quigo-backend/app/services"
	"github.com/vikashkumar2020/quigo-backend/infra/eth"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
)

func CreateRide() gin.HandlerFunc {
	// create ride
	return func(ctx *gin.Context) {

		var payload *models.RideRequest

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(400, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user := ctx.MustGet("currentUser").(models.User)

		ride := models.Rides{
			RiderEmail:      user.Email,
			Origin:          payload.Origin,
			Destination:     payload.Destination,
			Price:           payload.Amount,
			RideStatus:      "requested",
			PaymentStatus:   "pending",
			RiderAddress:    user.Address,
			RiderPrivateKey: user.PrivateKey,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		db := pgdatabase.GetDBInstance().GetDB()
		result := db.Create(&ride)

		if result.Error != nil {
			ctx.JSON(400, gin.H{"status": "error", "message": "Something bad happened"})
			return
		}

		ctx.JSON(200, gin.H{
			"message":      "create ride",
			"ride_details": ride,
		})
	}
}

func GetRiderRideDetails() gin.HandlerFunc {
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
			c.JSON(200, gin.H{"status": "requested", "message": "Ride is still pending"})
			return
		}

		var driverDetails models.User
		result = db.Where("email = ?", ride.DriverEmail).First(&driverDetails)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Driver not found"})
			return
		}

		riderRideDetails := models.RiderRideDetails{}
		riderRideDetails.Origin = ride.Origin
		riderRideDetails.Destination = ride.Destination
		riderRideDetails.Price = ride.Price
		riderRideDetails.RideStatus = ride.RideStatus
		riderRideDetails.PaymentStatus = ride.PaymentStatus
		riderRideDetails.DriverName = driverDetails.Name
		riderRideDetails.DriverNumer = driverDetails.Phone

		c.JSON(200, gin.H{
			"message":      "ride details",
			"ride_details": riderRideDetails,
		})
	}
}

func Payment() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		db := pgdatabase.GetDBInstance().GetDB()

		var ride models.Rides

		result := db.Where("id = ?", id).First(&ride)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Ride not found"})
			return
		}

		if ride.PaymentStatus == "paid" {
			c.JSON(200, gin.H{"status": "paid", "message": "Payment already done"})
			return
		}

		// Get rider wallet details

		var riderWallet models.Wallet
		result = db.Where("email = ?", ride.RiderEmail).First(&riderWallet)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Rider wallet not found"})
			return
		}

		// Get driver wallet details

		var driverWallet models.Wallet
		result = db.Where("email = ?", ride.DriverEmail).First(&driverWallet)

		if result.Error != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Driver wallet not found"})
			return
		}

		// eth client

		ethClient := eth.GetEthClient();

		riderConn, riderAuth := services.GetConnection(ethClient,riderWallet.PrivateKey)
		driverConn, driverAuth := services.GetConnection(ethClient,driverWallet.PrivateKey)

		// Deduct amount from rider wallet

		price, err := strconv.ParseInt(ride.Price, 10, 64)
		if err != nil {
			// handle the error if necessary
			return
		}

		riderConn.Withdrawl(riderAuth,big.NewInt(price))
		driverConn.Deposite(driverAuth,big.NewInt(price))
	
		c.JSON(200, gin.H{
			"message": "payment Successfull",
			"ride_details": ride,
		})
	}
}

func Atoi(s string) {
	panic("unimplemented")
}
