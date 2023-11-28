package usecase

import (
	"errors"
	"github.com/ariwiraa/my-gram/domain/dtos/request"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/repository"
)

type UserUsecase interface {
	Register(payload request.UserRegister) (domain.User, error)
	Login(payload request.UserLogin) (domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// Login implements UserUsecase
func (u *userUsecase) Login(payload request.UserLogin) (domain.User, error) {
	username := payload.Username
	password := payload.Password

	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, err
	}

	comparePassword := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePassword {
		return user, err
	}

	return user, nil
}

// Register implements UserUsecase
func (u *userUsecase) Register(payload request.UserRegister) (domain.User, error) {
	var user domain.User

	isEmailUsed, _ := u.userRepository.IsEmailExists(payload.Email)
	if isEmailUsed {
		return user, errors.New("email sudah digunakan")
	}

	isUsernameUsed, _ := u.userRepository.IsUsernameExists(payload.Username)
	if isUsernameUsed {
		return user, errors.New("username sudah digunakan")
	}

	hashingPassword := helpers.HashPass(payload.Password)

	user = domain.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashingPassword,
	}

	newUser, err := u.userRepository.AddUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
