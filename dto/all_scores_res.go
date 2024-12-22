package dto

type AllScoresRes struct {
	KodeMatkul  string `gorm:"column:kode_matkul" json:"kode_matkul"`
	MataKuliah  string `gorm:"column:matakuliah" json:"mata_kuliah"`
	Sks         int    `gorm:"column:sks" json:"sks"`
	Category    string `gorm:"column:category" json:"category"`
	JenisMatkul string `gorm:"column:jenis_matkul" json:"jenis_matkul"`
	Nilai       string `gorm:"column:nilai" json:"nilai"`
}
