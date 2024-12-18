package repository

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"koriebruh/try/domain"
)

type UserRepository interface {
	Register(ctx context.Context, db *gorm.DB, user domain.MahasiswaDinus) error
	Login(ctx context.Context, db *gorm.DB, user domain.MahasiswaDinus) (*string, error)
	FindByNIM(ctx context.Context, db *gorm.DB, NIM string) (*domain.MahasiswaDinus, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r UserRepositoryImpl) Register(ctx context.Context, db *gorm.DB, user domain.MahasiswaDinus) error {
	var existingUser *domain.MahasiswaDinus

	err := db.WithContext(ctx).
		Where("nim_dinus = ?", user.NimDinus).
		Select("nim_dinus"). // Select specific fields only
		First(&existingUser).
		Error

	if err == nil {
		if existingUser.NimDinus == user.NimDinus {
			return fmt.Errorf("NimDinus Already Registered")
		}
		return fmt.Errorf("NimDinus Taken")
	}

	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking user existence: %v", err)
	}

	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("error Create New Mahasiswa %v", err)
	}

	return nil
}

func (r UserRepositoryImpl) Login(ctx context.Context, db *gorm.DB, user domain.MahasiswaDinus) (*string, error) {

	// temp data from db
	var result *domain.MahasiswaDinus

	err := db.WithContext(ctx).
		Where("nim_dinus = ?", user.NimDinus).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("nim_dinus un register")
		}
		return nil, fmt.Errorf("incorrect nim_dinus and password")
	}

	// validation pass
	if errPass := bcrypt.CompareHashAndPassword([]byte(result.PassMhs), []byte(user.PassMhs)); errPass != nil {
		return nil, fmt.Errorf("incorrect nim_dinus and password ")
	}

	return &result.NimDinus, nil

}

func (r UserRepositoryImpl) FindByNIM(ctx context.Context, db *gorm.DB, NIM string) (*domain.MahasiswaDinus, error) {
	var result *domain.MahasiswaDinus

	if err := db.WithContext(ctx).Where("nim_dinus = ?", NIM).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("nim_dinus un register")
		}
		return nil, fmt.Errorf(" Fail to get detail Current User %v", err)
	}

	return result, nil
}
