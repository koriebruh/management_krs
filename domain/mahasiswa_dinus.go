package domain

type MahasiswaDinus struct {
	ID       uint    `gorm:"primaryKey;autoIncrement"`
	NimDinus string  `gorm:"column:nim_dinus;type:varchar(50);not null;uniqueIndex"`
	TaMasuk  *int    `gorm:"size:10"`
	Prodi    *string `gorm:"size:5"`
	PassMhs  *string `gorm:"size:128"`
	Kelas    int8    `gorm:"not null;comment:0 = not choose,1 = pagi,2 = malam,3 = karyawan/pegawai"`
	AkdmStat string  `gorm:"size:2;not null;comment:1 = aktif, 2 = cuti, 3 = keluar, 4 = lulus, 5 = mangkir, 6 = meninggal, 7 = DO, 8 = Aktif Keuangan;index:STS_AKD_MHS"`
}

func (m *MahasiswaDinus) TableName() string {
	return "mahasiswa_dinus"
}
