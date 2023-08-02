package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	Email      string
	Balance    int64  `gorm:"not null;default:100"`
	PrivateKey string `gorm:"uniqueIndex;not null;primary_key"`
	PublicKey  string
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

type PaymentRequest struct {
	Amount int64 `json:"amount"`
}
