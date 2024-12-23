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
	KrsOffers(ctx context.Context, kodeTA string) ([]dto.KrsOfferRes, error)
	KrsSchedule(ctx context.Context, nimMhs string) (dto.KrsScheduleRes, error)
	InformationStudent(ctx context.Context, NimMhs string) (dto.InfoStudentRes, error)
	SetClassTime(ctx context.Context, nimDinus string, req dto.ChangeClassReq) error
	GetAllKRSPick(ctx context.Context, nimDinus string) ([]dto.SelectedKrs, error)
	InsertKRSPermit(ctx context.Context, nimDinus string) (string, error)
	StatusKRS(ctx context.Context, nimDinus string) (dto.StatusKrsRes, error)
	KrsOffersProdi(ctx context.Context, nimDinus string, kodeTA string) ([]dto.KrsOffersProdiResponse, error)
	GetAllScores(ctx context.Context, nimDinus string) ([]dto.AllScoresRes, error)
	ScheduleConflicts(ctx context.Context, nimDinus string, kodeTA string) ([]dto.ScheduleConflictRes, error)
}
type StudentStatusServicesImpl struct {
	*gorm.DB
	repository.StudentStatusRepository
	*validator.Validate
}

func NewStudentStatusServices(DB *gorm.DB, studentStatusRepository repository.StudentStatusRepository, validate *validator.Validate) *StudentStatusServicesImpl {
	return &StudentStatusServicesImpl{DB: DB, StudentStatusRepository: studentStatusRepository, Validate: validate}
}

func (s StudentStatusServicesImpl) KrsOffers(ctx context.Context, kodeTA string) ([]dto.KrsOfferRes, error) {
	KrsListOffers, err := s.StudentStatusRepository.KrsOffers(ctx, s.DB, kodeTA)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", helper.ErrNotFound, err)
	}
	return KrsListOffers, nil
}

func (s StudentStatusServicesImpl) KrsSchedule(ctx context.Context, nimMhs string) (dto.KrsScheduleRes, error) {
	var result dto.KrsScheduleRes
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		exist, err := s.StudentStatusRepository.CheckUserExist(ctx, tx, nimMhs)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		schedule, err := s.StudentStatusRepository.KrsSchedule(ctx, tx, exist.Prodi)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		result = dto.KrsScheduleRes{
			TA:         schedule.TA,
			Prodi:      schedule.Prodi,
			TglMulai:   schedule.TglMulai.Format("2006-01-02 15:04:05"),
			TglSelesai: schedule.TglSelesai.Format("2006-01-02 15:04:05"),
		}

		return nil
	})

	if err != nil {
		return dto.KrsScheduleRes{}, err
	}

	return result, nil

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

func (s StudentStatusServicesImpl) InsertKRSPermit(ctx context.Context, nimDinus string) (string, error) {
	permit, err := s.StudentStatusRepository.InsertKRSPermit(ctx, s.DB, nimDinus)
	if err != nil {
		return "", fmt.Errorf("%w: %v", helper.ErrNotFound, err)
	}

	if permit == false {
		return "not allowed insert krs", nil
	}

	return "allowed insert krs", nil

}

func (s StudentStatusServicesImpl) StatusKRS(ctx context.Context, nimDinus string) (dto.StatusKrsRes, error) {
	var result dto.StatusKrsRes

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		krsStatus, err := s.StudentStatusRepository.StatusKRS(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}
		result = krsStatus
		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil

}

func (s StudentStatusServicesImpl) KrsOffersProdi(ctx context.Context, nimDinus string, kodeTA string) ([]dto.KrsOffersProdiResponse, error) {
	var results []dto.KrsOffersProdiResponse

	err := s.DB.Transaction(func(tx *gorm.DB) error {

		userExist, err := s.StudentStatusRepository.CheckUserExist(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		prodiSchedule, err := s.StudentStatusRepository.KrsOffersProdi(ctx, tx, nimDinus, kodeTA, userExist.Prodi)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		results = prodiSchedule
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil

}

func (s StudentStatusServicesImpl) GetAllScores(ctx context.Context, nimDinus string) ([]dto.AllScoresRes, error) {
	var results []dto.AllScoresRes

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		scores, err := s.StudentStatusRepository.GetAllScores(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}
		results = scores
		return nil
	})

	if err != nil {
		return results, err
	}

	return results, nil
}

func (s StudentStatusServicesImpl) ScheduleConflicts(ctx context.Context, nimDinus string, kodeTA string) ([]dto.ScheduleConflictRes, error) {
	var results []dto.ScheduleConflictRes

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		conflicts, err := s.StudentStatusRepository.ScheduleConflicts(ctx, tx, nimDinus, kodeTA)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		results = conflicts
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil

}
