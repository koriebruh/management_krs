package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"koriebruh/try/conf"
	"koriebruh/try/domain"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

////////////////////////////////////////
/*
	&domain.MahasiswaDinus{},
	&domain.MatkulKurikulum{},

 	&domain.Hari{},
	&domain.SesiKuliah{},
	&domain.Ruang{},
	&domain.TahunAjaran{},
*/

func IfErrNotNil(err error) {
	if err != nil {
		log.Fatalf("err di sini bg, %e", err)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0 // Atau bisa return nilai default lainnya
	}
	return i
}

// AGA LAMA TOLOL 3OK ASW NTAR BUAT CONCURENT
func TestInsertMahasiswaDinus(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/mahasiswa_dinus.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 5 {
			log.Printf("Skipping line %d due to insufficient columns: %v", i, record)
			continue
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(record[5]), bcrypt.DefaultCost)

		mhsDinus := domain.MahasiswaDinus{
			NimDinus: record[0],
			TaMasuk:  int(atoi(record[1])),
			Prodi:    record[2],
			PassMhs:  string(password),
			Kelas:    atoi(record[4]),
			AkdmStat: record[3],
		}

		if err := db.Create(&mhsDinus).Error; err != nil {
			log.Fatalf("err in line %v and err bcs %e", i, err)
		}

	}

	log.Println("YEY SUCCESS")

}

func TestInsertMatkulkurikulum(t *testing.T) {
	db := conf.InitDB()

	file, err := os.Open("data_krs/matkul_kurikulum.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 15 {
			log.Printf("Skipping line %d due to insufficient columns: %v", i, record)
			continue
		}

		mk := domain.MatkulKurikulum{
			KurID:         atoi(record[0]),
			Kdmk:          record[1],
			Nmmk:          record[2],
			Nmen:          record[3],
			Tp:            record[4],
			Sks:           atoi(record[5]),
			SksT:          int16(atoi(record[6])),
			SksP:          int16(atoi(record[7])),
			Smt:           atoi(record[8]),
			JnsSmt:        atoi(record[9]),
			Aktif:         record[10] == "true",
			KurNama:       record[11],
			KelompokMakul: record[12],
			KurAktif:      record[13] == "true",
			JenisMatkul:   record[14],
		}

		if err := db.Create(&mk).Error; err != nil {
			log.Fatalf("err in line %v and err bcs %e", i, err)
		}

		log.Println("success nih bg line ke ", i)

	}
}

// gtw err input id 0 ga bisa bodomat tolol
func TestInsertHari(t *testing.T) {
	db := conf.InitDB()
	hariData := []domain.Hari{
		{ID: 0, Nama: "-", NamaEn: "-"},
		{ID: 1, Nama: "SENIN", NamaEn: "MONDAY"},
		{ID: 2, Nama: "SELASA", NamaEn: "TUESDAY"},
		{ID: 3, Nama: "RABU", NamaEn: "WEDNESDAY"},
		{ID: 4, Nama: "KAMIS", NamaEn: "THURSDAY"},
		{ID: 5, Nama: "JUMAT", NamaEn: "FRIDAY"},
		{ID: 6, Nama: "SABTU", NamaEn: "SATURDAY"},
		{ID: 7, Nama: "ALLDAY", NamaEn: "ALLDAY"},
		{ID: 8, Nama: "MINGGU", NamaEn: "SUNDAY"},
	}

	for _, hari := range hariData {
		if err := db.Create(&hari).Error; err != nil {
			log.Fatalf("Failed to insert data: %v", err)
		}
	}

	log.Println("YEY SUCCESS")
}

func TestInsertSesiKuliah(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/sesi_kuliah.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {
		if i == 0 {
			continue
		}
		jamMulai, err := parseTimeOnly(record[3])
		if err != nil {
			log.Printf("Skipping line %d due to invalid jamMulai format: %v", i, err)
			continue
		}

		jamSelesai, err := parseTimeOnly(record[4])
		if err != nil {
			log.Printf("Skipping line %d due to invalid jamSelesai format: %v", i, err)
			continue
		}

		sesiKuliah := domain.SesiKuliah{
			Jam:        record[1],
			Sks:        int16(atoi(record[2])),        // Kolom SKS
			JamMulai:   jamMulai.Format("15:04:05"),   // Format ke HH:mm:ss
			JamSelesai: jamSelesai.Format("15:04:05"), // Format ke HH:mm:ss
			Status:     atoi(record[5]),               // Kolom Status
		}

		// Insert ke database
		if err := db.Create(&sesiKuliah).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}
	}

	log.Println("YEY SUCCESS")
}

func parseTimeOnly(timeStr string) (time.Time, error) {
	layout := "15:04:05" // Format waktu (jam, menit, detik)
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %s, error: %v", timeStr, err)
	}
	return parsedTime, nil
}

func TestInsertRuang(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/ruang.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {

		if i == 0 {
			continue
		}

		ruang := domain.Ruang{
			ID:           atoi(record[0]),
			Nama:         record[1],
			Nama2:        record[2],
			IDJenisMakul: atoi(record[3]),
			IDFakultas:   record[4],
			Kapasitas:    atoi(record[5]),
			KapUjian:     atoi(record[6]),
			Status:       int16(atoi(record[7])),
			Luas:         record[8],
			Kondisi:      record[9],
			Jumlah:       atoi(record[10]),
		}

		if err := db.Create(&ruang).Error; err != nil {
			log.Fatalf("err in line %v and err bcs %e", i, err)
		}
	}

	log.Println("YEY SUCCESS")
}

func TestInsertTahunAjaran(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/tahun_ajaran.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	parseTime := func(timeStr string) time.Time {
		if timeStr == "" {
			return time.Time{}
		}
		// Format waktu sesuai dengan CSV
		layout := "2006-01-02 15:04:05.000"
		parsedTime, err := time.Parse(layout, timeStr)
		if err != nil {
			log.Printf("Failed to parse time: %s, error: %v", timeStr, err)
			return time.Time{}
		}
		return parsedTime
	}

	for i, record := range records {

		if i == 0 {
			continue
		}

		ajaran := domain.TahunAjaran{
			ID:             int64(atoi(record[0])),
			Kode:           atoi(record[1]),
			TahunAkhir:     record[2],
			TahunAwal:      record[3],
			JnsSmt:         atoi(record[4]),
			SetAktif:       record[5] == "1", //if 1 return true
			BikuTagihJenis: int8(atoi(record[6])),
			UpdateTime:     parseTime(record[7]),
			UpdateID:       record[8],
			UpdateHost:     record[9],
			AddedTime:      parseTime(record[10]),
			AddedID:        record[11],
			AddedHost:      record[12],
			TglMasuk:       parseTime(record[13]),
		}

		if err := db.Create(&ajaran).Error; err != nil {
			log.Fatalf("err in line %v and err bcs %e", i, err)
		}
	}

	log.Println("YEY SUCCESS")
}

///////////////////////////////////////////
/*
	&domain.TagihanMhs{},
	&domain.IpSemester{},
	&domain.JadwalInputKrs{},
	&domain.MhsIjinKrs{},
	&domain.HerregistMahasiswa{},
	&domain.MhsDipaketkan{},
	&domain.DaftarNilai{},
	&domain.ValidasiKrsMhs{},
	&domain.SesiKuliahBentrok{},
*/
