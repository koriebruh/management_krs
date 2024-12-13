package domain

import "time"

type TahunAjaran struct {
	ID             int64     `gorm:"primaryKey;autoIncrement"`
	Kode           int       `gorm:"unique;not null"`
	TahunAkhir     string    `gorm:"not null"`
	TahunAwal      string    `gorm:"not null"`
	JnsSmt         int       `gorm:"not null;comment:1 = reg ganjil, 2 = reg genap, 3 = sp ganjil, 4 = sp genap"`
	SetAktif       bool      `gorm:"not null"`
	BikuTagihJenis int8      `gorm:"default:0;comment:1 = spp; 2 = sks; 3 = kekurangan"`
	UpdateTime     time.Time `gorm:"default:null"`
	UpdateID       string    `gorm:"default:null"`
	UpdateHost     string    `gorm:"default:null"`
	AddedTime      time.Time `gorm:"default:null"`
	AddedID        string    `gorm:"default:null"`
	AddedHost      string    `gorm:"default:null"`
	TglMasuk       time.Time `gorm:"type:date;default:null"`
}

func (t *TahunAjaran) TableName() string {
	return "tahun_ajaran"
}
