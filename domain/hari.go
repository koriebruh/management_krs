package domain

type Hari struct {
	ID     int8   `gorm:"unique;not null"`
	Nama   string `gorm:"type:varchar(6);not null"`
	NamaEn string `gorm:"type:varchar(20);not null"`
}

func (h *Hari) TableName() string {
	return "hari"
}
