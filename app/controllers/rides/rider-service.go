package rides

import (
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
			Duration:        payload.Duration,
			Distance:        payload.Distance,
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
		riderRideDetails.Duration = ride.Duration
		riderRideDetails.Distance = ride.Distance

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

		ethClient := eth.GetEthClient()

		price, err := strconv.ParseInt(ride.Price, 10, 64)
		if err != nil {
			// handle the error if necessary
			return
		}

		Conn := services.GetConnection(ethClient,"28eee86e1836f578030dc7b77d5a90bbaa460f95074c7864a0948cc3a778af79")

		driverAuth := services.GetAccountAuth(ethClient, driverWallet.PrivateKey)
		res, err := Conn.Deposite(driverAuth, big.NewInt(price))
		if err != nil {
			// handle error
			log.Println(err)
		}
		log.Println(res)

		reply, err := Conn.Balance(&bind.CallOpts{})
		if err != nil {
			// handle error
			log.Println(err)
		}
		log.Println(reply.Int64())
		driverWallet.Balance = driverWallet.Balance + price

		riderAuth := services.GetAccountAuth(ethClient, riderWallet.PrivateKey)
		res, err = Conn.Withdrawl(riderAuth, big.NewInt(price))
		if err != nil {
			// handle error
			log.Println(err)
		}
		log.Println(res)
		reply, err = Conn.Balance(&bind.CallOpts{})
		if err != nil {
			// handle error
			log.Println(err)
		}
		riderWallet.Balance = riderWallet.Balance - price

		db.Save(&riderWallet)
		db.Save(&driverWallet)

		ride.PaymentStatus = "paid"

		db.Save(&ride)
		c.JSON(200, gin.H{
			"message":"payment Successfull",
			"ride_details": ride,
		})
	}
}

func Atoi(s string) {
	panic("unimplemented")
}
