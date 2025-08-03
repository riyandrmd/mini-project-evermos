package model

import "time"

type Toko struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	NamaToko  string `json:"nama_toko"`
	Deskripsi string `json:"deskripsi"`
	Foto      string `json:"foto"`

	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
