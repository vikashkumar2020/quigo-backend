package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	Email                 string    `gorm:"uniqueIndex;not null;primary_key"`
	Amount                string    `gorm:"not null"`
	PaymentTime           time.Time `gorm:"not null"`
	SenderWalletAddress   string    `gorm:"not null"`
	ReceiverWalletAddress string    `gorm:"not null"`
	CreatedAt             time.Time `gorm:"not null"`
	UpdatedAt             time.Time `gorm:"not null"`
}
