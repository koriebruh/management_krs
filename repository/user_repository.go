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
	Register(ctx context.Context, db *gorm.DB, user domain.User) error
	Login(ctx context.Context, db *gorm.DB, user domain.User) (*string, error)
	FindById(ctx context.Context, db *gorm.DB, id int) (*domain.User, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r UserRepositoryImpl) Register(ctx context.Context, db *gorm.DB, user domain.User) error {
	var existingUser *domain.User

	err := db.WithContext(ctx).
		Where("email = ? OR username = ?", user.Email, user.Username).
		Select("email, username"). // Select specific fields only
		First(&existingUser).
		Error

	if err == nil {
		if existingUser.Email == user.Email {
			return fmt.Errorf("email already registered")
		}
		return fmt.Errorf("username already taken")
	}

	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking user existence: %v", err)
	}

	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("error create account %v", err)
	}

	return nil
}

func (r UserRepositoryImpl) Login(ctx context.Context, db *gorm.DB, user domain.User) (*string, error) {

	// temp data from db
	var result *domain.User

	err := db.WithContext(ctx).
		Where("nim = ?", user.NIM).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("login failed: incorrect user and password")
		}
		return nil, fmt.Errorf("login failed: %v", err)
	}

	// validation pass
	if errPass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); errPass != nil {
		return nil, fmt.Errorf("login failed decrypt : %v", errPass)
	}

	return result.NIM, nil

}

func (r UserRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, id int) (*domain.User, error) {
	var result *domain.User

	if err := db.WithContext(ctx).First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to fetch product: %v", err)
	}

	return result, nil
}
