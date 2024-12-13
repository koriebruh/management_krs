package domain

import "time"

type ValidasiKrsMhs struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	NimDinus    string         `gorm:"not null"`
	JobDate     time.Time      `gorm:"default:null"`
	JobHost     string         `gorm:"default:null"`
	JobAgent    string         `gorm:"default:null"`
	TA          int            `gorm:"not null;default:0"`
	Mahasiswa   MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
	TahunAjaran TahunAjaran    `gorm:"foreignKey:TA;references:Kode"`
}

func (v *ValidasiKrsMhs) TableName() string {
	return "validasi_krs_mhs"
}
