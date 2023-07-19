package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	Phone         string  `json:"phone"`
	IsVerified    bool    `json:"is_verified"`
	VechileNumber string  `json:"vechile_number"`
	VechileType   string  `json:"vechile_type"`
	Rating        float32 `json:"rating"`
	Address       string  `json:"address"`
}
