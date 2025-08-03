package dto

type CreateProductRequest struct {
	NamaProduk string `json:"nama_produk" validate:"required"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int    `json:"harga" validate:"required,gt=0"`
	Stok       int    `json:"stok" validate:"required"`
	CategoryID uint   `json:"id_category" validate:"required"`
}
