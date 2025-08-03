package model

type TrxDetail struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	TrxID       uint `json:"trx_id"`
	LogProdukID uint `json:"log_produk_id"`
	TokoID      uint `json:"toko_id"`
	Qty         int  `json:"qty"`
	HargaTotal  int  `json:"harga_total"`

	LogProduk LogProduk `json:"log_produk" gorm:"foreignKey:LogProdukID"`
}
