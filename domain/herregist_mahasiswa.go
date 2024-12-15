package domain

import (
	"time"
)

type HerregistMahasiswa struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	NimDinus    string         `gorm:"default:null"`
	TA          int            `gorm:"not null;default:0"`
	DateReg     time.Time      `gorm:"default:null"`
	Mahasiswa   MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
	TahunAjaran TahunAjaran    `gorm:"foreignKey:TA;references:Kode"`
}

func (h *HerregistMahasiswa) TableName() string {
	return "herregist_mahasiswa"
}
