package service

import (
	"context"
	"gorm.io/gorm"
	"koriebruh/try/dto"
	"koriebruh/try/repository"
	"strconv"
)

type StudentStatusService interface {
	InformationStudent(ctx context.Context, NimMhs string) (dto.StudentInfo, error)
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

func (s StudentStatusServicesImpl) InformationStudent(ctx context.Context, NimMhs string) (dto.StudentInfo, error) {
	var result dto.StudentInfo

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		student, err := s.StudentStatusRepository.InformationStudent(ctx, tx, NimMhs)
		if err != nil {
			return err
		}

		//CONVERT
		BayarStatus := true
		if student.TagihanMhs.SppBayar != 1 {
			BayarStatus = false
		}

		//CONVERT
		var StatusAkdemik string
		AkdmNum, _ := strconv.Atoi(student.MahasiswaDinus.AkdmStat)
		switch AkdmNum {
		case 1:
			StatusAkdemik = "Aktif"
		case 2:
			StatusAkdemik = "Cuti"
		case 3:
			StatusAkdemik = "Keluar"
		case 4:
			StatusAkdemik = "Lulus"
		case 5:
			StatusAkdemik = "Mangkir"
		case 6:
			StatusAkdemik = "Meninggal"
		case 7:
			StatusAkdemik = "DO"
		case 8:
			StatusAkdemik = "Aktif Keuangan"
		default:
			StatusAkdemik = "Status Tidak Dikenal"
		}

		//CONVERT
		var kelas string
		if student.MahasiswaDinus.Kelas == 1 {
			kelas = "pagi"
		} else if student.MahasiswaDinus.Kelas == 2 {
			kelas = "pagi"

		} else if student.MahasiswaDinus.Kelas == 3 {
			kelas = "pagi"

		} else {
			kelas = "not choose"
		}

		result = dto.StudentInfo{
			AkademikStat:     StatusAkdemik,
			Bayar:            BayarStatus,
			Bank:             student.TagihanMhs.SppBank,
			BayarDate:        student.TagihanMhs.SppBayarDate,
			TransactionId:    student.TagihanMhs.SppTransaksi,
			RegistrationDate: student.HerregistMahasiswa.DateReg,
			Kelas:            kelas,
		}

		return nil

	})

	if err != nil {
		return dto.StudentInfo{}, err
	}

	return result, nil
}
