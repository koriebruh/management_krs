package dto

type KrsLogRes struct {
	IdRec          int    `gorm:"column:id_krs" json:"id_rec"`
	NimDinus       string `gorm:"column:nim_dinus" json:"nim_dinus"`
	KodeMataKuliah string `gorm:"column:kdmk" json:"kode_mata_kuliah"`
	Aksi           int8   `gorm:"column:aksi" json:"aksi"`
	IdJadwal       int    `gorm:"column:id_jadwal" json:"id_jadwal"`
	IpAddress      int    `gorm:"column:ip_addr" json:"ip_address"`
	LastUpdate     string `gorm:"column:last_update" json:"last_update"`
}
