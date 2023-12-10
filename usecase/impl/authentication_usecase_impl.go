package impl

import (
	"context"
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
	"time"
)

type authenticationUsecaseImpl struct {
	repo           repository.AuthenticationRepository
	userRepository repository.UserRepository
}

func (u *authenticationUsecaseImpl) Register(ctx context.Context, payload request.UserRegister) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	isEmailUsed, _ := u.userRepository.IsEmailExists(ctx, payload.Email)
	if isEmailUsed {
		return &user, errors.New("email sudah digunakan")
	}

	isUsernameUsed, _ := u.userRepository.IsUsernameExists(ctx, payload.Username)
	if isUsernameUsed {
		return &user, errors.New("username sudah digunakan")
	}

	hashingPassword := helpers.HashPass(payload.Password)

	user = domain.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashingPassword,
	}

	newUser, err := u.userRepository.AddUser(ctx, user)
	if err != nil {
		return &newUser, err
	}

	return &newUser, nil
}

func (u *authenticationUsecaseImpl) Login(ctx context.Context, payload request.UserLogin) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	username := payload.Username
	password := payload.Password

	user, err := u.userRepository.FindByUsername(ctx, username)
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
func (u *authenticationUsecaseImpl) Add(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	authentication := new(domain.Authentication)

	authentication.RefreshToken = token

	err := u.repo.Add(ctx, *authentication)
	if err != nil {
		return err
	}

	return nil

}

// Delete implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) Delete(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	authentication, err := u.repo.FindByRefreshToken(ctx, token)
	if err != nil {
		return err
	}

	err = u.repo.Delete(ctx, *authentication)
	if err != nil {
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) ExistsByRefreshToken(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := u.repo.FindByRefreshToken(ctx, token)
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
