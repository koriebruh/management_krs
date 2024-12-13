package domain

type SesiKuliahBentrok struct {
	ID          int        `gorm:"primaryKey"`
	IDBentrok   int        `gorm:"index"`
	Sesi        SesiKuliah `gorm:"foreignKey:ID;references:ID"`
	SesiBentrok SesiKuliah `gorm:"foreignKey:IDBentrok;references:ID"`
}

func (s *SesiKuliahBentrok) TableName() string {
	return "sesi_kuliah_bentrok"
}
