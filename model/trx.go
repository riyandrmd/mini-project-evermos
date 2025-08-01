package model

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	UserID      uint        `json:"user_id"`
	AddressID   uint        `json:"address_id"`
	Method      string      `json:"method"`
	InvoiceCode string      `json:"invoice_code"`
	TotalPrice  float64     `json:"total_price"`
	User        User        `json:"-" gorm:"foreignKey:UserID"`
	Address     *Address    `json:"-" gorm:"foreignKey:AddressID"`
	DetailTrx   []DetailTrx `json:"-" gorm:"foreignKey:TrxID"`
}
