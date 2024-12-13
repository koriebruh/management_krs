package domain

import "time"

type TagihanMhs struct {
	ID            int    `gorm:"primaryKey;autoIncrement;comment:id biasa"`
	TA            int    `gorm:"not null;comment:tahun ajaran"`
	NimDinus      string `gorm:"not null;comment:nim mahasiswa"`
	SppBank       string
	SppBayar      int       `gorm:"not null;default:0;comment:status bayar spp 1: bayar 0: belum bayar"`
	SppBayarDate  time.Time `gorm:"comment:tanggal pada saat operator input pembayaran"`
	SppDispensasi int
	SppHost       string         `gorm:"comment:ip/host operator"`
	SppStatus     int            `gorm:"not null;comment:1 : full payment 0 : dispensasi"`
	SppTransaksi  string         `gorm:"comment:jenis pembayaran : langsung, transfer"`
	TahunAjaran   TahunAjaran    `gorm:"foreignKey:TA;references:Kode"`
	Mahasiswa     MahasiswaDinus `gorm:"foreignKey:NimDinus;references:NimDinus"`
}

func (t *TagihanMhs) TableName() string {
	return "tagihan_mhs"
}
