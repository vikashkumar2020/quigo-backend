package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	Name          string  `json:"name" gorm:"size:255;not null;"`
	Email         string  `json:"email" gorm:"size:255;not null;"`
	Password      string  `json:"password" gorm:"size:255;not null;"`
	Phone         string  `json:"phone" gorm:"size:255;not null;"`
	IsVerified    bool    `json:"is_verified" gorm:"size:255;not null;"`
	VechileNumber string  `json:"vechile_number"`
	VechileType   string  `json:"vechile_type"`
	Rating        float32 `json:"rating" `
	Address       string  `json:"address" gorm:"size:255;not null;"`
}
