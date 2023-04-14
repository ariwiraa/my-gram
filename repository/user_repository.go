package repository

import (
	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	IsUsernameExists(username string) error
}

type userRepository struct {
	db *gorm.DB
}

// IsUsernameExists implements UserRepository
func (r *userRepository) IsUsernameExists(username string) error {
	var user domain.User
	err := r.db.Debug().Select("id", "username").Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}

	return nil
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
