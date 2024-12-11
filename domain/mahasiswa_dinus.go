package domain

type MahasiswaDinus struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	NimDinus      string `gorm:"size:50;unique;not null"`
	TAMasuk       int    `gorm:"size:10"`
	Prodi         string `gorm:"size:5"`
	PassMhs       string `gorm:"size:128;default:null"`
	Kelas         int    `gorm:"size:1;comment:0 = not choose, 1 = pagi, 2 = malam, 3 = karyawan/pegawai"`
	AkdmStat      string `gorm:"size:2;not null;comment:1 = aktif, 2 = cuti, 3 = keluar, 4 = lulus, 5 = mangkir, 6 = meninggal, 7 = DO"`
	TahunAjaranID string `gorm:"foreignKey:NimDinus;references:NimDinus"`
}
