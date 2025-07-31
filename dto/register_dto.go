package dto

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`

	Phone   string `json:"phone"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
