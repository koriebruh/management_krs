package domain

type SesiKuliahBentrok struct {
	ID        uint16 `gorm:"primaryKey"`
	IDBentrok uint16 `gorm:"primaryKey"`
	//Sesi      SesiKuliah `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	//Bentrok   SesiKuliah `gorm:"foreignKey:IDBentrok;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;index:FK_sesi_kuliah_bentrok2"`
}

func (s *SesiKuliahBentrok) TableName() string {
	return "sesi_kuliah_bentrok"
}
