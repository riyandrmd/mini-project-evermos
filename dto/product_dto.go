package dto

type CreateProductRequest struct {
	Name       string `json:"name" validate:"required"`
	Price      int    `json:"price" validate:"required,gt=0"`
	Stock      int    `json:"stock" validate:"required,gte=0"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

type ProductResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"category_id"`
	StoreID    uint   `json:"store_id"`
}
