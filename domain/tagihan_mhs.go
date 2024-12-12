package domain

import "time"

type TagihanMhs struct {
	ID             uint       `gorm:"primaryKey;autoIncrement;comment:id biasa"`
	TA             int        `gorm:"not null;comment:tahun ajaran"`
	NimDinus       string     `gorm:"size:50;not null;comment:nim mahasiswa;uniqueIndex:nim"`
	SPPBank        *string    `gorm:"size:11"`
	SPPBayar       int8       `gorm:"not null;default:0;comment:status bayar spp 1: bayar 0: belum bayar"`
	SPPBayarDate   *time.Time `gorm:"comment:tanggal pada saat operator input pembayaran"`
	SPPDispensasi  *int
	SPPHost        *string        `gorm:"size:25;comment:ip/host operator"`
	SPPStatus      int8           `gorm:"not null;comment:1 : full payment 0 : dispensasi"`
	SPPTransaksi   *string        `gorm:"size:20;comment:jenis pembayaran : langsung, transfer"`
	TahunAjaran    TahunAjaran    `gorm:"foreignkey:TA"`
	MahasiswaDinus MahasiswaDinus `gorm:"foreignkey:NimDinus"`
}

func (t *TagihanMhs) TableName() string {
	return "tagihan_mhs"
}
