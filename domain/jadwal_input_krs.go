package domain

import "time"

type JadwalInputKrs struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	TA         int    `gorm:"not null;default:0"`
	Prodi      string `gorm:"type:char(3)"`
	TglMulai   time.Time
	TglSelesai time.Time
}

func (j *JadwalInputKrs) TableName() string {
	return "jadwal_input_krs"
}
