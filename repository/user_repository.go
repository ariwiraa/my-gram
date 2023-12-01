package repository

import (
	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user domain.User) (domain.User, error)
	FindByUsername(username string) (domain.User, error)
	FindUsersByIDList(id []uint) ([]domain.User, error)
	IsUsernameExists(username string) (bool, error)
	IsEmailExists(email string) (bool, error)
	IsUserExists(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) IsUserExists(id uint) error {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUsersByIDList(userIds []uint) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Debug().Preload("Photos").Find(&users, "users.id IN ?", userIds).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

// IsEmailExists implements UserRepository
func (r *userRepository) IsEmailExists(email string) (bool, error) {
	var user domain.User
	err := r.db.Debug().Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// IsUsernameExists implements UserRepository
func (r *userRepository) IsUsernameExists(username string) (bool, error) {
	var user domain.User
	err := r.db.Debug().Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// IsEmailExist implements UserRepository
func (r *userRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	// err := r.db.Debug().Where("username = ?", username).Find(&user).Error
	err := r.db.Preload("Photos").First(&user, "username = ?", username).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// AddUser implements domain.UserRepository
func (r *userRepository) AddUser(user domain.User) (domain.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
