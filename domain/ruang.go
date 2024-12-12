package domain

type Ruang struct {
	ID           uint8   `gorm:"primaryKey;autoIncrement"`
	Nama         string  `gorm:"size:250;not null"`
	Nama2        string  `gorm:"size:250;default:'-'"`
	IDJenisMakul *uint   `gorm:"size:11"`
	IDFakultas   *string `gorm:"size:5"`
	Kapasitas    *int    `gorm:"size:3;default:0"`
	KapUjian     *int    `gorm:"size:3;default:0"`
	Status       *int8   `gorm:"default:1;comment:1: buka 0: tutup 2: hapus"`
	Luas         string  `gorm:"size:5;default:'0';comment:meter persegi"`
	Kondisi      *string `gorm:"size:50"`
	Jumlah       *int    `gorm:"size:5"`
}

func (r *Ruang) TableName() string {
	return "ruang"
}
