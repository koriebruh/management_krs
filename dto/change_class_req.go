package dto

type ChangeClassReq struct {
	Kelas int `json:"kelas" validate:"required,min=0,max=3"`
}
