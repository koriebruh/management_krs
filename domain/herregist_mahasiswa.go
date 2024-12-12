package domain

import "time"

type HerregistMahasiswa struct {
	ID             uint    `gorm:"primaryKey;autoIncrement"`
	NimDinus       *string `gorm:"size:50"`
	TA             *string `gorm:"size:5"`
	DateReg        *time.Time
	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
}

func (h *HerregistMahasiswa) TableName() string {
	return "herregis_mahasiswa"
}
