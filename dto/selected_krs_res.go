package dto

type SelectedKrs struct {
	NamaMatkul   string `json:"nama_matkul"`
	NamaMatkulEN string `json:"nama_matkul_en"`
	Tipe         string `json:"tipe"`
	Semester     int    `json:"semester"`
	JenisMatkul  string `json:"jenis_matkul"`
	Hari1        string `json:"hari_1"`
	Hari2        string `json:"hari_2"`
	Hari3        string `json:"hari_3"`
}
