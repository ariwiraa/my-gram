package repository

import (
	"context"
	"errors"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(ctx context.Context, user domain.User) (domain.User, error)
	FindByUsername(ctx context.Context, username string) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindById(ctx context.Context, id uint) (*domain.User, error)
	FindUsersByIDList(ctx context.Context, id []uint) ([]domain.User, error)
	IsUsernameExists(ctx context.Context, username string) (bool, error)
	IsEmailExists(ctx context.Context, email string) (bool, error)
	IsUserExists(ctx context.Context, id uint) error
	UpdateUser(ctx context.Context, user domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// UpdateUser implements UserRepository.
func (r *userRepository) UpdateUser(ctx context.Context, user domain.User) error {
	err := r.db.WithContext(ctx).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ErrUserNotFound
		}
		log.Printf("[UpdateUser] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil
}

// FindByEmail implements UserRepository.
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, helpers.ErrEmailNotFound
		}
		log.Printf("[FindByEmail] with error detail %v", err.Error())
		return &user, helpers.ErrRepository
	}

	return &user, nil
}

func (r *userRepository) FindById(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, helpers.ErrUserNotFound
		}
		log.Printf("[FindById] with error detail %v", err.Error())
		return &user, helpers.ErrRepository
	}

	return &user, nil
}

func (r *userRepository) IsUserExists(ctx context.Context, id uint) error {
	var user domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ErrUserNotFound
		}
		log.Printf("[IsUserExists] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil
}

func (r *userRepository) FindUsersByIDList(ctx context.Context, userIds []uint) ([]domain.User, error) {
	var users []domain.User
	err := r.db.WithContext(ctx).Preload("Photos").Find(&users, "users.id IN ?", userIds).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, helpers.ErrUserNotFound
		}
		log.Printf("[FindUsersByIDList] with error detail %v", err.Error())
		return users, helpers.ErrRepository
	}

	return users, nil
}

// IsEmailExists implements UserRepository
func (r *userRepository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Printf("[IsEmailExists] with error detail %v", err.Error())
		return false, helpers.ErrRepository
	}

	return true, nil
}

// IsUsernameExists implements UserRepository
func (r *userRepository) IsUsernameExists(ctx context.Context, username string) (bool, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Printf("[IsUsernameExists] with error detail %v", err.Error())
		return false, helpers.ErrRepository
	}

	return true, nil
}

// IsEmailExist implements UserRepository
func (r *userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User

	err := r.db.WithContext(ctx).Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "photo_url", "caption", "created_at", "user_id")
	}).
		First(&user, "username = ?", username).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, helpers.ErrUserNotFound
		}
		log.Printf("[FindByUsername] with error detail %v", err.Error())
		return user, helpers.ErrRepository
	}

	return user, nil
}

// AddUser implements domain.UserRepository
func (r *userRepository) AddUser(ctx context.Context, user domain.User) (domain.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, helpers.ErrRepository
	}

	return user, nil
}
