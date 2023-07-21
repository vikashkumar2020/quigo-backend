package models

import "gorm.io/gorm"

type Rider struct {
	gorm.Model
	Name       string `json:"name" gorm:"size:255;not null;"`
	Email      string `json:"email" gorm:"size:255;not null;"`
	Password   string `json:"password" gorm:"size:255;not null;"`
	Phone      string `json:"phone" gorm:"size:255;not null;"`
	IsVerified bool   `json:"is_verified" gorm:"size:255;not null;"`
	Address    string `json:"address" gorm:"size:255;not null;"`
}
