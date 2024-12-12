package domain

type JadwalTawar struct {
	ID              uint    `gorm:"primaryKey;autoIncrement"`
	TA              int     `gorm:"not null;default:0"`
	Kdmk            string  `gorm:"size:15;not null"`
	Klpk            string  `gorm:"size:15;not null"`
	Klpk2           *string `gorm:"size:15"`
	Kdds            int     `gorm:"not null"`
	Kdds2           *int
	Jmax            *int            `gorm:"size:3;default:0"`
	Jsisa           *int            `gorm:"size:3;default:0"`
	IDHari1         uint8           `gorm:"not null"`
	IDHari2         uint8           `gorm:"not null"`
	IDHari3         uint8           `gorm:"not null"`
	IDSesi1         uint16          `gorm:"size:3;not null"`
	IDSesi2         uint16          `gorm:"size:3;not null"`
	IDSesi3         uint16          `gorm:"size:3;not null"`
	IDRuang1        uint            `gorm:"size:3;not null"`
	IDRuang2        uint            `gorm:"size:3;not null"`
	IDRuang3        uint            `gorm:"size:3;not null"`
	JnsJam          uint8           `gorm:"not null;comment:1=pagi, 2=malam, 3=pagi-malam"`
	OpenClass       uint8           `gorm:"not null;default:1;comment:kelas dibuka utk KRS : 1 = open; 0 = close"`
	Hari1           Hari            `gorm:"foreignKey:IDHari1"`
	Hari2           Hari            `gorm:"foreignKey:IDHari2"`
	Hari3           Hari            `gorm:"foreignKey:IDHari3"`
	Sesi1           SesiKuliah      `gorm:"foreignKey:IDSesi1"`
	Sesi2           SesiKuliah      `gorm:"foreignKey:IDSesi2"`
	Sesi3           SesiKuliah      `gorm:"foreignKey:IDSesi3"`
	Ruang1          Ruang           `gorm:"foreignKey:IDRuang1"`
	Ruang2          Ruang           `gorm:"foreignKey:IDRuang2"`
	Ruang3          Ruang           `gorm:"foreignKey:IDRuang3"`
	MatkulKurikulum MatkulKurikulum `gorm:"foreignKey:Kdmk"`
}

func (j *JadwalTawar) TableName() string {
	return "jatwal_tawar"
}
