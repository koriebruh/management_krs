package domain

import "time"

type SesiKuliah struct {
	ID         uint8  `gorm:"primaryKey;autoIncrement"`
	Jam        string `gorm:"size:11;not null;default:''"`
	SKS        uint8  `gorm:"not null;default:0"`
	JamMulai   *time.Time
	JamSelesai *time.Time
	Status     *int `gorm:"default:1;comment:0=tidak valid, 1= jam valid(kelipatan 50menit), 2 = jam yang harusnya tisak di pakai(jam istirahat)"`
}

func (s *SesiKuliah) TableName() string {
	return "sesi_kuliah"
}
