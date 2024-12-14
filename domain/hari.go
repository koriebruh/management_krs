package domain

type Hari struct {
	ID     int8   `gorm:"unique;not null;autoIncrement:false"` //tadi di tambah auto iinc false
	Nama   string `gorm:"type:varchar(6);not null"`
	NamaEn string `gorm:"type:varchar(20);not null"`
}

func (h *Hari) TableName() string {
	return "hari"
}
