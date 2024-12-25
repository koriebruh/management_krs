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
	"strconv"
	"strings"
	"time"
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
	InsertSchedule(ctx context.Context, nimDinus string, kodeTA string, idSchedule int) (string, error)
	GetKrsLog(ctx context.Context, nimDinus string, kodeTA string) ([]dto.KrsLogRes, error)
	DeleteKrsRecByIdKrs(ctx context.Context, nimDinus string, idKrs int) (string, error)
	UpdateValidate(ctx context.Context, kodeTA string, req dto.UpdateValidateReq) (string, error)
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

func (s StudentStatusServicesImpl) InsertSchedule(ctx context.Context, nimDinus string, kodeTA string, idSchedule int) (string, error) {
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		userExist, err := s.StudentStatusRepository.CheckUserExist(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}
		Prodi := userExist.Prodi

		//CEK APAKAH PUNYA IJIN INSERT KRS
		_, err = s.StudentStatusRepository.InsertKRSPermit(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("no have permit insert krs"))
		}

		//KRS YG SUDAH DI VALIDASI TIDAK BISA INSERT
		statusKRS, _ := s.StudentStatusRepository.StatusKRS(ctx, tx, nimDinus)
		if statusKRS.Validate == "Validated" {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("krs already validated can't insert"))
		}

		//AMBIL DATA JADWAL
		conflictsSchedule, err := s.StudentStatusRepository.ScheduleConflicts(ctx, tx, nimDinus, kodeTA)
		var foundSchedule dto.ScheduleConflictRes
		for _, s := range conflictsSchedule {
			if s.Id == idSchedule {
				foundSchedule = dto.ScheduleConflictRes{
					Id:             s.Id,
					TahunAjaran:    s.TahunAjaran,
					Kelompok:       s.Kelompok,
					NamaMataKuliah: s.NamaMataKuliah,
					JumlahSKS:      s.JumlahSKS,
					Hari:           s.Hari,
					JamMulai:       s.JamMulai,
					JamSelesai:     s.JamSelesai,
					Ruang:          s.Ruang,
					StatusBentrok:  s.StatusBentrok,
					KeteranganSlot: s.KeteranganSlot,
				}
				break
			}
		}
		fmt.Println(foundSchedule)

		//CEK JADWAL BENTROK KAH ?
		//CEK SLOT PENUH KAH ?
		if foundSchedule.StatusBentrok == "BENTROK" {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("krs bentrok bang"))
		} else if strings.Contains(foundSchedule.KeteranganSlot, "SLOT PENUH") {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("slot sudah penuh"))
		}

		//UNTUK PENGECEKAN APAKAH YG DI PILIH ADALAH YG SEBELUM NYA APA A  DAN MENGECEK APAKAH KRS MENCUKUPI
		offersProdi, err := s.StudentStatusRepository.KrsOffersProdi(ctx, tx, nimDinus, kodeTA, Prodi)
		var foundOfferProdi dto.KrsOffersProdiResponse
		for _, p := range offersProdi {
			if p.Id == idSchedule {
				foundOfferProdi = dto.KrsOffersProdiResponse{
					Id:              p.Id,
					TahunAjaran:     p.TahunAjaran,
					KodeMataKuliah:  p.KodeMataKuliah,
					Kelompok:        p.Kelompok,
					NamaMataKuliah:  p.NamaMataKuliah,
					JumlahSKS:       p.JumlahSKS,
					Hari:            p.Hari,
					JamMulai:        p.JamMulai,
					JamSelesai:      p.JamSelesai,
					Ruang:           p.Ruang,
					StatusPemilihan: p.StatusPemilihan,
					StatusKrs:       p.StatusKrs,
				}
			}
		}
		fmt.Println(foundOfferProdi)

		if foundOfferProdi.StatusPemilihan == "Tidak Bisa" {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("anda sebelumnya sudah dapat A tidak bisa memilih ini"))
		} else if foundOfferProdi.StatusKrs == "Tidak Mencukupi" {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("krs anda tidak mencukupi"))
		}

		// MEMASTIKAN JAM (PAGI,MALAM) SESUAI YG DI SET
		var jnsJam string
		if err = tx.WithContext(ctx).
			Model(domain.JadwalTawar{}).
			Select("jns_jam").
			Where("id = ?", idSchedule).
			Scan(&jnsJam).Error; err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("jenis kelas dari shecdule yg di pilih tidak ketemu"))
		}

		fmt.Println("JENIS JAM ", jnsJam)
		fmt.Println("JENIS JAM ", userExist.Kelas)

		if jnsJam != strconv.Itoa(userExist.Kelas) {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("kelas anda %v sedangkan kelas yang ingin ada pilih %v ", userExist.Kelas, jnsJam))
		}

		//JIKA NANTI ERROR BEARRTI ANTARA foundOfferProdi DAN foundSchedule MALSAHNAYA DATANYA ANEH KADANG TIDAKA DI SALAH 1
		TA, _ := strconv.Atoi(kodeTA)
		record := domain.KrsRecord{
			TA:       TA,
			Kdmk:     foundOfferProdi.KodeMataKuliah,
			IDJadwal: foundOfferProdi.Id,
			NimDinus: nimDinus,
			Sts:      "B", // B MAKSUNYA APA G PAHAM GA DI JELASIN JUGA DI
			Sks:      foundOfferProdi.JumlahSKS,
			Modul:    0,
		}

		if err = s.StudentStatusRepository.InsertKrs(ctx, tx, record); err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		if err = s.StudentStatusRepository.InsertKrsLog(ctx, tx, nimDinus, record, 1); err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)

		}

		// TAMBAH JSISA
		if err = tx.WithContext(ctx).
			Model(&domain.JadwalTawar{}).
			Where("id = ?", idSchedule).
			UpdateColumn("jsisa", gorm.Expr("jsisa + ?", 1)).Error; err != nil {
			return fmt.Errorf("%w: %v", helper.ErrInternalServer, fmt.Errorf("gagal update j sisa"))

		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return "SUCCESS ADD NEW SCHEDULE", nil

}

func (s StudentStatusServicesImpl) GetKrsLog(ctx context.Context, nimDinus string, kodeTA string) ([]dto.KrsLogRes, error) {

	var results []dto.KrsLogRes

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		_, err := s.StudentStatusRepository.CheckUserExist(ctx, tx, nimDinus)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		log, err := s.StudentStatusRepository.GetKrsLog(ctx, tx, nimDinus, kodeTA)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		results = log
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s StudentStatusServicesImpl) DeleteKrsRecByIdKrs(ctx context.Context, nimDinus string, idKrs int) (string, error) {
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		var krsRec domain.KrsRecord
		if err := tx.WithContext(ctx).Where("id=?", idKrs).First(&krsRec).Error; err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, fmt.Errorf("wrong id_krs %d not found", idKrs))
		}
		//ADD TO KRS LOG
		if err := s.StudentStatusRepository.InsertKrsLog(ctx, tx, nimDinus, krsRec, 0); err != nil {
			return fmt.Errorf("%w: %v", helper.ErrInternalServer, err)
		}
		idSchedule := krsRec.ID

		//VALIDASI YG SUDAH DI VALUDASI TIDAK BISA DELETE
		statusKRS, err := s.StudentStatusRepository.StatusKRS(ctx, tx, nimDinus)
		if statusKRS.Validate == "Validated" {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		// MENGAHPUS KRS RECORD, INI SOFT DELETE
		if err := tx.WithContext(ctx).Delete(&krsRec).Error; err != nil {
			return fmt.Errorf("%w: %v", helper.ErrInternalServer, fmt.Errorf("err delete"))
		}

		// MENGURAGI Jsisakan di batal kan krs nya
		if err := tx.WithContext(ctx).
			Model(&domain.JadwalTawar{}).
			Where("id = ?", idSchedule).
			UpdateColumn("jsisa", gorm.Expr("jsisa - ?", 1)).Error; err != nil {
			return fmt.Errorf("%w: %v", helper.ErrInternalServer, fmt.Errorf("err update jsisa"))
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("success delete shecudle where id krs = %d", idKrs), nil
}

func (s StudentStatusServicesImpl) UpdateValidate(ctx context.Context, nimDinus string, req dto.UpdateValidateReq) (string, error) {
	if err := s.Validate.Struct(req); err != nil {
		return "", fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
	}

	err := s.DB.Transaction(func(tx *gorm.DB) error {

		if err := s.StudentStatusRepository.CheckTA(ctx, tx, strconv.Itoa(req.TA)); err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		data := domain.ValidasiKrsMhs{
			NimDinus: nimDinus,
			JobDate:  time.Now(),
			JobHost:  req.JobHost,
			JobAgent: req.JobAgent,
			TA:       req.TA,
		}

		if err := s.StudentStatusRepository.UpdateValidate(ctx, tx, data); err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return "success update validate status", nil
}
