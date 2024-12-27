package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/try/domain"
	"koriebruh/try/dto"
)

type StudentStatusRepository interface {
	KrsOffers(ctx context.Context, db *gorm.DB, kodeTA string) ([]dto.KrsOfferRes, error)
	KrsSchedule(ctx context.Context, db *gorm.DB, prodi string) (domain.JadwalInputKrs, error)
	CheckUserExist(ctx context.Context, db *gorm.DB, nimDinus string) (domain.MahasiswaDinus, error)
	InformationStudent(ctx context.Context, db *gorm.DB, nimDinus string) (*dto.InfoStudentDB, error)
	SetClassTime(ctx context.Context, db *gorm.DB, nimDinus string, classOption int) error
	GetAllKRSPick(ctx context.Context, db *gorm.DB, nimDinus string) ([]dto.SelectedKrs, error)
	InsertKRSPermit(ctx context.Context, db *gorm.DB, nimDinus string) (bool, error)
	StatusKRS(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string) (dto.StatusKrsRes, error)
	KrsOffersProdi(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string, kelompok string) ([]dto.KrsOffersProdiResponse, error)
	GetAllScores(ctx context.Context, db *gorm.DB, nimDinus string) ([]dto.AllScoresRes, error)
	ScheduleConflicts(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string) ([]dto.ScheduleConflictRes, error)
	InsertKrsLog(ctx context.Context, db *gorm.DB, nimDinus string, rec domain.KrsRecord, Aksi int8) error
	InsertKrs(ctx context.Context, db *gorm.DB, rec domain.KrsRecord) error
	GetKrsLog(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string) ([]dto.KrsLogRes, error)
	UpdateValidate(ctx context.Context, db *gorm.DB, mhs domain.ValidasiKrsMhs) error
	CheckTA(ctx context.Context, db *gorm.DB, kodeTA string) error
}

type StudentStatusRepositoryImpl struct {
}

func NewStudentStatusRepository() *StudentStatusRepositoryImpl {
	return &StudentStatusRepositoryImpl{}
}

func (s StudentStatusRepositoryImpl) KrsOffers(ctx context.Context, db *gorm.DB, kodeTA string) ([]dto.KrsOfferRes, error) {
	var krsOffers []dto.KrsOfferRes

	err := db.WithContext(ctx).Raw(`
		SELECT DISTINCT
		    jt.id AS id,
			jt.ta AS tahun_ajaran,
			jt.klpk AS kelompok,
			mk.nmmk AS nama_mata_kuliah,
			mk.sks AS jumlah_sks,
			h.nama AS hari,
			sk.jam_mulai,
			sk.jam_selesai,
			jt.jns_jam AS jns_jam,
			r.nama AS ruang
		FROM jadwal_tawar jt
			LEFT JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
			LEFT JOIN hari h ON jt.id_hari1 = h.id
			LEFT JOIN sesi_kuliah sk ON jt.id_sesi1 = sk.id
			LEFT JOIN ruang r ON jt.id_ruang1 = r.id
		WHERE
			mk.kur_aktif = 1 AND
			jt.ta = ?;
	`, kodeTA).Scan(&krsOffers).Error

	if err != nil {
		return nil, fmt.Errorf("error kode Tahun Ajar %v not found", kodeTA)
	}

	return krsOffers, nil
}

func (s StudentStatusRepositoryImpl) KrsSchedule(ctx context.Context, db *gorm.DB, prodi string) (domain.JadwalInputKrs, error) {
	var jadwal domain.JadwalInputKrs
	if err := db.Where("prodi = ?", prodi).First(&jadwal).Error; err != nil {
		return domain.JadwalInputKrs{}, fmt.Errorf("failed to get jadwal where prodi %v", prodi)
	}

	return jadwal, nil
}

func (s StudentStatusRepositoryImpl) CheckUserExist(ctx context.Context, db *gorm.DB, nimDinus string) (domain.MahasiswaDinus, error) {
	mahasiswaDinus := domain.MahasiswaDinus{}
	if err := db.WithContext(ctx).Where("nim_dinus =?", nimDinus).First(&mahasiswaDinus).Error; err != nil {
		return domain.MahasiswaDinus{}, fmt.Errorf("failed to get status user %v", nimDinus)
	}

	return mahasiswaDinus, nil
}

func (s StudentStatusRepositoryImpl) InformationStudent(ctx context.Context, db *gorm.DB, nimDinus string) (*dto.InfoStudentDB, error) {
	var studentStatus dto.InfoStudentDB

	err := db.WithContext(ctx).Model(&domain.MahasiswaDinus{}).
		Select("mahasiswa_dinus.nim_dinus, mahasiswa_dinus.ta_masuk, mahasiswa_dinus.prodi, mahasiswa_dinus.akdm_stat, mahasiswa_dinus.kelas, herregist_mahasiswa.date_reg, tagihan_mhs.spp_bayar, tagihan_mhs.spp_status, tagihan_mhs.spp_transaksi").
		Joins("JOIN herregist_mahasiswa ON mahasiswa_dinus.nim_dinus = herregist_mahasiswa.nim_dinus").
		Joins("JOIN tagihan_mhs ON herregist_mahasiswa.nim_dinus = tagihan_mhs.nim_dinus").
		Where("mahasiswa_dinus.nim_dinus = ?", nimDinus).
		Scan(&studentStatus).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get status user %v, %e", nimDinus, err)
	}

	fmt.Println(studentStatus)
	return &studentStatus, nil
}

func (s StudentStatusRepositoryImpl) SetClassTime(ctx context.Context, db *gorm.DB, nimDinus string, classOption int) error {

	var CountKrsInsert int64

	if err := db.WithContext(ctx).Model(&domain.KrsRecord{}).
		Where("nim_dinus = ?", nimDinus).
		Count(&CountKrsInsert).Error; err != nil {
		return fmt.Errorf("failed to check KRS record for nim_dinus=%s: %w", nimDinus, err)
	}

	if CountKrsInsert > 0 {
		return fmt.Errorf("you have added a total of %d Krs, you can't change the class type", CountKrsInsert)
	}

	if err := db.WithContext(ctx).Model(&domain.MahasiswaDinus{}).
		Where("nim_dinus = ?", nimDinus).Update("kelas", classOption).
		Error; err != nil {
		return fmt.Errorf("failed to update class for nim_dinus=%s: %w", nimDinus, err)
	}

	return nil
}

func (s StudentStatusRepositoryImpl) GetAllKRSPick(ctx context.Context, db *gorm.DB, nimDinus string) ([]dto.SelectedKrs, error) {

	//FIND mahasiswa dinus where akdm stat 1
	//AMBIL JADWAL TAWAR WHERE KDMK = xx

	var results []dto.SelectedKrs

	err := db.WithContext(ctx).Model(&domain.JadwalTawar{}).
		Select(`
        DISTINCT jadwal_tawar.id AS id,
        krs_record.id AS krs_record_id,
        jadwal_tawar.ta AS tahun_ajaran,
        jadwal_tawar.kdmk AS kode_mata_kuliah,
        jadwal_tawar.klpk AS kelompok,
        matkul_kurikulum.nmmk AS nama_mata_kuliah,
        matkul_kurikulum.sks AS jumlah_sks,
        hari.nama AS hari,
        sesi_kuliah.jam_mulai,
        sesi_kuliah.jam_selesai,
        ruang.nama AS ruang,
        jadwal_tawar.jns_jam AS jns_jam
    `).
		Joins("LEFT JOIN matkul_kurikulum ON jadwal_tawar.kdmk = matkul_kurikulum.kdmk").
		Joins("LEFT JOIN hari ON jadwal_tawar.id_hari1 = hari.id").
		Joins("LEFT JOIN sesi_kuliah ON jadwal_tawar.id_sesi1 = sesi_kuliah.id").
		Joins("LEFT JOIN ruang ON jadwal_tawar.id_ruang1 = ruang.id").
		Joins("JOIN krs_record ON krs_record.kdmk = matkul_kurikulum.kdmk").
		Where("jadwal_tawar.ta IS NOT NULL AND krs_record.nim_dinus = ? AND krs_record.deleted_at IS NULL", nimDinus).
		Scan(&results).Error

	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, fmt.Errorf("failed to get KRS data: %w", err)
	}

	return results, nil
}

func (s StudentStatusRepositoryImpl) InsertKRSPermit(ctx context.Context, db *gorm.DB, nimDinus string) (bool, error) {

	var ijinKrs domain.MhsIjinKrs
	if err := db.WithContext(ctx).Where("nim_dinus = ?", nimDinus).First(&ijinKrs).Error; err != nil {
		return false, fmt.Errorf("error cant get permit status where nim")
	}

	if ijinKrs.Ijinkan == false {
		return false, nil
	}

	return true, nil
}

func (s StudentStatusRepositoryImpl) StatusKRS(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string) (dto.StatusKrsRes, error) {
	var status dto.StatusKrsRes

	//CHECK VALDASI DULU
	var validasi string

	exists := db.WithContext(ctx).Model(&domain.ValidasiKrsMhs{}).
		Select("1").
		Where("nim_dinus = ? AND ta = ?", nimDinus, kodeTA).
		Limit(1).
		Find(&validasi).RowsAffected

	if exists > 0 {
		validasi = "Validated" // Data ditemukan
	} else {
		validasi = "Not Validated" // Data tidak ditemukan
	}

	//CHECK DI PAKETKAN ATAU TIDAK
	var diPaketkanKah domain.MhsDipaketkan
	var paket string
	if err := db.WithContext(ctx).Where("nim_dinus = ?", nimDinus).First(&diPaketkanKah).Error; err != nil {
		paket = "tidak di paketkan"
	} else {
		paket = "dipaketkan"
	}

	//AMBIL DATA YG DI PERLUKAN
	var IpSemester struct {
		TahunAjaran string
		Sks         int
		Ips         string
		TahunMasuk  string
	}
	if err := db.WithContext(ctx).Raw(`
		SELECT
			ip_s.ta AS tahun_ajaran,
			ip_s.sks,
			ip_s.ips,
			md.ta_masuk AS tahun_masuk
		FROM ip_semester ip_s
		JOIN mahasiswa_dinus md ON ip_s.nim_dinus = md.nim_dinus
		WHERE ip_s.nim_dinus = ?
		LIMIT 1
	`, nimDinus).Scan(&IpSemester).Error; err != nil {
		return status, fmt.Errorf("error %v get data", nimDinus)
	}
	status = dto.StatusKrsRes{
		Validate:    validasi,
		Dipaketkan:  paket,
		TahunAjaran: IpSemester.TahunAjaran,
		TahunMasuk:  IpSemester.TahunMasuk,
		Sks:         IpSemester.Sks,
		Ips:         IpSemester.Ips,
	}

	return status, nil

}

func (s StudentStatusRepositoryImpl) KrsOffersProdi(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string, kelompok string) ([]dto.KrsOffersProdiResponse, error) {
	var krsOfferByProdi []dto.KrsOffersProdiResponse

	klp := fmt.Sprintf("%s%%", kelompok)
	query := `
		SELECT DISTINCT 
		    			jt.id   AS id,
		    			jt.ta   AS tahun_ajaran,
						jt.kdmk AS kode_mata_kuliah,
						jt.klpk AS kelompok,
						mk.nmmk AS nama_mata_kuliah,
						mk.sks  AS jumlah_sks,
						h.nama  AS hari,
						jt.jns_jam AS jns_jam,
						sk.jam_mulai,
						sk.jam_selesai,
						r.nama  AS ruang,
						CASE
							WHEN EXISTS (SELECT 1
										 FROM daftar_nilai dn
										 WHERE dn.kdmk = jt.kdmk
										   AND dn.nl = 'A'
										   AND dn.nim_dinus = ? ) THEN 'Tidak Bisa'
							ELSE 'Bisa'
							END AS status_pemilihan,
						CASE
							WHEN (
									 (SELECT COALESCE(SUM(sks), 0)
									  FROM krs_record
									  WHERE nim_dinus = ?)
										 + mk.sks
									 ) > (SELECT COALESCE(MAX(sks), 0)
										  FROM ip_semester
										  WHERE nim_dinus = ?
										  ORDER BY last_update
										  limit 1) THEN 'Tidak Mencukupi'
							ELSE CONCAT(
									'Jika di ambil Sisa ',
									(SELECT COALESCE(MAX(sks), 0)
									 FROM ip_semester
									 WHERE nim_dinus = ? 
									 ORDER BY last_update
									 limit 1)
										- (SELECT COALESCE(SUM(sks), 0)
										   FROM krs_record
										   WHERE nim_dinus = ?)
										- mk.sks
								 )
							END AS status_krs
		FROM jadwal_tawar jt
				 LEFT JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
				 LEFT JOIN hari h ON jt.id_hari1 = h.id
				 LEFT JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
				 LEFT JOIN ruang r ON jt.id_ruang1 = r.id
		WHERE jt.ta IS NOT NULL   -- Pastikan hanya menampilkan data valid
		  AND jt.klpk LIKE ? -- Hanya tampilkan kelompok yang dimulai dengan 'B11'
		  AND jt.ta = ?
		ORDER BY jt.ta, mk.nmmk;
		`

	if err := db.WithContext(ctx).Raw(query, nimDinus, nimDinus, nimDinus, nimDinus, nimDinus, klp, kodeTA).Scan(&krsOfferByProdi).Error; err != nil {
		return nil, fmt.Errorf("error show where nim %v tahunAjar %v not found and kel prodi %v", nimDinus, kodeTA)
	}

	fmt.Println(krsOfferByProdi)

	return krsOfferByProdi, nil
}

func (s StudentStatusRepositoryImpl) GetAllScores(ctx context.Context, db *gorm.DB, nimDinus string) ([]dto.AllScoresRes, error) {
	var scores []dto.AllScoresRes

	if err := db.WithContext(ctx).Table("daftar_nilai dn").
		Select(` DISTINCT
			mk.kdmk AS kode_matkul,
			mk.nmmk AS matakuliah,
			mk.sks AS sks,
			mk.tp AS category,
			mk.jenis_matkul AS jenis_matkul,
			dn.nl AS nilai
		`).
		Joins("JOIN matkul_kurikulum mk ON dn.kdmk = mk.kdmk").
		Where("dn.nim_dinus = ? AND dn.hide = 0", nimDinus).
		Scan(&scores).Error; err != nil {
		return nil, fmt.Errorf("error nilai where nim %v not found", nimDinus)
	}

	return scores, nil

}

func (s StudentStatusRepositoryImpl) ScheduleConflicts(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string) ([]dto.ScheduleConflictRes, error) {
	var schedules []dto.ScheduleConflictRes
	query := `
        SELECT DISTINCT
            jt.id   AS id,
            jt.ta AS tahun_ajaran,
            jt.klpk AS kelompok,
            mk.nmmk AS nama_mata_kuliah,
            mk.sks AS jumlah_sks,
            h.nama AS hari,
            sk.jam_mulai,
            sk.jam_selesai,
            jt.jns_jam AS jns_jam,
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
                 LEFT JOIN matkul_kurikulum mk ON jt.kdmk = mk.kdmk
                 LEFT JOIN hari h ON jt.id_hari1 = h.id
                 LEFT JOIN sesi_kuliah sk ON sk.id = jt.id_sesi1
                 JOIN ruang r ON jt.id_ruang1 = r.id
        WHERE mk.kur_aktif = 1
          AND jt.ta = ?
          AND jt.jns_jam IN (1, 2)
          AND jt.jsisa <= jt.jmax;
    `

	// Jalankan raw query
	if err := db.WithContext(ctx).Raw(query, nimDinus, kodeTA).Scan(&schedules).Error; err != nil {
		return schedules, fmt.Errorf("err not found ur nim %v and kodeTA %v not exist", nimDinus, kodeTA)
	}

	return schedules, nil
}

func (s StudentStatusRepositoryImpl) InsertKrsLog(ctx context.Context, db *gorm.DB, nimDinus string, rec domain.KrsRecord, Aksi int8) error {

	recordLog := domain.KrsRecordLog{
		IDKrs:    rec.ID,
		NimDinus: nimDinus,
		Kdmk:     rec.Kdmk,
		Aksi:     Aksi,
		IDJadwal: rec.IDJadwal,
		IpAddr:   "",
	}

	if err := db.WithContext(ctx).Create(&recordLog).Error; err != nil {
		return fmt.Errorf("err add to tabel record_log")
	}

	return nil
}

func (s StudentStatusRepositoryImpl) InsertKrs(ctx context.Context, db *gorm.DB, rec domain.KrsRecord) error {
	fmt.Println(rec)
	if err := db.WithContext(ctx).Create(&rec).Error; err != nil {
		fmt.Println(err)
		return fmt.Errorf("err add to tabel record")
	}

	return nil
}

func (s StudentStatusRepositoryImpl) GetKrsLog(ctx context.Context, db *gorm.DB, nimDinus string, kodeTA string) ([]dto.KrsLogRes, error) {

	var log []dto.KrsLogRes
	if err := db.WithContext(ctx).Raw(`
        SELECT krl.*
        FROM krs_record_log krl
        LEFT JOIN krs_record kr ON krl.id_krs = kr.id
        WHERE kr.ta = ? AND kr.nim_dinus = ?;
    `, kodeTA, nimDinus).Scan(&log).Error; err != nil {
		return nil, fmt.Errorf("err get record where nim %v and ta %v", nimDinus, kodeTA)
	}

	return log, nil

}

func (s StudentStatusRepositoryImpl) UpdateValidate(ctx context.Context, db *gorm.DB, mhs domain.ValidasiKrsMhs) error {
	if err := db.WithContext(ctx).Save(&mhs).Error; err != nil {
		return fmt.Errorf("err update validation status")
	}
	return nil
}

func (s StudentStatusRepositoryImpl) CheckTA(ctx context.Context, db *gorm.DB, kodeTA string) error {
	ta := domain.TahunAjaran{}
	if err := db.WithContext(ctx).Where("kode =?", kodeTA).First(&ta).Error; err != nil {
		return fmt.Errorf("kode %v not register", ta)
	}

	return nil
}
