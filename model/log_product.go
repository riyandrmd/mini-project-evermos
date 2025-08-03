package model

type LogProduk struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	NamaProduk string `json:"nama_produk"`
	Slug       string `json:"slug"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`
	TokoID     uint   `json:"toko_id"`
	CategoryID uint   `json:"category_id"`
}
