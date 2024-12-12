package domain

type KRSRecord struct {
	ID              uint64          `gorm:"primaryKey;autoIncrement"`
	TA              string          `gorm:"not null;index:PERIODE"`
	Kdmk            string          `gorm:"size:255;not null"`
	IDJadwal        uint            `gorm:"not null"`
	NimDinus        string          `gorm:"size:50;not null;index:MAHASISWA"`
	Sts             string          `gorm:"size:1;not null"`
	SKS             int             `gorm:"not null"`
	Modul           int8            `gorm:"not null;default:0"`
	TahunAjaran     TahunAjaran     `gorm:"foreignkey:TA"`
	MatkulKurikulum MatkulKurikulum `gorm:"foreignkey:Kdmk"`
	JadwalTawar     JadwalTawar     `gorm:"foreignkey:IDJadwal"`
	MahasiswaDinus  MahasiswaDinus  `gorm:"foreignkey:NimDinus"`
}

func (k *KRSRecord) TableName() string {
	return "krs_record"
}
