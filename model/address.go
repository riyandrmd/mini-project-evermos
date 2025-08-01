package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Label     string `json:"label"`
	Recipient string `json:"recipient"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	User      *User  `json:"-"`
}
