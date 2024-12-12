package domain

type DaftarNilai struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	//NIM DINUS KALO DI BIKIN NULL ERR
	NimDinus        string          `gorm:"size:50;index:nim"`
	Kdmk            *string         `gorm:"size:20;index:nim"`
	Nl              *string         `gorm:"size:2"`
	Hide            *int8           `gorm:"default:0;comment:0 = nilai muncul; 1 = nilai disembunyikan (utk keperluan spt hapus mata kuliah)"`
	MahasiswaDinus  MahasiswaDinus  `gorm:"foreignkey:NimDinus"`
	MatkulKurikulum MatkulKurikulum `gorm:"foreignkey:Kdmk"`
}

func (d *DaftarNilai) TableName() string {
	return "daftar_nilai"
}
