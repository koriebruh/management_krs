package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"koriebruh/try/domain"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/repository"
	"koriebruh/try/req_db"
)

type StudentStatusService interface {
	InformationStudent(ctx context.Context, NimMhs string) (req_db.InformationStudent, error)
	SetClassTime(ctx context.Context, nimDinus string, req dto.ChangeClassReq) error
	GetAllKRSPick(ctx context.Context, nimDinus string) ([]domain.KrsRecord, error)
	//ExceptionInsertKRS()
	//StatusKRS()
}
type StudentStatusServicesImpl struct {
	*gorm.DB
	repository.StudentStatusRepository
	*validator.Validate
}

func NewStudentStatusServices(DB *gorm.DB, studentStatusRepository repository.StudentStatusRepository, validate *validator.Validate) *StudentStatusServicesImpl {
	return &StudentStatusServicesImpl{DB: DB, StudentStatusRepository: studentStatusRepository, Validate: validate}
}

func (s StudentStatusServicesImpl) InformationStudent(ctx context.Context, NimMhs string) (req_db.InformationStudent, error) {
	var result req_db.InformationStudent

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		student, err := s.StudentStatusRepository.InformationStudent(ctx, tx, NimMhs)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}
		result = *student

		return nil
	})

	if err != nil {
		return req_db.InformationStudent{}, err
	}

	return result, nil
}

func (s StudentStatusServicesImpl) SetClassTime(ctx context.Context, nimDinus string, req dto.ChangeClassReq) error {
	if err := s.Validate.Struct(req); err != nil {
		return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
	}

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := s.StudentStatusRepository.SetClassTime(ctx, tx, nimDinus, req.Kelas)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s StudentStatusServicesImpl) GetAllKRSPick(ctx context.Context, nimDinus string) ([]domain.KrsRecord, error) {
	var result []domain.KrsRecord

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		pick, err := s.StudentStatusRepository.GetAllKRSPick(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		result = pick

		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}
