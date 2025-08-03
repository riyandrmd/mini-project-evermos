package dto

type CreateTrxItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Qty       int  `json:"qty" validate:"required,gt=0"`
}

type CreateTrxRequest struct {
	IDAlamat         uint            `json:"id_alamat" validate:"required"`
	MetodePembayaran string          `json:"metode_pembayaran" validate:"required"`
	Items            []CreateTrxItem `json:"items" validate:"required,dive"`
}
