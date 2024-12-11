package dto

type LoginReq struct {
	NimDinus string `json:"nim_dinus" validate:"required"`
	PassMhs  string `json:"pass_mhs" validate:"required"`
}
