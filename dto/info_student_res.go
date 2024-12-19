package dto

type InfoStudentRes struct {
	NimDinus  string `json:"nim_dinus"`
	TaMasuk   string `json:"ta_masuk"`
	Prodi     string `json:"prodi"`
	AkdmStat  string `json:"akdm_stat"`
	DateReg   string `json:"date_reg"`
	SppStatus string `json:"spp_status"`
	Kelas     string `json:"kelas"`
}
