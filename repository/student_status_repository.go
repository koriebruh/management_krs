package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/try/req_db"
)

type StudentStatusRepository interface {
	InformationStudent(ctx context.Context, db *gorm.DB, nimDinus string) (*req_db.InformationStudent, error)
	//SetClassTime(ctx context.Context, db *gorm.DB)
	//GetAllKRSPick(ctx context.Context, db *gorm.DB)
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

	if err := db.WithContext(ctx).Preload("TagihanMhs").
		Preload("HerregistMahasiswa").
		Preload("MahasiswaDinus").
		Where("mahasiswa_dinus.nim_dinus = ?", nimDinus).
		First(&studentStatus).Error; err != nil {

		return nil, fmt.Errorf("error nih bg %e", err)
	}

	return &studentStatus, nil
}
