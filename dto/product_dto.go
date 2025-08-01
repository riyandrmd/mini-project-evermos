package dto

type CreateProductRequest struct {
	Name          string  `json:"name" validate:"required"`
	Slug          string  `json:"slug" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	PriceReseller float64 `json:"price_reseller" validate:"required"`
	PriceCustomer float64 `json:"price_customer" validate:"required"`
	CategoryID    uint    `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	Description   string  `json:"description"`
	PriceReseller float64 `json:"price_reseller"`
	PriceCustomer float64 `json:"price_customer"`
	CategoryID    uint    `json:"category_id"`
}

// type ProductResponse struct {
// 	ID         uint   `json:"id"`
// 	Name       string `json:"name"`
// 	Price      int    `json:"price"`
// 	Stock      int    `json:"stock"`
// 	CategoryID uint   `json:"category_id"`
// 	StoreID    uint   `json:"store_id"`
// }
