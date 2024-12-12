package domain

import "time"

type MhsIjinKRS struct {
	ID             uint    `gorm:"primaryKey;autoIncrement"`
	TA             *int    `gorm:"uniqueIndex:nim"`
	NimDinus       *string `gorm:"size:50;uniqueIndex:nim"`
	Ijinkan        *int8
	Time           time.Time      `gorm:"autoCreateTime"`
	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
}

func (m *MhsIjinKRS) TableName() string {
	return "mhs_ijin_krs"
}
