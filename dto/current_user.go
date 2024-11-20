package dto

type CurrentUser struct {
	NIM      string `json:"nim" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
