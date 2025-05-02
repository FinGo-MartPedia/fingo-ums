package requests

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	UserAgent string `json:"user_agent" validate:"required"`
	IPAddress string `json:"ip_address" validate:"required"`
}

func (l LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
