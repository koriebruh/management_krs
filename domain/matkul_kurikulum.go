package domain

type MatkulKurikulum struct {
	KurID         int    `gorm:"column:kur_id"`
	Kdmk          string `gorm:"index;not null"`
	Nmmk          string
	Nmen          string
	Tp            string `gorm:"type:enum('T','P','TP')"`
	Sks           int
	SksT          int16
	SksP          int16
	Smt           int
	JnsSmt        int
	Aktif         bool
	KurNama       string
	KelompokMakul string `gorm:"type:enum('MPK','MKK','MKB','MKD','MBB','MPB')"`
	KurAktif      bool   `gorm:"type:bit(1)"`
	JenisMatkul   string `gorm:"type:enum('wajib','pilihan')"`
}

func (m *MatkulKurikulum) TableName() string {
	return "matkul_kurikulum"
}
