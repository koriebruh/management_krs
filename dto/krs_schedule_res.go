package dto

type KrsScheduleRes struct {
	TA         int    `json:"ta"`
	Prodi      string `json:"prodi"`
	TglMulai   string `json:"tgl_mulai"`
	TglSelesai string `json:"tgl_selesai"`
}
