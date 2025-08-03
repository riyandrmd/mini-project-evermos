package model

import "time"

type Trx struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	UserID           uint   `json:"user_id"`
	KodeInvoice      string `json:"kode_invoice"`
	HargaTotal       int    `json:"harga_total"`
	MetodePembayaran string `json:"metode_pembayaran"`

	AlamatID uint   `json:"alamat_id"`
	Alamat   Alamat `json:"alamat" gorm:"foreignKey:AlamatID"`

	User   User        `json:"user" gorm:"foreignKey:UserID"`
	Detail []TrxDetail `json:"detail" gorm:"foreignKey:TrxID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
