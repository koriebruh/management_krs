package domain

type DaftarNilai struct {
	ID         int             `gorm:"column:_id;primaryKey;autoIncrement"`
	NimDinus   string          `gorm:"default:null"`
	Kdmk       string          `gorm:"default:null"`
	Nl         string          `gorm:"type:char(2);default:null"`
	Hide       int16           `gorm:"default:0;comment:0 = nilai muncul;1 = nilai disembunyikan (utk keperluan spt hapus mata kuliah)"`
	Mahasiswa  MahasiswaDinus  `gorm:"foreignKey:NimDinus;references:NimDinus"`
	MataKuliah MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
}

func (d *DaftarNilai) TableName() string {
	return "daftar_nilai"
}
