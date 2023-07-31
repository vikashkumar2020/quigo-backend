package models

import (
	"time"

	"github.com/google/uuid"
)

type Rides struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	RiderEmail       string    `gorm:"uniqueIndex;not null;primary_key"`
	DriverEmail 	 string    `gorm:"uniqueIndex"`
	From             string    `gorm:"not null"`
	To               string    `gorm:"not null"`
	Price            string
	Seats            string
	Departure        time.Time
	Arrival          time.Time
	RideStatus       string 	`gorm:"not null"`
	PaymentStatus    string 	`gorm:"not null"`
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
	Address   string `json:"address"`
	PrivateKey string `json:"private_key"`

}
