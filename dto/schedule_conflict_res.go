package dto

type ScheduleConflictRes struct {
	TahunAjaran    string `gorm:"column:tahun_ajaran" json:"tahun_ajaran"`
	Kelompok       string `gorm:"column:kelompok" json:"kelompok"`
	NamaMataKuliah string `gorm:"column:nama_mata_kuliah" json:"nama_mata_kuliah"`
	JumlahSKS      int    `gorm:"column:jumlah_sks" json:"jumlah_sks"`
	Hari           string `gorm:"column:hari" json:"hari"`
	JamMulai       string `gorm:"column:jam_mulai" json:"jam_mulai"`
	JamSelesai     string `gorm:"column:jam_selesai" json:"jam_selesai"`
	Ruang          string `gorm:"column:ruang" json:"ruang"`
	StatusBentrok  string `gorm:"column:status_bentrok" json:"status_bentrok"`
	KeteranganSlot string `gorm:"column:keterangan_slot" json:"keterangan_slot"`
}
