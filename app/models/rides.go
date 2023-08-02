package models

import (
	"time"

	"github.com/google/uuid"
)

type Rides struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	RiderEmail       string    `gorm:"uniqueIndex;not null;primary_key"`
	DriverEmail      string    `gorm:"uniqueIndex"`
	Origin           string    `gorm:"not null"`
	Destination      string    `gorm:"not null"`
	Price            string
	Departure        time.Time
	Arrival          time.Time
	RideStatus       string `gorm:"not null"`
	PaymentStatus    string `gorm:"not null"`
	VerificationCode string
	RiderAddress     string
	RiderPrivateKey  string
	DriverAddress    string
	DriverPrivateKey string
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

type RideRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	FromLng     string `json:"from_lng"`
	FromLat     string `json:"from_lat"`
	ToLng       string `json:"to_lng"`
	ToLat       string `json:"to_lat"`
	Amount      string `json:"amount"`
}

type RiderRideDetails struct {
	DriverName    string `json:"driver_name"`
	DriverNumer   string `json:"driver_number"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Price         string `json:"price"`
	RideStatus    string `json:"ride_status"`
	PaymentStatus string `json:"payment_status"`
}

type RideDetail struct {
	ID            string `json:"id"`
	RideStatus    string `json:"ride_status"`
	Price         string `json:"price"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	PaymentStatus string `json:"payment_status"`
}

type DriverRideDetails struct {
	RiderName     string `json:"rider_name"`
	RiderNumer    string `json:"rider_number"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Price         string `json:"price"`
	RideStatus    string `json:"ride_status"`
	PaymentStatus string `json:"payment_status"`
}
