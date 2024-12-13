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

	//err = db.Transaction(func(tx *gorm.DB) error {
	//	// Migrasi tabel yang tidak memiliki dependensi foreign key
	//	if err := tx.Debug().AutoMigrate(
	//		&domain.Hari{},              // Hari
	//		&domain.SesiKuliah{},        // SesiKuliah
	//		&domain.SesiKuliahBentrok{}, // SesiKuliahBentrok
	//		&domain.Ruang{},             // Ruang
	//		&domain.TahunAjaran{},       // TahunAjaran
	//	); err != nil {
	//		return err
	//	}
	//
	//	// Migrasi tabel yang memiliki dependensi foreign key
	//	if err := tx.Debug().AutoMigrate(
	//		&domain.KRSRecord{},       // KRSRecord
	//		&domain.MatkulKurikulum{}, // MatkulKurikulum
	//		&domain.JadwalTawar{},     // JadwalTawar
	//		&domain.TagihanMhs{},      // TagihanMhs
	//		&domain.IPSemester{},      // IPSemester
	//	); err != nil {
	//		return err
	//	}
	//
	//	// Migrasi tabel MhsIjinKRS sebelum MahasiswaDinus
	//	if err := tx.Debug().AutoMigrate(
	//		&domain.MhsIjinKRS{}, // MhsIjinKRS
	//	); err != nil {
	//		return err
	//	}
	//
	//	// Migrasi tabel MahasiswaDinus yang tergantung pada tabel sebelumnya
	//	if err := tx.Debug().AutoMigrate(
	//		&domain.MahasiswaDinus{}, // MahasiswaDinus
	//	); err != nil {
	//		return err
	//	}
	//
	//	// Migrasi tabel lainnya dengan dependensi yang sudah terpenuhi
	//	if err := tx.Debug().AutoMigrate(
	//		&domain.KRSRecordLog{},       // KRSRecordLog
	//		&domain.MhsDipaketkan{},      // MhsDipaketkan
	//		&domain.HerregistMahasiswa{}, // HerregistMahasiswa
	//		&domain.DaftarNilai{},        // DaftarNilai
	//		&domain.ValidasiKRSMhs{},     // ValidasiKRSMhs
	//	); err != nil {
	//		return err
	//	}
	//
	//	// Semua migrasi berhasil
	//	return nil
	//})

	if err != nil {
		log.Fatalf("failed to migrate in data base %v", err)
	}

	slog.Info("connection establish")
	return db
}
