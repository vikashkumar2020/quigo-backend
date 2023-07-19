package models

import "gorm.io/gorm"

type Rider struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	IsVerified bool   `json:"is_verified"`
	Address    string `json:"address"`
}
