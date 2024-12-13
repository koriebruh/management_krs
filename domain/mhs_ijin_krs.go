package domain

import "time"

type MhsIjinKrs struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	TA       int    `gorm:"not null;default:0"`
	NimDinus string `gorm:"default:null"`
	Ijinkan  bool   `gorm:"default:null"`
	//Time        time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Time        time.Time      `gorm:"not null;"`
	TahunAjaran TahunAjaran    `gorm:"foreignKey:TA;references:Kode"`
	Mahasiswa   MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
}

func (m *MhsIjinKrs) TableName() string {
	return "mhs_ijin_krs"
}
