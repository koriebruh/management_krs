package dto

type RegisterReq struct {
	NimDinus string `json:"nim_dinus" validate:"required"`
	TAMasuk  int    `json:"ta_masuk" validate:"required"`
	Prodi    string `json:"prodi" validate:"required"`
	PassMhs  string `json:"pass_mhs" validate:"required"`
	Kelas    int    `json:"kelas" validate:"required,min=1,max=3"`
	AkdmStat string `json:"akdm_stat" validate:"required,oneof=1 2 3 4 5 6 7"`
}
