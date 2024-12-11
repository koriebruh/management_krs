package domain

import "time"

type TahunAjaran struct {
	ID             int64     `gorm:"primaryKey;autoIncrement"`
	Kode           string    `gorm:"size:255;not null"`
	TahunAkhir     string    `gorm:"size:255;not null"`
	TahunAwal      string    `gorm:"size:255;not null"`
	JnsSmt         int       `gorm:"size:1;not null"`
	SetAktif       bool      `gorm:"type:tinyint(1)"`
	BikuTagihJenis int       `gorm:"size:1;default:0"`
	UpdateTime     time.Time `gorm:"type:datetime"`
	UpdateID       string    `gorm:"size:18"`
	UpdateHost     string    `gorm:"size:18"`
	AddedTime      time.Time `gorm:"type:datetime"`
	AddedID        string    `gorm:"size:18"`
	AddedHost      string    `gorm:"size:18"`
	TglMasuk       time.Time `gorm:"type:date"`
}
