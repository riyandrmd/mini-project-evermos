package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
	Phone    string `json:"phone" gorm:"uniqueIndex"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	About    string `json:"about"`
	Job      string `json:"job"`
	Province string `json:"province"`
	City     string `json:"city"`
	Store    *Store `json:"store"`
}
