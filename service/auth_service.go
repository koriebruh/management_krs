package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"koriebruh/try/domain"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/repository"
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
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) *AuthServiceImpl {
	return &AuthServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (service AuthServiceImpl) Register(ctx context.Context, req dto.RegisterReq) error {
	if err := service.Validate.Struct(req); err != nil {
		return fmt.Errorf("%w: %v", helper.ErrValidationFailed, err)
	}

	return service.DB.Transaction(func(tx *gorm.DB) error {
		var registerData domain.User

		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrPasswordEncryption, err)
		}
		req.Password = string(password)

		registerData = domain.User{
			NIM:      req.NIM,
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
		}

		if err = service.UserRepository.Register(ctx, tx, registerData); err != nil {
			return fmt.Errorf("%w: %v", helper.ErrUserRegistration, err)
		}

		return nil
	})
}

func (service AuthServiceImpl) Login(ctx context.Context, req dto.LoginReq) (string, error) {
	var userNIM string

	if err := service.Validate.Struct(req); err != nil {
		return userNIM, fmt.Errorf("%w: %v", helper.ErrValidationFailed, err)
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		loginData := domain.User{
			NIM:      req.NIM,
			Password: req.Password,
		}

		result, err := service.UserRepository.Login(ctx, tx, loginData)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrLoginFailed, err)
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
	var result dto.CurrentUser

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		user, err := service.UserRepository.FindByNIM(ctx, tx, nim)
		if err != nil {
			return fmt.Errorf("%w: %v", helper.ErrNotFound, err)
		}

		result = dto.CurrentUser{
			NIM:      user.NIM,
			Username: user.Username,
			Email:    user.Email,
		}

		return nil
	})

	if err != nil {
		return dto.CurrentUser{}, err
	}

	return result, nil

}
