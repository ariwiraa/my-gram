package impl

import (
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type authenticationUsecaseImpl struct {
	repo           repository.AuthenticationRepository
	userRepository repository.UserRepository
}

func (u *authenticationUsecaseImpl) Register(payload request.UserRegister) (*domain.User, error) {
	var user domain.User

	isEmailUsed, _ := u.userRepository.IsEmailExists(payload.Email)
	if isEmailUsed {
		return &user, errors.New("email sudah digunakan")
	}

	isUsernameUsed, _ := u.userRepository.IsUsernameExists(payload.Username)
	if isUsernameUsed {
		return &user, errors.New("username sudah digunakan")
	}

	hashingPassword := helpers.HashPass(payload.Password)

	user = domain.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashingPassword,
	}

	newUser, err := u.userRepository.AddUser(user)
	if err != nil {
		return &newUser, err
	}

	return &newUser, nil
}

func (u *authenticationUsecaseImpl) Login(payload request.UserLogin) (*domain.User, error) {
	username := payload.Username
	password := payload.Password

	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return &user, err
	}

	if user.ID == 0 {
		return &user, err
	}

	comparePassword := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePassword {
		return &user, err
	}

	return &user, nil
}

// Add implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) Add(token string) error {
	authentication := new(domain.Authentication)

	authentication.RefreshToken = token

	err := u.repo.Add(*authentication)
	if err != nil {
		return err
	}

	return nil

}

// Delete implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) Delete(token string) error {
	authentication, err := u.repo.FindByRefreshToken(token)
	if err != nil {
		return err
	}

	err = u.repo.Delete(*authentication)
	if err != nil {
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) ExistsByRefreshToken(token string) error {
	_, err := u.repo.FindByRefreshToken(token)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthenticationUsecaseImpl(repo repository.AuthenticationRepository, userRepository repository.UserRepository) usecase.AuthenticationUsecase {
	return &authenticationUsecaseImpl{
		repo:           repo,
		userRepository: userRepository,
	}
}
