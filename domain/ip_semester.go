package domain

import (
	"time"
)

type IpSemester struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	TA          int            `gorm:"not null;default:0"`
	NimDinus    string         `gorm:"not null"`
	Sks         int            `gorm:"not null"`
	Ips         string         `gorm:"not null"`
	LastUpdate  time.Time      `gorm:"default:null"`
	TahunAjaran TahunAjaran    `gorm:"foreignKey:TA;references:Kode"`
	Mahasiswa   MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
}

func (i *IpSemester) TableName() string {
	return "ip_semester"
}
