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

	if err := db.WithContext(ctx).Model(&domain.MahasiswaDinus{}).
		Where("nim_dinus = ?", nimDinus).Update("kelas", classOption).
		Error; err != nil {
		return fmt.Errorf("failed to update class for nim_dinus=%s: %w", nimDinus, err)
	}

	return nil
}
