package dto

type StatusKrsRes struct {
	Validate    string `json:"validate,omitempty"`
	TahunAjaran string `json:"tahun_ajaran,omitempty"`
	Dipaketkan  string `json:"dipaketkan,omitempty"`
	TahunMasuk  string `json:"tahun_masuk,omitempty"`
	Sks         int    `json:"sks,omitempty"`
	Ips         string `json:"ips,omitempty"`
}
