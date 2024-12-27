package dto

type KrsOffersProdiResponse struct {
	Id              int    `gorm:"column:id" json:"id"`
	TahunAjaran     string `gorm:"column:tahun_ajaran"`
	KodeMataKuliah  string `gorm:"column:kode_mata_kuliah"`
	Kelompok        string `gorm:"column:kelompok"`
	NamaMataKuliah  string `gorm:"column:nama_mata_kuliah"`
	JumlahSKS       int    `gorm:"column:jumlah_sks"`
	Hari            string `gorm:"column:hari"`
	JamMulai        string `gorm:"column:jam_mulai"` // Perbaikan di sini, sebelumnya ada typo `:=`
	JamSelesai      string `gorm:"column:jam_selesai"`
	Ruang           string `gorm:"column:ruang"`
	JnsJam          string `gorm:"column:jns_jam"`
	StatusPemilihan string `gorm:"column:status_pemilihan"`
	StatusKrs       string `gorm:"column:status_krs"`
}
