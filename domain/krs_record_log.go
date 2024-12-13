package domain

import "time"

type KrsRecordLog struct {
	IDKrs      int             `gorm:"default:null"`
	NimDinus   string          `gorm:"default:null"`
	Kdmk       string          `gorm:"default:null"`
	Aksi       int8            `gorm:"default:null;comment:1=insert,2=delete"`
	IDJadwal   int             `gorm:"default:null"`
	IpAddr     string          `gorm:"default:null"`
	LastUpdate time.Time       `gorm:"not null"`
	KrsRecord  KrsRecord       `gorm:"foreignKey:IDKrs;references:ID"`
	Mahasiswa  MahasiswaDinus  `gorm:"foreignKey:NimDinus;references:NimDinus"`
	MataKuliah MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
}

func (k *KrsRecordLog) TableName() string {
	return "krs_record_log"
}
