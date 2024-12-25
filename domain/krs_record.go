package domain

import "gorm.io/gorm"

type KrsRecord struct {
	ID          int             `gorm:"primaryKey;autoIncrement=false"`
	TA          int             `gorm:"not null;default:0"`
	Kdmk        string          `gorm:"not null"`
	IDJadwal    int             `gorm:"not null"`
	NimDinus    string          `gorm:"not null"`
	Sts         string          `gorm:"type:char(1);not null"`
	Sks         int             `gorm:"not null"`
	Modul       int             `gorm:"not null;default:0"`
	TahunAjaran TahunAjaran     `gorm:"foreignKey:TA;references:Kode"`
	MataKuliah  MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
	Jadwal      JadwalTawar     `gorm:"foreignKey:IDJadwal;references:ID"`
	Mahasiswa   MahasiswaDinus  `gorm:"foreignKey:NimDinus;references:NimDinus"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
}

func (k *KrsRecord) TableName() string {
	return "krs_record"
}
