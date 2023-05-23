package repository

import (
	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	IsUsernameExists(username string) (bool, error)
	IsEmailExists(email string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
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
func (r *userRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	// err := r.db.Debug().Where("email = ?", email).Find(&user).Error
	err := r.db.Debug().First(&user, "email = ?", email).Error
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
