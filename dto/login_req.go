package dto

type LoginReq struct {
	NIM      string `json:"nim" validate:"required"`
	Password string `json:"password" validate:"required"`
}
