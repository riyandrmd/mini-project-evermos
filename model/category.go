package model

type Category struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	NamaCategory string `json:"nama_category"`
}
