package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"koriebruh/try/domain"
	"log/slog"
)

func InitDB() *gorm.DB {

	config := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DataBase.User,
		config.DataBase.Pass,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		slog.Error("failed make connection to database", err)
	}

	if err = db.AutoMigrate(
		&domain.User{},
	); err != nil {
		slog.Error("failed auto migrate ", err)
	}

	slog.Info("success migrate")
	return db
}
