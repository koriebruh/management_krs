package domain

import "time"

type IPSemester struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	TA             int    `gorm:"not null;default:0;uniqueIndex:nim"`
	NimDinus       string `gorm:"size:50;not null;uniqueIndex:nim"`
	SKS            int    `gorm:"not null"`
	IPS            string `gorm:"size:5;not null"`
	LastUpdate     *time.Time
	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
}

func (i *IPSemester) TableName() string {
	return "ip_semester"
}
