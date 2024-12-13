package domain

type MhsDipaketkan struct {
	NimDinus   string `gorm:"primaryKey"`
	TaMasukMhs int
	Mahasiswa  MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
}

func (m *MhsDipaketkan) TableName() string {
	return "mhs_dipaketkan"
}
