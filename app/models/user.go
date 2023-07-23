package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name               string    `gorm:"type:varchar(255);not null"`
	Email              string    `gorm:"uniqueIndex;not null"`
	Phone              string    `gorm:"uniqueIndex;not null"`
	Password           string    `gorm:"not null"`
	Role               string    `gorm:"type:varchar(255);not null"`
	Verified           bool      `gorm:"not null"`
	VerificationCode   string
	PasswordResetToken string
	PasswordResetAt    time.Time
	Address            string    `gorm:"not null"`
	PrivateKey         string    `gorm:"not null"`
	CreatedAt          time.Time `gorm:"not null"`
	UpdatedAt          time.Time `gorm:"not null"`
}
