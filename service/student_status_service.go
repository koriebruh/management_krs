package service

import (
	"context"
	"gorm.io/gorm"
	"koriebruh/try/repository"
	"koriebruh/try/req_db"
)

type StudentStatusService interface {
	InformationStudent(ctx context.Context, NimMhs string) (req_db.InformationStudent, error)
	//SetClassTime()
	//GetAllKRSPick()
	//ExceptionInsertKRS()
	//StatusKRS()
}
type StudentStatusServicesImpl struct {
	*gorm.DB
	repository.StudentStatusRepository
}

func NewStudentStatusServices(DB *gorm.DB, studentStatusRepository repository.StudentStatusRepository) *StudentStatusServicesImpl {
	return &StudentStatusServicesImpl{DB: DB, StudentStatusRepository: studentStatusRepository}
}

func (s StudentStatusServicesImpl) InformationStudent(ctx context.Context, NimMhs string) (req_db.InformationStudent, error) {
	var result req_db.InformationStudent

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		student, err := s.StudentStatusRepository.InformationStudent(ctx, tx, NimMhs)
		if err != nil {
			return err
		}
		result = *student

		return nil
	})

	if err != nil {
		return req_db.InformationStudent{}, err
	}

	return result, nil
}
