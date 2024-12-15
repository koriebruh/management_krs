package domain

// JadwalTawar represents the jadwal_tawar table
type JadwalTawar struct {
	ID         int             `gorm:"primaryKey;autoIncrement=false"`
	TA         int             `gorm:"not null;default:0"`
	Kdmk       string          `gorm:"not null"`
	Klpk       string          `gorm:"not null"`
	Klpk2      string          `gorm:"default:null"`
	Kdds       int             `gorm:"not null"`
	Kdds2      int             `gorm:"default:null"`
	Jmax       int             `gorm:"default:0"`
	Jsisa      int             `gorm:"default:0"`
	IDHari1    int8            `gorm:"not null"`
	IDHari2    int8            `gorm:"not null"`
	IDHari3    int8            `gorm:"not null"`
	IDSesi1    int             `gorm:"not null"`
	IDSesi2    int             `gorm:"not null"`
	IDSesi3    int             `gorm:"not null"`
	IDRuang1   int             `gorm:"not null"`
	IDRuang2   int             `gorm:"not null"`
	IDRuang3   int             `gorm:"not null"`
	JnsJam     int8            `gorm:"not null;comment:1=pagi, 2=malam, 3=pagi-malam"`
	OpenClass  bool            `gorm:"not null;default:1;comment:kelas dibuka utk KRS : 1 = open; 0 = close"`
	MataKuliah MatkulKurikulum `gorm:"foreignKey:Kdmk;references:Kdmk"`
	Hari1      Hari            `gorm:"foreignKey:IDHari1;references:ID"`
	Hari2      Hari            `gorm:"foreignKey:IDHari2;references:ID"`
	Hari3      Hari            `gorm:"foreignKey:IDHari3;references:ID"`
	Sesi1      SesiKuliah      `gorm:"foreignKey:IDSesi1;references:ID"`
	Sesi2      SesiKuliah      `gorm:"foreignKey:IDSesi2;references:ID"`
	Sesi3      SesiKuliah      `gorm:"foreignKey:IDSesi3;references:ID"`
	Ruang1     Ruang           `gorm:"foreignKey:IDRuang1;references:ID"`
	Ruang2     Ruang           `gorm:"foreignKey:IDRuang2;references:ID"`
	Ruang3     Ruang           `gorm:"foreignKey:IDRuang3;references:ID"`
}

func (j *JadwalTawar) TableName() string {
	return "jadwal_tawar"
}
