package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"koriebruh/try/conf"
	"koriebruh/try/domain"
	"log"
	"os"
	"strconv"
	"sync"
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
			Aktif:         int8(atoi(record[10])),
			KurNama:       record[11],
			KelompokMakul: record[12],
			KurAktif:      int8(atoi(record[13])),
			JenisMatkul:   record[14],
		}

		if err := db.Create(&mk).Error; err != nil {
			log.Fatalf("err in line %v and err bcs %e", i, err)
		}

	}
}

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

// DUNGU DATA NIM DI TABEL TAGHAN AMA DI MAHASISWA DINUS BEDA MANA BISA FORIGEN KEY NYA,
// (INI SUCCES DATA TERAHIR YG EMANG BEGO)
func TestTagihanMhs(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/tagihan_mhs.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	const layout = "2006-01-02 15:04:05.000"

	for i, record := range records {
		if i == 0 {
			continue
		}

		//BUAT NGESEKIP DATA YG TOLOL NIM NYA GA ADA DI TABEL mahasiswa_dinus
		var mahasiswa domain.MahasiswaDinus
		err = db.Where("nim_dinus = ?", record[2]).First(&mahasiswa).Error
		if err != nil {
			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[2], i)
			continue
		}

		sppBayarDate, err := time.Parse(layout, record[4])
		if err != nil {
			log.Fatalf("Error parsing date in line %v: %v", i, err)
		}

		tagihanMhs := domain.TagihanMhs{
			ID:            atoi(record[0]),
			TA:            atoi(record[1]),
			NimDinus:      record[2],
			SppBayar:      atoi(record[3]),
			SppBayarDate:  sppBayarDate,
			SppHost:       record[5],
			SppStatus:     atoi(record[6]),
			SppDispensasi: atoi(record[7]),
			SppBank:       record[8],
			SppTransaksi:  record[9],
		}

		// Insert ke database
		if err := db.Create(&tagihanMhs).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// SKIP SOME DATA NIM  and KODE YG TUDAK ADA JG
// (WORK) AGA LAMA NANTI UBAH PAKE GO ROUTRINE
func TestIpSemester(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/ip_semester.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	const layout = "2006-01-02 15:04:05.000"

	for i, record := range records {
		if i == 0 {
			continue
		}

		if atoi(record[1]) == 20231 && record[5] == "95cedfff35b96393991dd55e982120c2" {
			fmt.Println("harus nya ke insert ini")
		}

		//BUAT NGESEKIP DATA YG TOLOL NIM NYA GA ADA DI TABEL mahasiswa_dinus
		var mahasiswa domain.MahasiswaDinus
		err = db.Where("nim_dinus = ?", record[5]).First(&mahasiswa).Error
		if err != nil {
			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[5], i)
			continue
		}
		log.Printf("NIM INI KETEMU %v", record[5])

		var kodeExist domain.TahunAjaran
		err = db.Where("kode = ?", record[1]).First(&kodeExist).Error
		if err != nil {
			log.Printf("KODE %s not found in tahun_ajaran, skipping line %v", record[1], i)
			continue
		}
		log.Println()

		lu, err := time.Parse(layout, record[4])
		if err != nil {
			log.Fatalf("Error parsing date in line %v: %v", i, err)
		}

		ipSemester := domain.IpSemester{
			ID:         atoi(record[0]),
			TA:         atoi(record[1]),
			Sks:        atoi(record[2]),
			Ips:        record[3],
			LastUpdate: lu,
			NimDinus:   record[5],
		}

		// Insert ke database
		if err := db.Create(&ipSemester).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}
		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// (AMAN)
func TestJadwalInputKrs(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/jadwal_input_krs.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	const layout = "02/01/06 15.04"

	for i, record := range records {
		if i == 0 {
			continue
		}

		tgsMulai, err := time.Parse(layout, record[3])
		if err != nil {
			log.Fatalf("Error parsing date in line %v: %v", i, err)
		}
		tgsSelesai, err := time.Parse(layout, record[4])
		if err != nil {
			log.Fatalf("Error parsing date in line %v: %v", i, err)
		}

		jadwalInputKrs := domain.JadwalInputKrs{
			ID:         atoi(record[0]),
			TA:         atoi(record[1]),
			Prodi:      record[2],
			TglMulai:   tgsMulai,
			TglSelesai: tgsSelesai,
		}

		// Insert ke database
		if err := db.Create(&jadwalInputKrs).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// SKIP SOME DATA NIM YG TOLOL GA ADA
// (AMAN)
func TestMhsIjinKrs(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/mhs_ijin_krs.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	const layout = "2006-01-02 15:04:05.000"

	for i, record := range records {
		if i == 0 {
			continue
		}

		//BUAT NGESEKIP DATA YG TOLOL NIM NYA GA ADA DI TABEL mahasiswa_dinus
		var mahasiswa domain.MahasiswaDinus
		err = db.Where("nim_dinus = ?", record[4]).First(&mahasiswa).Error
		if err != nil {
			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[4], i)
			continue
		}

		TimeIjin, err := time.Parse(layout, record[3])
		if err != nil {
			log.Fatalf("Error parsing date in line %v: %v", i, err)
		}

		ijinKrs := domain.MhsIjinKrs{
			ID:       atoi(record[0]),
			TA:       atoi(record[1]),
			Ijinkan:  record[2] == "1",
			Time:     TimeIjin,
			NimDinus: record[4],
		}

		// Insert ke database
		if err := db.Create(&ijinKrs).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// (AMAN)
func TestHerregistMhs(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/herregist_mahasiswa.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	const layout = "2006-01-02 15:04:05.000"

	for i, record := range records {
		if i == 0 {
			continue
		}

		//BUAT NGESEKIP DATA YG TOLOL NIM NYA GA ADA DI TABEL mahasiswa_dinus
		var mahasiswa domain.MahasiswaDinus
		err = db.Where("nim_dinus = ?", record[1]).First(&mahasiswa).Error
		if err != nil {
			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[1], i)
			continue
		}

		timeReg, err := time.Parse(layout, record[3])
		if err != nil {
			log.Fatalf("Error parsing date in line %v: %v", i, err)
		}

		herregistMahasiswa := domain.HerregistMahasiswa{
			ID:       atoi(record[0]),
			NimDinus: record[1],
			TA:       atoi(record[2]),
			DateReg:  timeReg,
		}

		// Insert ke database
		if err := db.Create(&herregistMahasiswa).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// WORK BUAT DATA EMANG ADA YG TOLOL JADI ADA YG GA MASUK
func TestMahasiswaDiPaketkan(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/mhs_dipaketkan.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {
		if i == 0 {
			continue
		}

		//BUAT NGESEKIP DATA YG TOLOL NIM NYA GA ADA DI TABEL mahasiswa_dinus
		var mahasiswa domain.MahasiswaDinus
		err = db.Where("nim_dinus = ?", record[0]).First(&mahasiswa).Error
		if err != nil {
			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[0], i)
			continue
		}

		dipaketkan := domain.MhsDipaketkan{
			NimDinus:   record[0],
			TaMasukMhs: atoi(record[1]),
		}

		// Insert ke database
		if err := db.Create(&dipaketkan).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// WORK tapi lama nanti kalo sempat pakai goroutine
func TestDaftarNilai(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/daftar_nilai.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {
		if i == 0 {
			continue
		}

		//BUAT NGESEKIP DATA YG TOLOL NIM NYA GA ADA DI TABEL mahasiswa_dinus
		var mahasiswa domain.MahasiswaDinus
		err = db.Where("nim_dinus = ?", record[1]).First(&mahasiswa).Error
		if err != nil {
			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[1], i)
			continue
		}

		daftarNilai := domain.DaftarNilai{
			ID:       atoi(record[0]),
			NimDinus: record[1],
			Kdmk:     record[2],
			Nl:       record[3],
			Hide:     int16(atoi(record[4])),
		}

		// Insert ke database
		if err := db.Create(&daftarNilai).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// WORK
func TestValidasiKrsMhs(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/validasi_krs_mhs.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	const layout = "2006-01-02 15:04:05.000"

	for i, record := range records {
		if i == 0 {
			continue
		}

		//BUAT NGESEKIP DATA YG TOLOL NIM NYA GA ADA DI TABEL mahasiswa_dinus
		var mahasiswa domain.MahasiswaDinus
		err = db.Where("nim_dinus = ?", record[1]).First(&mahasiswa).Error
		if err != nil {
			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[1], i)
			continue
		}

		var kodeExist domain.TahunAjaran
		err = db.Where("kode = ?", record[5]).First(&kodeExist).Error
		if err != nil {
			log.Printf("KODE %s not found in tahun_ajaran, skipping line %v", record[5], i)
			continue
		}

		JobDate, err := time.Parse(layout, record[2])
		if err != nil {
			log.Fatalf("Error parsing date in line %v: %v", i, err)
		}

		validasiKrsMhs := domain.ValidasiKrsMhs{
			ID:       atoi(record[0]),
			NimDinus: record[1],
			JobDate:  JobDate,
			JobHost:  record[3],
			JobAgent: record[4],
			TA:       atoi(record[5]),
		}

		// Insert ke database
		if err := db.Create(&validasiKrsMhs).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// WORK
func TestSesiKuliahBentrok(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/sesi_kuliah_bentrok.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {
		if i == 0 {
			continue
		}

		idSesiKuliah := atoi(record[0])
		idBentrok := atoi(record[1])

		// JIKA ID DAN ID BENTOK SUDAH ADA DI DB SKIP
		var bentrokExist domain.SesiKuliahBentrok
		err := db.Where("id = ? AND id_bentrok = ?", idSesiKuliah, idBentrok).First(&bentrokExist).Error
		if err == nil {
			log.Printf("Sesi kuliah bentrok dengan id %d dan id_bentrok %d sudah ada, skipping line %d", idSesiKuliah, idBentrok, i)
			continue
		}

		//	CEK APAKAH ID SESIMKUALIAH ADA
		var sesiKuliah domain.SesiKuliah
		err = db.Where("id = ?", idSesiKuliah).First(&sesiKuliah).Error
		if err != nil {
			log.Printf("Sesi kuliah dengan id %d tidak ditemukan, skipping line %d", idSesiKuliah, i)
			continue
		}

		// CEK APAKAH IDKULIAHBENTROK ADA DI SESI KULIAH
		var sesiKuliahBentrok domain.SesiKuliah
		err = db.Where("id = ?", idBentrok).First(&sesiKuliahBentrok).Error
		if err != nil {
			// Jika id_bentrok tidak ditemukan di sesi_kuliah, lewati baris ini
			log.Printf("ID Bentrok %d tidak ditemukan di tabel sesi_kuliah, skipping line %d", idBentrok, i)
			continue
		}

		bentrok := domain.SesiKuliahBentrok{
			ID:        idSesiKuliah,
			IDBentrok: idBentrok,
		}

		// Insert ke database
		if err := db.Create(&bentrok).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

/*
	&domain.JadwalTawar{},
	&domain.KrsRecord{},
	&domain.KrsRecordLog{},
*/

/*
DATA TOLOL MANA ADA ID DAN RUANG 0 OK KITA BUAT DI BOLEH NULL AJA
OK NOW WORK
*/
func TestJadwalTawar(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/jadwal_tawar.csv")
	IfErrNotNil(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	IfErrNotNil(err)

	for i, record := range records {
		if i == 0 {
			continue
		}

		// JIKA HARI/SESI/RUANG == 0 DI ISI null
		ifZeroRetturnNul := func(x string) *int {
			if x == "" || x == "0" {
				return nil // Return nil jika string kosong atau "0"
			}
			result := atoi(x)
			return &result // Return pointer dari hasil konversi
		}
		// PENGECEKAN KETIKA RUANG HARI / RUANG / SESI / TIDAK EXIST
		SkipKah := false
		var HariExist domain.Hari
		ExistHarikah := func(hari string, db2 *gorm.DB) string {
			if hari == "0" || hari == "" {
				log.Println("GA PERLU DI CEK")
			} else {
				if err = db2.Where("id = ?", hari).First(&HariExist).Error; err != nil {
					SkipKah = true
				}
			}
			return "0"
		}
		ExistHarikah(record[9], db)
		ExistHarikah(record[11], db)
		ExistHarikah(record[12], db)

		var RuangExist domain.Ruang
		ExistRuangKah := func(ruang string, db2 *gorm.DB) string {
			if ruang == "0" || ruang == "" {
				log.Println("GA PERLU DI CEK")
			} else {
				if err = db2.Where("id = ?", ruang).First(&RuangExist).Error; err != nil {
					SkipKah = true
				}
			}
			return "0"
		}
		ExistRuangKah(record[15], db)
		ExistRuangKah(record[16], db)
		ExistRuangKah(record[17], db)

		var SesiExist domain.SesiKuliah
		ExistSesi := func(sesi string, db2 *gorm.DB) string {
			if sesi == "0" || sesi == "" {
				log.Println("GA PERLU DI CEK")
			} else {
				if err = db2.Where("id = ?", sesi).First(&SesiExist).Error; err != nil {
					SkipKah = true
				}
			}
			return "0"
		}
		ExistSesi(record[12], db)
		ExistSesi(record[13], db)
		ExistSesi(record[14], db)

		if SkipKah == true {
			log.Println("ADA LINE YG DI SKIP")
			continue
		}

		tawar := domain.JadwalTawar{
			ID:        atoi(record[0]),
			TA:        atoi(record[1]),
			Kdmk:      record[2],
			Klpk:      record[3],
			Klpk2:     record[4],
			Kdds:      atoi(record[5]),
			Kdds2:     atoi(record[6]),
			Jmax:      atoi(record[7]),
			Jsisa:     atoi(record[8]),
			IDHari1:   int8(atoi(record[9])),
			IDHari2:   int8(atoi(record[10])),
			IDHari3:   int8(atoi(record[11])),
			IDSesi1:   ifZeroRetturnNul(record[12]),
			IDSesi2:   ifZeroRetturnNul(record[13]),
			IDSesi3:   ifZeroRetturnNul(record[14]),
			IDRuang1:  ifZeroRetturnNul(record[15]),
			IDRuang2:  ifZeroRetturnNul(record[16]),
			IDRuang3:  ifZeroRetturnNul(record[17]),
			JnsJam:    int8(atoi(record[18])),
			OpenClass: record[19] == "1",
		}

		// Insert ke database
		if err := db.Create(&tawar).Error; err != nil {
			log.Fatalf("Error in line %v: %v", i, err)
		}

		log.Println("insert index ", i)
	}

	log.Println("YEY SUCCESS")
}

// test ualang pake go routine
func TestKrsRecord(t *testing.T) {
	db := conf.InitDB()
	file, err := os.Open("data_krs/krs_record.csv")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	var wg sync.WaitGroup          // Membuat WaitGroup untuk menunggu semua goroutine selesai
	sem := make(chan struct{}, 50) // Batas maksimal 5 goroutine concurrent untuk menghindari overload DB

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		wg.Add(1)
		sem <- struct{}{} // Mengambil slot goroutine

		go func(i int, record []string) {
			defer func() {
				<-sem // Melepaskan slot goroutine setelah selesai
				wg.Done()
			}()

			// CHECK NIM_DINUS
			var mahasiswa domain.MahasiswaDinus
			err = db.Where("nim_dinus = ?", record[4]).First(&mahasiswa).Error
			if err != nil {
				log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[4], i)
				return
			}

			// CHECK TA
			var kodeExist domain.TahunAjaran
			err = db.Where("kode = ?", record[1]).First(&kodeExist).Error
			if err != nil {
				log.Printf("KODE %s not found in tahun_ajaran, skipping line %v", record[1], i)
				return
			}

			var kdmkExist domain.MatkulKurikulum
			if err := db.Where("kdmk = ?", record[2]).First(&kdmkExist).Error; err != nil {
				log.Printf("kdmk %s not found in matkul_kurikulum, skipping line %v", record[2], i)
				return
			}

			var idJadwalExist domain.JadwalTawar
			if err := db.Where("id = ?", record[3]).First(&idJadwalExist).Error; err != nil {
				log.Printf("KODE %s not found in jadwal_tawar, skipping line %v", record[3], i)
				return
			}

			// Mengecek apakah data KrsRecord sudah ada
			var KrsIdExist domain.KrsRecord
			if err := db.Where("id = ?", record[0]).First(&KrsIdExist).Error; err == nil {
				log.Printf("Krs %s already exists, skipping line %v", record[0], i)
				return
			}

			// Membuat objek krsRecord
			krsRecord := domain.KrsRecord{
				ID:       atoi(record[0]),
				TA:       atoi(record[1]),
				Kdmk:     record[2],
				IDJadwal: atoi(record[3]),
				NimDinus: record[4],
				Sts:      record[5],
				Sks:      atoi(record[6]),
				Modul:    atoi(record[7]),
			}

			// Memulai transaksi untuk insert data
			tx := db.Begin()
			if err := tx.Create(&krsRecord).Error; err != nil {
				tx.Rollback() // Rollback jika ada error
				log.Printf("Error inserting data at index %d: %v", i, err)
				return
			}

			tx.Commit()
			log.Printf("Insert successful at index %d", i)
		}(i, record)
	}

	wg.Wait() // Menunggu semua goroutine selesai
	log.Println("YEY SUCCESS")
}

//func TestKrsRecordLog(t *testing.T) {
//	db := conf.InitDB()
//	file, err := os.Open("data_krs/krs_record_log.csv")
//	IfErrNotNil(err)
//	defer file.Close()
//
//	reader := csv.NewReader(file)
//	records, err := reader.ReadAll()
//	IfErrNotNil(err)
//
//	const layout = "2006-01-02 15:04:05.000"
//
//	for i, record := range records {
//		if i == 0 {
//			continue
//		}
//
//		//CHECK NIM_DINUS
//		var mahasiswa domain.MahasiswaDinus
//		err = db.Where("nim_dinus = ?", record[5]).First(&mahasiswa).Error
//		if err != nil {
//			log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[5], i)
//			continue
//		}
//
//		var kdmkExist domain.MatkulKurikulum
//		if err := db.Where("kdmk = ?", record[1]).First(&kdmkExist).Error; err != nil {
//			log.Printf("kdmk %s not found in matkul_kurikulum, skipping line %v", record[1], i)
//			continue
//		}
//
//		var idKrsExist domain.KrsRecord
//		if err := db.Where("id = ?", record[0]).First(&idKrsExist).Error; err != nil {
//			log.Printf("KrsID %s not found in tahun_ajaran, skipping line %v", record[0], i)
//			continue
//		}
//
//		var KrsIdExist domain.KrsRecordLog
//		if err := db.Where("id = ?", record[0]).First(&KrsIdExist).Error; err != nil {
//			log.Printf("KODE %s not found in tahun_ajaran, skipping line %v", record[0], i)
//			continue
//		}
//
//		LastUpdate, err := time.Parse(layout, record[4])
//		if err != nil {
//			log.Fatalf("Error parsing date in line %v: %v", i, err)
//		}
//
//		recordLog := domain.KrsRecordLog{
//			IDKrs:      atoi(record[0]), // validasi dulu apkah ada
//			Kdmk:       record[1],       // validasi dulu apakh ada
//			Aksi:       int8(atoi(record[2])),
//			IDJadwal:   atoi(record[3]),
//			LastUpdate: LastUpdate,
//			NimDinus:   record[5], // validasi dulu apakh ada
//		}
//
//		// Insert ke database
//		if err := db.Create(&recordLog).Error; err != nil {
//			log.Fatalf("Error in line %v: %v", i, err)
//		}
//
//		log.Println("insert index ", i)
//	}
//
//	log.Println("YEY SUCCESS")
//}

func TestKrsRecordLog(t *testing.T) {
	db := conf.InitDB()

	// Membuka file CSV
	file, err := os.Open("data_krs/krs_record_log.csv")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Membaca seluruh isi file CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	const layout = "2006-01-02 15:04:05.000"

	// Membatasi goroutine menjadi 50
	var wg sync.WaitGroup
	sem := make(chan struct{}, 50) // Channel untuk membatasi jumlah goroutine

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		// Mulai goroutine
		wg.Add(1)
		sem <- struct{}{} // Mengambil slot di channel (menunggu jika lebih dari 50)

		go func(i int, record []string) {
			defer wg.Done()
			defer func() { <-sem }() // Membebaskan slot setelah goroutine selesai

			// Validasi NIM_DINUS
			var mahasiswa domain.MahasiswaDinus
			err := db.Where("nim_dinus = ?", record[5]).First(&mahasiswa).Error
			if err != nil {
				log.Printf("NIM %s not found in mahasiswa_dinus, skipping line %v", record[5], i)
				return
			}

			// Validasi KDMK
			var kdmkExist domain.MatkulKurikulum
			if err := db.Where("kdmk = ?", record[1]).First(&kdmkExist).Error; err != nil {
				log.Printf("kdmk %s not found in matkul_kurikulum, skipping line %v", record[1], i)
				return
			}

			// Validasi ID KRS
			var idKrsExist domain.KrsRecord
			if err := db.Where("id = ?", record[0]).First(&idKrsExist).Error; err != nil {
				log.Printf("KrsID %s not found in tahun_ajaran, skipping line %v", record[0], i)
				return
			}

			// Cek apakah KrsRecordLog sudah ada
			var KrsIdExist domain.KrsRecordLog
			if err := db.Where("id = ?", record[0]).First(&KrsIdExist).Error; err == nil {
				log.Printf("KrsRecordLog %s already exists, skipping line %v", record[0], i)
				return
			}

			// Parsing LastUpdate
			LastUpdate, err := time.Parse(layout, record[4])
			if err != nil {
				log.Printf("Error parsing date in line %v: %v", i, err)
				return
			}

			// Membuat record log
			recordLog := domain.KrsRecordLog{
				IDKrs:      atoi(record[0]),
				Kdmk:       record[1],
				Aksi:       int8(atoi(record[2])),
				IDJadwal:   atoi(record[3]),
				LastUpdate: LastUpdate,
				NimDinus:   record[5],
			}

			// Insert ke database
			if err := db.Create(&recordLog).Error; err != nil {
				log.Printf("Error inserting record at line %v: %v", i, err)
				return
			}

			log.Printf("Inserted record at index %v", i)
		}(i, record)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	log.Println("YEY SUCCESS")
}
