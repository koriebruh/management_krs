package domain

type MatkulKurikulum struct {
	KurID         uint    `gorm:"size:11"`
	Kdmk          string  `gorm:"primaryKey;size:255"`
	Nmmk          *string `gorm:"size:255"`
	Nmen          *string `gorm:"size:255"`
	TP            *string `gorm:"type:enum('T', 'P', 'TP')"`
	SKS           *int    `gorm:"size:7"`
	SKST          *int16  `gorm:"size:3"`
	SKSP          *int16  `gorm:"size:3"`
	Smt           *uint
	JnsSmt        *int8
	Aktif         *int8
	KurNama       *string `gorm:"size:255"`
	KelompokMakul *string `gorm:"type:enum('MPK', 'MKK', 'MKB', 'MKD', 'MBB', 'MPB')"`
	KurAktif      *bool
	JenisMatkul   *string `gorm:"type:enum('wajib', 'pilihan')"`
}

func (m *MatkulKurikulum) TableName() string {
	return "matkul_kurikulum"
}
