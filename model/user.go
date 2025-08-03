package model

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Nama         string `json:"nama"`
	KataSandi    string `json:"-"` // disembunyikan di response
	Notelp       string `gorm:"unique" json:"notelp"`
	Email        string `gorm:"unique" json:"email"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
	IsAdmin      bool   `json:"is_admin" gorm:"default:false"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Toko Toko `json:"toko" gorm:"foreignKey:UserID"`
}
