package domain

type MahasiswaDinus struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	NimDinus string `gorm:"unique;not null;default:''"`
	TaMasuk  int
	Prodi    string
	PassMhs  string `gorm:"default:null"`
	Kelas    int    `gorm:"not null;comment:0 = not choose,1 = pagi,2 = malam,3 = karyawan/pegawai"`
	AkdmStat string `gorm:"type:char(2);not null;comment:1 = aktif, 2 = cuti, 3 = keluar, 4 = lulus, 5 = mangkir, 6 = meninggal, 7 = DO, 8 = Aktif Keuangan"`
}

func (m *MahasiswaDinus) TableName() string {
	return "mahasiswa_dinus"
}
