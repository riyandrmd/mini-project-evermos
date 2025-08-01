package model

import "gorm.io/gorm"

type DetailTrx struct {
	gorm.Model
	TrxID        uint        `json:"trx_id"`
	LogProductID uint        `json:"log_product_id"`
	StoreID      uint        `json:"store_id"`
	Qty          int         `json:"qty"`
	TotalPrice   float64     `json:"total_price"`
	Trx          *Trx        `json:"-" gorm:"foreignKey:TrxID"`
	LogProduct   *LogProduct `json:"-" gorm:"foreignKey:LogProductID"`
	Store        *Store      `json:"-" gorm:"foreignKey:StoreID"`
}
