package dto

type CreateTrxItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Qty       int  `json:"qty" validate:"required,min=1"`
}

type CreateTrxRequest struct {
	Method    string          `json:"method" validate:"required"`
	AddressID uint            `json:"address_id" validate:"required"`
	Items     []CreateTrxItem `json:"items" validate:"required,dive"`
}
