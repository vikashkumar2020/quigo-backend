package models

import (
	"time"

	"github.com/google/uuid"
)

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Role            string `json:"role" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

