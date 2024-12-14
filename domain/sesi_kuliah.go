package domain

type SesiKuliah struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	Jam        string `gorm:"not null;default:''"`
	Sks        int16  `gorm:"not null;default:0"`
	JamMulai   string `gorm:"type:TIME"` // Tipe waktu
	JamSelesai string `gorm:"type:TIME"`
	Status     int    `gorm:"default:1;comment:0=tidak valid, 1= jam valid(kelipatan 50menit), 2 = jam yang harusnya tisak di pakai(jam istirahat)"`
}

func (s *SesiKuliah) TableName() string {
	return "sesi_kuliah"
}
