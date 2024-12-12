package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"koriebruh/try/domain"
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

	//17 table
	if err = db.AutoMigrate(
		//&domain.Hari{},
		//&domain.SesiKuliah{},        // rancu
		//&domain.SesiKuliahBentrok{}, // rancu
		//&domain.Ruang{},
		//&domain.TahunAjaran{},
		//&domain.MahasiswaDinus{},
		//&domain.MatkulKurikulum{},
		//&domain.JadwalTawar{},
		//&domain.TagihanMhs{},
		//&domain.IPSemester{},
		//&domain.KRSRecord{},
		//&domain.KRSRecordLog{},
		//&domain.MhsIjinKRS{},
		//&domain.HerregistMahasiswa{},
		//&domain.MhsDipaketkan{},
		//&domain.DaftarNilai{},
		//&domain.ValidasiKRSMhs{},

		&domain.Hari{},
		&domain.SesiKuliah{},
		&domain.SesiKuliahBentrok{},
		&domain.Ruang{},
		&domain.TahunAjaran{},
		&domain.KRSRecord{},       // Migrasi tabel krs_record terlebih dahulu
		&domain.MatkulKurikulum{}, // Migrasi tabel matkul_kurikulum setelah krs_record
		&domain.JadwalTawar{},
		&domain.TagihanMhs{},
		&domain.IPSemester{},
		&domain.MahasiswaDinus{},
		&domain.KRSRecordLog{},
		&domain.MhsIjinKRS{},
		&domain.HerregistMahasiswa{},
		&domain.MhsDipaketkan{},
		&domain.DaftarNilai{},
		&domain.ValidasiKRSMhs{},
	); err != nil {
		log.Fatalf("failed to migrate in data base %v", err)
	}

	slog.Info("success migrate")
	return db
}
