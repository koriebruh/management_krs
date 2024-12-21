package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/try/domain"
	"koriebruh/try/dto"
)

type StudentStatusRepository interface {
	CheckUserExist(ctx context.Context, db *gorm.DB, nimDinus string) (domain.MahasiswaDinus, error)
	InformationStudent(ctx context.Context, db *gorm.DB, nimDinus string) (*dto.InfoStudentDB, error)
	SetClassTime(ctx context.Context, db *gorm.DB, nimDinus string, classOption int) error
	GetAllKRSPick(ctx context.Context, db *gorm.DB, nimDinus string) ([]dto.SelectedKrs, error)

	//ExceptionInsertKRS(ctx context.Context, db *gorm.DB)
	//StatusKRS(ctx context.Context, db *gorm.DB)
}

type StudentStatusRepositoryImpl struct {
}

func NewStudentStatusRepository() *StudentStatusRepositoryImpl {
	return &StudentStatusRepositoryImpl{}
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
		Joins("JOIN krs_management.tagihan_mhs ON herregist_mahasiswa.nim_dinus = tagihan_mhs.nim_dinus").
		Where("mahasiswa_dinus.nim_dinus = ?", nimDinus).
		Scan(&studentStatus).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get status user %v", nimDinus)
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

	err := db.WithContext(ctx).Model(&domain.KrsRecord{}).
		Select("matkul_kurikulum.nmmk AS nama_matkul, matkul_kurikulum.nmen AS nama_matkul_en, matkul_kurikulum.tp AS tipe, matkul_kurikulum.smt AS semester, matkul_kurikulum.jenis_matkul AS jenis_matkul, hari1.nama AS hari1, hari2.nama AS hari2, hari3.nama AS hari3").
		Joins("JOIN matkul_kurikulum ON matkul_kurikulum.kdmk = krs_record.kdmk").
		Joins("JOIN jadwal_tawar ON jadwal_tawar.id = krs_record.id_jadwal").
		Joins("LEFT JOIN hari AS hari1 ON hari1.id = jadwal_tawar.id_hari1").
		Joins("LEFT JOIN hari AS hari2 ON hari2.id = jadwal_tawar.id_hari2").
		Joins("LEFT JOIN hari AS hari3 ON hari3.id = jadwal_tawar.id_hari3").
		Where("krs_record.nim_dinus = ?", nimDinus).
		Scan(&results).Error

	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, fmt.Errorf("failed to get KRS data: %w", err)
	}

	return results, nil
}
