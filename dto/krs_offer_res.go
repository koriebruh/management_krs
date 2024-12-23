package dto

type KrsOfferRes struct {
	Id          int    `gorm:"column:id" json:"id"`
	TahunAjaran int    `gorm:"column:tahun_ajaran" json:"tahun_ajaran"`
	Kelompok    string `gorm:"column:kelompok" json:"kelompok"`
	Matakuliah  string `gorm:"column:nama_mata_kuliah" json:"matakuliah"`
	Sks         int    `gorm:"column:jumlah_sks" json:"sks"`
	Hari        string `gorm:"column:hari" json:"hari"`
	JamMulai    string `gorm:"column:jam_mulai" json:"jam_mulai"`
	JamSelesai  string `gorm:"column:jam_selesai" json:"jam_selesai"`
	Ruang       string `gorm:"column:ruang" json:"ruang"`
}
