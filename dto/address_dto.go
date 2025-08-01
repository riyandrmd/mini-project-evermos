package dto

type CreateAddressRequest struct {
	Label     string `json:"label" validate:"required"`
	Recipient string `json:"recipient" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Address   string `json:"address" validate:"required"`
}

type AddressResponse struct {
	ID        uint   `json:"id"`
	Label     string `json:"label"`
	Recipient string `json:"recipient"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
