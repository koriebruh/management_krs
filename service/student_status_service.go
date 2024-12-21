package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/repository"
	"strconv"
)

type StudentStatusService interface {
	InformationStudent(ctx context.Context, NimMhs string) (dto.InfoStudentRes, error)
	SetClassTime(ctx context.Context, nimDinus string, req dto.ChangeClassReq) error
	GetAllKRSPick(ctx context.Context, nimDinus string) ([]dto.SelectedKrs, error)
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

func (s StudentStatusServicesImpl) InformationStudent(ctx context.Context, NimMhs string) (dto.InfoStudentRes, error) {

	var repositoryData dto.InfoStudentDB
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//FIND EXIST OR NO
		user, err := s.StudentStatusRepository.CheckUserExist(ctx, tx, NimMhs)
		if err != nil {
			return err
		}

		student, err := s.StudentStatusRepository.InformationStudent(ctx, tx, NimMhs)
		if err != nil {
			return err
		}
		repositoryData = *student

		//JIKA DATA INFO KOSONG MAKA
		isEmpty := func(data dto.InfoStudentDB) bool {
			return data == (dto.InfoStudentDB{})
		}

		if isEmpty(repositoryData) {
			atoi, _ := strconv.Atoi(user.AkdmStat)

			repositoryData.Kelas = strconv.Itoa(user.Kelas)
			repositoryData.Prodi = user.Prodi
			repositoryData.NimDinus = user.NimDinus
			repositoryData.AkdmStat = atoi
			repositoryData.TaMasuk = user.TaMasuk
		}

		return nil
	})

	fmt.Println(&repositoryData)

	var result dto.InfoStudentRes
	if err != nil {
		return result, err
	}

	if repositoryData.SppStatus == 1 {
		result.SppStatus = "paid"
	} else {
		result.SppStatus = "unpaid"
	}

	switch repositoryData.AkdmStat {
	case 1:
		result.AkdmStat = "aktif"
	case 2:
		result.AkdmStat = "Cuti"
	case 3:
		result.AkdmStat = "Keluar"
	case 4:
		result.AkdmStat = "Lulus"
	case 5:
		result.AkdmStat = "Mangkir"
	case 6:
		result.AkdmStat = "Meninggal"
	case 7:
		result.AkdmStat = "DO"
	case 8:
		result.AkdmStat = "Aktif Keuangan"
	default:
		result.AkdmStat = "Status Tidak Dikenal"
	}

	if repositoryData.Kelas == "1" {
		result.Kelas = "pagi"
	} else if repositoryData.Kelas == "2" {
		result.Kelas = "pagi"

	} else if repositoryData.Kelas == "3" {
		result.Kelas = "pagi"
	} else {
		result.Kelas = "belum dipilih"
	}

	result.NimDinus = repositoryData.NimDinus
	result.TaMasuk = strconv.Itoa(repositoryData.TaMasuk)
	result.Prodi = repositoryData.Prodi
	result.DateReg = repositoryData.DateReg

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

func (s StudentStatusServicesImpl) GetAllKRSPick(ctx context.Context, nimDinus string) ([]dto.SelectedKrs, error) {
	var results []dto.SelectedKrs

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		pick, err := s.StudentStatusRepository.GetAllKRSPick(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		results = pick

		return nil
	})

	if err != nil {
		return results, err
	}

	return results, nil
}
