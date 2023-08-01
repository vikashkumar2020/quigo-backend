package models

import (
	"time"

	"github.com/google/uuid"
)

type Rides struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	RiderEmail       string    `gorm:"uniqueIndex;not null;primary_key"`
	DriverEmail      string    `gorm:"uniqueIndex"`
	From             string    `gorm:"not null"`
	To               string    `gorm:"not null"`
	Price            string
	Seats            string
	Departure        time.Time
	Arrival          time.Time
	RideStatus       string `gorm:"not null"`
	PaymentStatus    string `gorm:"not null"`
	VerificationCode string
	Role             string    `gorm:"type:varchar(255);not null"`
	Address          string    `gorm:"not null"`
	PrivateKey       string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

type RideRequest struct {
	RiderEmail string `json:"email"`
	From       string `json:"from"`
	To         string `json:"to"`
	FromLng    string `json:"from_lng"`
	FromLat    string `json:"from_lat"`
	ToLng      string `json:"to_lng"`
	ToLat      string `json:"to_lat"`
	Amnt       string `json:"amnt"`
}

type RiderRideDetails struct {
	DriverName  string `json:"driver_name"`
	DriverNumer string `json:"driver_number"`
	From        string `json:"from"`
	To          string `json:"to"`
	Price       string `json:"price"`
	RideStatus  string `json:"ride_status"`
	PaymentStatus     string `json:"payment_status"`
}

type RideDetail struct{
	RideStatus  string `json:"ride_status"`
	Price 	 string `json:"price"`
	From 	  string `json:"from"`
	To 		  string `json:"to"`
	PaymentStatus     string `json:"payment_status"`
}

type DriverRideDetails struct {
	RiderName string `json:"rider_name"`
	RiderNumer string `json:"rider_number"`
	From      string `json:"from"`
	To        string `json:"to"`
	Price     string `json:"price"`
	RideStatus  string `json:"ride_status"`
	PaymentStatus     string `json:"payment_status"`
}


