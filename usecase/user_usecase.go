package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/repository"
)

type UserUsecase interface {
	Register(payload domain.UserRequest) (domain.User, error)
	Login(payload domain.UserLogin) (domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// Login implements UserUsecase
func (u *userUsecase) Login(payload domain.UserLogin) (domain.User, error) {
	email := payload.Email
	password := payload.Password

	user, err := u.userRepository.FindByEmail(email)
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
func (u *userUsecase) Register(payload domain.UserRequest) (domain.User, error) {

	user := domain.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Age:      payload.Age,
	}

	newUser, err := u.userRepository.AddUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}
