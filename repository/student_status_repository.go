package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/try/domain"
	"koriebruh/try/req_db"
)

type StudentStatusRepository interface {
	InformationStudent(ctx context.Context, db *gorm.DB, nimDinus string) (*req_db.InformationStudent, error)
	SetClassTime(ctx context.Context, db *gorm.DB, nimDinus string, classOption int) error
	GetAllKRSPick(ctx context.Context, db *gorm.DB, nimDinus string) ([]domain.KrsRecord, error)

	//ExceptionInsertKRS(ctx context.Context, db *gorm.DB)
	//StatusKRS(ctx context.Context, db *gorm.DB)
}

type StudentStatusRepositoryImpl struct {
}

func NewStudentStatusRepository() *StudentStatusRepositoryImpl {
	return &StudentStatusRepositoryImpl{}
}

func (s StudentStatusRepositoryImpl) InformationStudent(ctx context.Context, db *gorm.DB, nimDinus string) (*req_db.InformationStudent, error) {
	var studentStatus req_db.InformationStudent

	var Mhs domain.MahasiswaDinus
	if err := db.WithContext(ctx).Where("nim_dinus = ?", nimDinus).First(&Mhs).Error; err != nil {
		return nil, fmt.Errorf("error find where nim := %v and err is %e", nimDinus, err)
	}

	var Heregis domain.HerregistMahasiswa
	if err := db.WithContext(ctx).Where("nim_dinus = ?", nimDinus).First(&Heregis).Error; err != nil {
		return nil, fmt.Errorf("error find herregis where nim := %v and err is %e", nimDinus, err)
	}

	var Tagihan []domain.TagihanMhs
	if err := db.WithContext(ctx).Where("nim_dinus = ?", nimDinus).Find(&Tagihan).Error; err != nil {
		return nil, fmt.Errorf("error find tagihan  where nim := %v and err is %e", nimDinus, err)
	}

	//MAPING
	studentStatus = req_db.InformationStudent{
		MahasiswaDinus:     Mhs,
		HerregistMahasiswa: Heregis,
		TagihanMahasiswa:   Tagihan,
	}

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

func (s StudentStatusRepositoryImpl) GetAllKRSPick(ctx context.Context, db *gorm.DB, nimDinus string) ([]domain.KrsRecord, error) {

	//FIND mahasiswa dinus where akdm stat 1
	//AMBIL JADWAL TAWAR WHERE KDMK = xx

	var krsList []domain.KrsRecord

	err := db.Preload("TahunAjaran").
		Preload("MataKuliah").
		Preload("Jadwal").
		Preload("Mahasiswa").
		Where("nim_dinus = ?", nimDinus).
		Find(&krsList).Error

	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, fmt.Errorf("failed to get KRS data: %w", err)
	}

	return krsList, nil
}
