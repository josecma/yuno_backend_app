package dtos

type CreateUserRequestDTO struct {
	Username    string `json:"username" validate:"required,min=3"`
	Password    string `json:"password" validate:"required,min=8"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
