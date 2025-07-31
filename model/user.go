package model

import "gorm.io/gorm"

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100)" json:"name"`
	Email     string         `gorm:"type:varchar(100);unique" json:"email"`
	Password  string         `gorm:"type:varchar(255)" json:"password"`
	IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
