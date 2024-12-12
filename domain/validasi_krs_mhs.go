package domain

import "time"

type ValidasiKRSMhs struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	NimDinus       string `gorm:"size:50;not null"`
	JobDate        *time.Time
	JobHost        *string `gorm:"size:255"`
	JobAgent       *string `gorm:"size:255"`
	TA             *int
	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
}

func (v *ValidasiKRSMhs) TableName() string {
	return "validasi_krs_mhs"
}
