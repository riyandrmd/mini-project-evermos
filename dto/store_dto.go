package dto

import "time"

type UpdateStoreRequest struct {
	Name string `json:"name" validate:"required"`
}

type StoreResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
