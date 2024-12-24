package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/try/conf"
	"koriebruh/try/domain"
	"koriebruh/try/dto"
	"koriebruh/try/repository"
	"log"
	"strconv"
	"strings"
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

type JadwalTawaran struct {
	TahunAjaran int    `gorm:"column:tahun_ajaran"`
	Kelompok    string `gorm:"column:kelompok"`
	Matakuliah  string `gorm:"column:nama_mata_kuliah"`
	Sks         int    `gorm:"column:jumlah_sks"`
	Hari        string `gorm:"column:hari"`
	JamMulai    string `gorm:"column:jam_mulai"`
	JamSelesai  string `gorm:"column:jam_selesai"`
	Ruang       string `gorm:"column:ruang"`
}

func TestQueryTawaran(t *testing.T) {
	db := conf.InitDB()

	var result []JadwalTawaran
	tahunAjaran := "20232"

	err := db.Raw(`
		SELECT
			jt.ta AS tahun_ajaran,
			jt.klpk AS kelompok,
			mk.nmmk AS nama_mata_kuliah,
			mk.sks AS jumlah_sks,
			h.nama AS hari,
			sk.jam_mulai,
			sk.jam_selesai,
			r.nama AS ruang
		FROM jadwal_tawar jt
			JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
			JOIN hari h ON jt.id_hari1 = h.id
			JOIN sesi_kuliah sk ON jt.id_sesi1 = sk.id
			JOIN ruang r ON jt.id_ruang1 = r.id
		WHERE
			mk.kur_aktif = 1 AND
			jt.ta = ?;
	`, tahunAjaran).Scan(&result).Error

	if err != nil {
		t.Fatalf("Gagal menjalankan query: %v", err)
	}

	// Debug output hasil query
	for _, row := range result {
		fmt.Printf("Data: %+v\n", row)
	}
}

func TestJadwalInput(t *testing.T) {
	db := conf.InitDB()
	var user domain.MahasiswaDinus
	nimDinus := "1a4421a533b58bb95212ca38610c13de"
	if err := db.Where("nim_dinus = ?", nimDinus).First(&user).Error; err != nil {
		panic(err)
	}

	var jadwal domain.JadwalInputKrs
	if err := db.Where("prodi = ?", user.Prodi).First(&jadwal).Error; err != nil {
		panic(err)
	}

	result := struct {
		TA         int
		Prodi      string
		TglMulai   string
		TglSelesai string
	}{
		TA:         jadwal.TA,
		Prodi:      jadwal.Prodi,
		TglMulai:   jadwal.TglMulai.Format("2006-01-02 15:04:05"),
		TglSelesai: jadwal.TglSelesai.Format("2006-01-02 15:04:05"),
	}

	fmt.Println(result)
}

func TestScheduleByProdi(t *testing.T) {

	db := conf.InitDB()

	var results []struct {
		TahunAjaran     string `gorm:"column:tahun_ajaran"`
		KodeMataKuliah  string `gorm:"column:kode_mata_kuliah"`
		Kelompok        string `gorm:"column:kelompok"`
		NamaMataKuliah  string `gorm:"column:nama_mata_kuliah"`
		JumlahSKS       int    `gorm:"column:jumlah_sks"`
		Hari            string `gorm:"column:hari"`
		JamMulai        string `gorm:"column:jam_mulai"` // Perbaikan di sini, sebelumnya ada typo `:=`
		JamSelesai      string `gorm:"column:jam_selesai"`
		Ruang           string `gorm:"column:ruang"`
		StatusPemilihan string `gorm:"column:status_pemilihan"`
	}

	nimDinus := "647e27c32c8935273e876a457b81b186"
	kelompok := "B11%"

	query := `
        SELECT DISTINCT
            jt.ta AS tahun_ajaran,
            jt.kdmk AS kode_mata_kuliah,
            jt.klpk AS kelompok,
            mk.nmmk AS nama_mata_kuliah,
            mk.sks AS jumlah_sks,
            h.nama AS hari,
            sk.jam_mulai,
            sk.jam_selesai,
            r.nama AS ruang,
            CASE
                WHEN EXISTS (
                    SELECT 1
                    FROM daftar_nilai dn
                    WHERE dn.kdmk = jt.kdmk AND dn.nl = 'A' AND dn.nim_dinus = ?
                ) THEN 'Tidak Bisa'
                ELSE 'Bisa'
            END AS status_pemilihan
        FROM
            jadwal_tawar jt
            JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
            JOIN hari h ON jt.id_hari1 = h.id
            JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
            JOIN ruang r ON jt.id_ruang1 = r.id
        WHERE
            jt.ta IS NOT NULL
            AND jt.klpk LIKE ?
        ORDER BY
            jt.ta, mk.nmmk;
    `

	if err := db.Raw(query, nimDinus, kelompok).Scan(&results).Error; err != nil {
		panic(err)
	}

	fmt.Println(results)

}

func TestQueryGetAllScores(t *testing.T) {

	db := conf.InitDB()

	var result []struct {
		KodeMatkul  string `gorm:"column:kode_matkul" json:"kode_matkul"`
		MataKuliah  string `gorm:"column:mata_kuliah" json:"mata_kuliah"`
		Sks         int    `gorm:"column:sks" json:"sks"`
		Category    string `gorm:"column:category" json:"category"`
		JenisMatkul string `gorm:"column:jenis_matkul" json:"jenis_matkul"`
		Nilai       string `gorm:"column:nilai" json:"nilai"`
	}

	nimDinus := "6f41ddf2e566f37089dd0e2f5fdbeca1"

	err := db.WithContext(ctx).Table("daftar_nilai dn").
		Select(`
			mk.kdmk AS kode_matkul,
			mk.nmen AS matakuliah,
			mk.sks AS sks,
			mk.tp AS category,
			mk.jenis_matkul AS jenis_matkul,
			dn.nl AS nilai
		`).
		Joins("JOIN matkul_kurikulum mk ON dn.kdmk = mk.kdmk").
		Where("dn.nim_dinus = ? AND dn.hide = 0", nimDinus).
		Scan(&result).Error

	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatalf("nilai not found bg")
	}

	for _, row := range result {
		t.Logf("Kode Matkul: %s, Mata Kuliah: %s, SKS: %d, Nilai: %s", row.KodeMatkul, row.MataKuliah, row.Sks, row.Nilai)
	}
}

func TestCheckScheduleConflict(t *testing.T) {

	db := conf.InitDB()

	var schedules []struct {
		TahunAjaran    string `gorm:"column:tahun_ajaran" json:"tahun_ajaran"`
		Kelompok       string `gorm:"column:kelompok" json:"kelompok"`
		NamaMataKuliah string `gorm:"column:nama_mata_kuliah" json:"nama_mata_kuliah"`
		JumlahSKS      int    `gorm:"column:jumlah_sks" json:"jumlah_sks"`
		Hari           string `gorm:"column:hari" json:"hari"`
		JamMulai       string `gorm:"column:jam_mulai" json:"jam_mulai"`
		JamSelesai     string `gorm:"column:jam_selesai" json:"jam_selesai"`
		Ruang          string `gorm:"column:ruang" json:"ruang"`
		StatusBentrok  string `gorm:"column:status_bentrok" json:"status_bentrok"`
		KeteranganSlot string `gorm:"column:keterangan_slot" json:"keterangan_slot"`
	}

	nimDinus := "6f41ddf2e566f37089dd0e2f5fdbeca1"
	kodeTA := "20232"

	query := `
        SELECT
            jt.ta AS tahun_ajaran,
            jt.klpk AS kelompok,
            mk.nmmk AS nama_mata_kuliah,
            mk.sks AS jumlah_sks,
            h.nama AS hari,
            sk.jam_mulai,
            sk.jam_selesai,
            r.nama AS ruang,
            CASE
                WHEN EXISTS (
                    SELECT 1
                    FROM krs_record kr
                             JOIN jadwal_tawar jt_inner ON kr.id_jadwal = jt_inner.id
                             JOIN sesi_kuliah sk_inner ON sk_inner.id = jt_inner.id_sesi1
                    WHERE kr.nim_dinus = ?
                      AND jt.id_hari1 = jt_inner.id_hari1
                      AND (
                        (sk.jam_mulai < sk_inner.jam_selesai AND sk.jam_selesai > sk_inner.jam_mulai)
                        )
                ) THEN 'BENTROK'
                ELSE NULL
            END AS status_bentrok,
            CASE
                WHEN jt.jsisa = jt.jmax THEN CONCAT(jt.jsisa, '/', jt.jmax, ' SLOT PENUH')
                ELSE CONCAT(jt.jsisa, '/', jt.jmax)
            END AS keterangan_slot
        FROM jadwal_tawar jt
                 JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
                 JOIN hari h ON jt.id_hari1 = h.id
                 JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
                 JOIN ruang r ON jt.id_ruang1 = r.id
        WHERE mk.kur_aktif = 1
          AND jt.ta = ?
          AND jt.jns_jam IN (1, 2)
          AND jt.jsisa <= jt.jmax;
    `

	// Jalankan raw query
	err := db.Raw(query, nimDinus, kodeTA).Scan(&schedules).Error
	if err != nil {
		fmt.Println("Query execution failed:", err)
	}

	for _, schedule := range schedules {
		fmt.Printf("%+v\n", schedule)
	}

}

func TestInsertSchedule(t *testing.T) {

	nim := "560b4d78fc163d57774e045317be842f"
	kodeTA := "275486"
	idSchduleUserInsert := 283180 //277345
	prodi := "A11"

	db := conf.InitDB()
	statusRepository := repository.NewStudentStatusRepository()

	err := db.Transaction(func(tx *gorm.DB) error {

		//Insert jadwal mata kuliah di luar kurikulum ditolak
		_, err := statusRepository.InsertKRSPermit(ctx, tx, nim) // if != err
		if err != nil {
			panic("TIDAK PUNYA IJIN INSERT") // berarti tidak punya ijin insert krs
		}

		//KRS yang sudah tervalidasi tidak dapat insert lagi
		statusKRS, _ := statusRepository.StatusKRS(ctx, tx, nim)
		if statusKRS.Validate == "Validated" {
			panic("SUDAH DI VALIDASI TIDAK BISA INSERT LAGI")
		}

		conflictsSchedule, err := statusRepository.ScheduleConflicts(ctx, tx, nim, kodeTA)
		var foundSchedule dto.ScheduleConflictRes
		for _, s := range conflictsSchedule {
			// Mengecek jika id pada elemen s sama dengan ischedule
			if s.Id == idSchduleUserInsert {
				// Menyimpan data jadwal yang sesuai dengan ischedule
				foundSchedule = dto.ScheduleConflictRes{
					Id:             s.Id,
					TahunAjaran:    s.TahunAjaran,
					Kelompok:       s.Kelompok,
					NamaMataKuliah: s.NamaMataKuliah,
					JumlahSKS:      s.JumlahSKS,
					Hari:           s.Hari,
					JamMulai:       s.JamMulai,
					JamSelesai:     s.JamSelesai,
					Ruang:          s.Ruang,
					StatusBentrok:  s.StatusBentrok,
					KeteranganSlot: s.KeteranganSlot,
				}
				break // Hentikan loop jika sudah menemukan data yang sesuai
			}
		}

		if foundSchedule.StatusBentrok == "BENTROK" {
			panic("BENTROK COK")
		} else if strings.Contains(foundSchedule.KeteranganSlot, "SLOT PENUH") {
			panic("SLOT PENUH")
		}

		offersProdi, err := statusRepository.KrsOffersProdi(ctx, tx, nim, kodeTA, prodi)
		var foundOfferProdi dto.KrsOffersProdiResponse
		for _, p := range offersProdi {
			if p.Id == idSchduleUserInsert {
				foundOfferProdi = dto.KrsOffersProdiResponse{
					Id:              p.Id,
					TahunAjaran:     p.TahunAjaran,
					KodeMataKuliah:  p.KodeMataKuliah,
					Kelompok:        p.Kelompok,
					NamaMataKuliah:  p.NamaMataKuliah,
					JumlahSKS:       p.JumlahSKS,
					Hari:            p.Hari,
					JamMulai:        p.JamMulai,
					JamSelesai:      p.JamSelesai,
					Ruang:           p.Ruang,
					StatusPemilihan: p.StatusPemilihan,
					StatusKrs:       p.StatusKrs,
				}
			}
		}

		if foundOfferProdi.StatusPemilihan == "Tidak Bisa" {
			panic("tidak bisa anda sudah dapat A")
		} else if foundOfferProdi.StatusKrs == "Tidak Mencukupi" {
			panic("krs tidak cukup")
		}

		existStudent, err := statusRepository.CheckUserExist(ctx, tx, nim)
		if err != nil {
			panic("SATATUS GA KETEMU")
		}

		var jnsJam string
		err = tx.WithContext(ctx).
			Model(domain.JadwalTawar{}).
			Select("jns_jam").
			Where("id = ?", idSchduleUserInsert).
			Scan(&jnsJam).Error

		if err != nil {
			panic(err) //jenis jam ga ketemu
		}

		fmt.Println(existStudent.Kelas)
		fmt.Println(jnsJam)

		if strconv.Itoa(existStudent.Kelas) != jnsJam {
			panic("JADWAL YG DI AMBIL TIDAK SESUAI PILIHAN KELAS")
		}

		//INSERT
		TA, _ := strconv.Atoi(kodeTA)
		record := domain.KrsRecord{
			TA:       TA,
			Kdmk:     foundOfferProdi.KodeMataKuliah,
			IDJadwal: foundSchedule.Id,
			NimDinus: nim,
			Sts:      "B", // ga paham maksud nya
			Sks:      foundSchedule.JumlahSKS,
			Modul:    0,
		}
		if err = tx.WithContext(ctx).Create(&record).Error; err != nil {
			panic("ERROR INSERT ADA YG KOSONG NI FIELDN NYA")
		}

		recordLog := domain.KrsRecordLog{
			IDKrs:    record.ID,
			NimDinus: nim,
			Kdmk:     foundOfferProdi.KodeMataKuliah,
			Aksi:     1,
			IDJadwal: foundSchedule.Id,
			IpAddr:   "",
		}
		if err = tx.WithContext(ctx).Create(&recordLog).Error; err != nil {
			panic("ERROR INSERT LOG")
		}

		// KURANGI J SISA

		err = tx.WithContext(ctx).
			Model(&domain.JadwalTawar{}).
			Where("id = ?", idSchduleUserInsert).
			UpdateColumn("jsisa", gorm.Expr("jsisa + ?", 1)).Error

		if err != nil {
			panic("Gagal mengupdate jsisa")
		}

		return err
	})

	if err != nil {
		panic(err)
	}
}
