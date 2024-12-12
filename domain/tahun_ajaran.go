package domain

import "time"

type TahunAjaran struct {
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	Kode           string `gorm:"not null;uniqueIndex"`
	TahunAkhir     string `gorm:"not null"`
	TahunAwal      string `gorm:"not null"`
	JnsSmt         int8   `gorm:"not null;comment:1 = reg ganjil, 2 = reg genap, 3 = sp ganjil, 4 = sp genap"`
	SetAktif       int8   `gorm:"not null"`
	BikuTagihJenis *int8  `gorm:"default:0;comment:1 = spp; 2 = sks; 3 = kekurangan"`
	UpdateTime     *time.Time
	UpdateID       *string `gorm:"size:18"`
	UpdateHost     *string `gorm:"size:18"`
	AddedTime      *time.Time
	AddedID        *string `gorm:"size:18"`
	AddedHost      *string `gorm:"size:18"`
	TglMasuk       *time.Time
	TahunAjarans   []TahunAjaran `gorm:"foreignkey:Kode"`
}

func (t *TahunAjaran) TableName() string {
	return "tahun_ajaran"
}
