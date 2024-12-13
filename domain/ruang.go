package domain

type Ruang struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Nama         string `gorm:"not null"`
	Nama2        string `gorm:"default:'-'"`
	IDJenisMakul int    `gorm:"default:null"`
	IDFakultas   string `gorm:"default:null"`
	Kapasitas    int    `gorm:"default:0"`
	KapUjian     int    `gorm:"default:0"`
	Status       int16  `gorm:"default:1;comment:1: buka 0: tutup 2: hapus"`
	Luas         string `gorm:"default:'0';comment:meter persegi"`
	Kondisi      string `gorm:"default:null"`
	Jumlah       int    `gorm:"default:null"`
}

func (r *Ruang) TableName() string {
	return "ruang"
}
