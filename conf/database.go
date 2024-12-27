package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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
		log.Fatalf("failed make connection to database %v", err)
	}
	//
	//// Disable foreign key checks before migration
	//db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	//
	//err = db.Transaction(func(tx *gorm.DB) error {
	//	// Independent tables (no foreign key dependencies)
	//	if err := tx.AutoMigrate(
	//		&domain.MahasiswaDinus{},
	//		&domain.MatkulKurikulum{},
	//		&domain.Hari{},
	//		&domain.SesiKuliah{},
	//		&domain.Ruang{},
	//		&domain.TahunAjaran{},
	//	); err != nil {
	//		return err
	//	}
	//
	//	// Tables with single-level dependencies
	//	if err := tx.AutoMigrate(
	//		&domain.TagihanMhs{},
	//		&domain.IpSemester{},
	//		&domain.JadwalInputKrs{},
	//		&domain.MhsIjinKrs{},
	//		&domain.HerregistMahasiswa{},
	//		&domain.MhsDipaketkan{},
	//		&domain.DaftarNilai{},
	//		&domain.ValidasiKrsMhs{},
	//		&domain.SesiKuliahBentrok{},
	//	); err != nil {
	//		return err
	//	}
	//
	//	// Tables with multiple dependencies
	//	if err := tx.AutoMigrate(
	//		&domain.JadwalTawar{},
	//		&domain.KrsRecord{},
	//		&domain.KrsRecordLog{},
	//	); err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})
	////PAKSA TIME  ORM NYA TOLOL MAKSA datetime
	//db.Exec(`
	//ALTER TABLE sesi_kuliahs
	//MODIFY COLUMN jam_mulai TIME,
	//MODIFY COLUMN jam_selesai TIME;
	//`)
	//
	//// Re-enable foreign key checks after migration
	//db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	//SETUP CONNECTION POOL
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.DB from GORM: %v", err)
	}
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(80)
	sqlDB.SetConnMaxLifetime(60)

	//if err != nil {
	//	log.Fatalf("failed to migrate in database: %v", err)
	//}

	slog.Info("connection established")
	return db
}
