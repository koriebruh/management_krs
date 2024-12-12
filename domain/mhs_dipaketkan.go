package domain

type MhsDipaketkan struct {
	NimDinus       string `gorm:"primaryKey;size:50"`
	TAMasukMhs     *int
	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
}

func (m *MhsDipaketkan) TableName() string {
	return "mhs_dipaketkan"
}
