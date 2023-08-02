package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	Email      string
	Balance    int64     `gorm:"not null;default:100"`
	Address    string    `gorm:"not null"`
	PrivateKey string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

type WalletProvider struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

type PaymentRequest struct {
	Amount int64 `json:"amount"`
}
