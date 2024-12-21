package main

import (
	"context"
	"fmt"
	"koriebruh/try/conf"
	"koriebruh/try/domain"
	"log"
	"testing"
)

var ctx = context.Background()

func TestFindExist(t *testing.T) {
	db := conf.InitDB()

	nimDinus := "262019ecd15e0169f7efdea9a64ad30e"
	mahasiswaDinus := domain.MahasiswaDinus{}
	if err := db.WithContext(ctx).Where("nim_dinus =?", nimDinus).First(&mahasiswaDinus).Error; err != nil {
		log.Fatalf("err find")
	}

	fmt.Println("yey ")

}

func TestQueryStatusMhs(t *testing.T) {
	db := conf.InitDB()
	var results []struct {
		NimDinus     string
		TaMasuk      int
		Prodi        string
		AkdmStat     int
		DateReg      string
		SppBayar     int
		SppStatus    int
		SppTransaksi string
		Kelas        string
	}

	nim := "262019ecd15e0169f7efdea9a64ad30e"

	err := db.Model(&domain.MahasiswaDinus{}).
		Select("mahasiswa_dinus.nim_dinus, mahasiswa_dinus.ta_masuk, mahasiswa_dinus.prodi, mahasiswa_dinus.akdm_stat, mahasiswa_dinus.kelas, herregist_mahasiswa.date_reg, tagihan_mhs.spp_bayar, tagihan_mhs.spp_status, tagihan_mhs.spp_transaksi").
		Joins("JOIN herregist_mahasiswa ON mahasiswa_dinus.nim_dinus = herregist_mahasiswa.nim_dinus").
		Joins("JOIN krs_management.tagihan_mhs ON herregist_mahasiswa.nim_dinus = tagihan_mhs.nim_dinus").
		Where("mahasiswa_dinus.nim_dinus = ?", nim).
		Scan(&results).Error

	if err != nil {
		log.Fatal(err)
	}

	// Print results
	for _, result := range results {
		fmt.Printf(
			"NimDinus: %s, TA Masuk: %d, Prodi: %s, Status Akademik: %v, DateReg: %s, SPP Bayar: %v, SPP Status: %v, SPP Transaksi: %s, Kelas%v \n",
			result.NimDinus,
			result.TaMasuk,
			result.Prodi,
			result.AkdmStat,
			result.DateReg,
			result.SppBayar,
			result.SppStatus,
			result.SppTransaksi,
			result.Kelas,
		)
	}

}

func TestQueryFindALlKrsPicked(t *testing.T) {
	db := conf.InitDB()

	var results []struct {
		NamaMatkul   string
		NamaMatkulEN string
		Tipe         string
		Semester     int
		JenisMatkul  string
		Hari1        string
		Hari2        string
		Hari3        string
	}

	nimDinus := "1a4421a533b58bb95212ca38610c13de"
	err := db.Model(&domain.KrsRecord{}).
		Select("matkul_kurikulum.nmmk AS nama_matkul, matkul_kurikulum.nmen AS nama_matkul_en, matkul_kurikulum.tp AS tipe, matkul_kurikulum.smt AS semester, matkul_kurikulum.jenis_matkul AS jenis_matkul, hari1.nama AS hari1, hari2.nama AS hari2, hari3.nama AS hari3").
		Joins("JOIN matkul_kurikulum ON matkul_kurikulum.kdmk = krs_record.kdmk").
		Joins("JOIN jadwal_tawar ON jadwal_tawar.id = krs_record.id_jadwal").
		Joins("LEFT JOIN hari AS hari1 ON hari1.id = jadwal_tawar.id_hari1").
		Joins("LEFT JOIN hari AS hari2 ON hari2.id = jadwal_tawar.id_hari2").
		Joins("LEFT JOIN hari AS hari3 ON hari3.id = jadwal_tawar.id_hari3").
		Where("krs_record.nim_dinus = ?", nimDinus).
		Scan(&results).Error

	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Printf("NamaMatkul: %s, Hari1: %s, Hari2: %s, Hari3: %s\n",
			result.NamaMatkul, result.Hari1, result.Hari2, result.Hari3)
	}
}

func TestStatusKRS(t *testing.T) {
	db := conf.InitDB()
	ctx := context.Background()
	nimDinus := "33f5cda80c5f2fc0bdb865cceb51550c"

	// Struct untuk menyimpan hasil akhir
	var Data struct {
		Validate    string
		TahunAjaran string
		Dipaketkan  string
		TahunMasuk  string
		Sks         int
		Ips         string
	}

	var validasi string
	err := db.WithContext(ctx).Model(&domain.ValidasiKrsMhs{}).
		Select("CASE WHEN job_date <= NOW() THEN 'Validated' ELSE 'Not Validated' END AS validation_status").
		Where("nim_dinus = ?", nimDinus).
		First(&validasi).Error
	if err != nil {
		log.Fatalf("Error fetching validation status: %v", err)
	}

	// Query apakah mahasiswa di-paketkan
	var diPaketkanKah domain.MhsDipaketkan
	if err = db.WithContext(ctx).Where("nim_dinus = ?", nimDinus).First(&diPaketkanKah).Error; err != nil {
		Data.Dipaketkan = "tidak di paketkan"
	} else {
		Data.Dipaketkan = "dipaketkan"
	}

	// Query informasi IP semester dan tahun masuk
	var ipSemester struct {
		TahunAjaran string
		Sks         int
		Ips         string
		TahunMasuk  string
	}
	err = db.WithContext(ctx).Raw(`
		SELECT
			ip_s.ta AS tahun_ajaran,
			ip_s.sks,
			ip_s.ips,
			md.ta_masuk AS tahun_masuk
		FROM ip_semester ip_s
		JOIN mahasiswa_dinus md ON ip_s.nim_dinus = md.nim_dinus
		WHERE ip_s.nim_dinus = ?
		LIMIT 1
	`, nimDinus).Scan(&ipSemester).Error
	if err != nil {
		log.Fatalf("Error fetching IP semester: %v", err)
	}

	// Mengisi data akhir
	Data.Validate = validasi
	Data.TahunAjaran = ipSemester.TahunAjaran
	Data.TahunMasuk = ipSemester.TahunMasuk
	Data.Sks = ipSemester.Sks
	Data.Ips = ipSemester.Ips

	// Output hasil akhir
	fmt.Printf("Final Data: %+v\n", Data)
}
