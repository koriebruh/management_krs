package domain

type Hari struct {
	ID     uint8  `gorm:"primaryKey"`
	Nama   string `gorm:"size:6;not null"`
	NamaEn string `gorm:"size:20;not null"`
}

func (h *Hari) TableName() string {
	return "hari"
}
