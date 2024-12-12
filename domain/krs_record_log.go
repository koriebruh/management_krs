package domain

import "time"

type KRSRecordLog struct {
	IDKRS           *uint
	NimDinus        *string `gorm:"size:50"`
	Kdmk            *string `gorm:"size:255"`
	Aksi            *int8   `gorm:"comment:1=insert,2=delete"`
	IDJadwal        *uint
	IPAddr          *string          `gorm:"size:50"`
	LastUpdate      time.Time        `gorm:"autoCreateTime"`
	KRSRecord       *KRSRecord       `gorm:"foreignkey:IDKRS"`
	MahasiswaDinus  *MahasiswaDinus  `gorm:"foreignkey:NimDinus"`
	MatkulKurikulum *MatkulKurikulum `gorm:"foreignkey:Kdmk"`
}

func (k *KRSRecordLog) TableName() string {
	return "krs_record_log"
}
