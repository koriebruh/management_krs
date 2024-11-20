package domain

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	NIM      string `gorm:"not null;unique"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`

	// auto create
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
