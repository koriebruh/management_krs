package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"koriebruh/try/domain"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/repository"
	"log"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterReq) error
	Login(ctx context.Context, req dto.LoginReq) (string, error)
	CurrentAcc(ctx context.Context, nim string) (dto.CurrentUser, error)
}

type AuthServiceImpl struct {
	repository.UserRepository
	*gorm.DB
	*validator.Validate
	repository.CacheRepository
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate, cacheRepository repository.CacheRepository) *AuthServiceImpl {
	return &AuthServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate, CacheRepository: cacheRepository}
}

func (service AuthServiceImpl) Register(ctx context.Context, req dto.RegisterReq) error {
	if err := service.Validate.Struct(req); err != nil {
		return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
	}

	return service.DB.Transaction(func(tx *gorm.DB) error {
		var registerData domain.MahasiswaDinus

		password, err := bcrypt.GenerateFromPassword([]byte(req.PassMhs), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrInternalServer, err)
		}
		req.PassMhs = string(password)

		registerData = domain.MahasiswaDinus{
			NimDinus: req.NimDinus,
			TaMasuk:  req.TAMasuk,
			Prodi:    req.Prodi,
			PassMhs:  req.PassMhs,
			Kelas:    req.Kelas,
			AkdmStat: req.AkdmStat,
		}

		if err = service.UserRepository.Register(ctx, tx, registerData); err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}

		return nil
	})
}

func (service AuthServiceImpl) Login(ctx context.Context, req dto.LoginReq) (string, error) {
	var userNIM string

	if err := service.Validate.Struct(req); err != nil {
		return userNIM, fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		loginData := domain.MahasiswaDinus{
			NimDinus: req.NimDinus,
			PassMhs:  req.PassMhs,
		}

		result, err := service.UserRepository.Login(ctx, tx, loginData)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrBadRequest, err)
		}
		userNIM = *result

		return nil
	})

	if err != nil {
		return "", err
	}

	return userNIM, nil
}

func (service AuthServiceImpl) CurrentAcc(ctx context.Context, nim string) (dto.CurrentUser, error) {
	cacheKey := fmt.Sprintf("NimDinus := %s", nim)

	// Try get from cache
	cacheUser, err := service.CacheRepository.Get(ctx, cacheKey)
	if err == nil {
		log.Printf("DEBUG - Found in cache: %s", cacheUser)
		var cachedResult dto.CurrentUser
		if err := json.Unmarshal([]byte(cacheUser), &cachedResult); err == nil {
			log.Printf("DEBUG - Successfully unmarshaled: %+v", cachedResult)
			return cachedResult, nil
		}
		log.Printf("DEBUG - Unmarshal error: %v", err)
	} else {
		log.Printf("DEBUG - Cache get error: %v", err)
	}

	var result dto.CurrentUser

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		user, err := service.UserRepository.FindByNIM(ctx, tx, nim)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		result = dto.CurrentUser{
			NimDinus: user.NimDinus,
			TAMasuk:  user.TaMasuk,
			Prodi:    user.Prodi,
		}

		cacheErr := service.CacheRepository.Set(ctx, fmt.Sprint("user:"+nim), result, 15*time.Minute) // SAVE IN 15 HOUR
		if cacheErr != nil {
			// LOGGING ERROR CACHE
			log.Printf("Failed to cache user: %v", cacheErr)
		}

		return nil
	})

	if err != nil {
		return dto.CurrentUser{}, err
	}

	return result, nil

}
