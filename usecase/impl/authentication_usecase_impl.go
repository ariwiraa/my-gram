package impl

import (
	"context"
	"errors"
	"strconv"
	"time"

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

func NewAuthenticationUsecaseImpl(repo repository.AuthenticationRepository, userRepository repository.UserRepository) usecase.AuthenticationUsecase {
	return &authenticationUsecaseImpl{
		repo:           repo,
		userRepository: userRepository,
	}
}

// ResendEmail implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) ResendEmail(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return errors.New("email is not found")
	}

	token := helpers.GenerateRandomOTP()
	tokenString := strconv.Itoa(token)

	configMail := helpers.DataMail{
		Username: user.Username,
		Email:    user.Email,
		Token:    tokenString,
		Subject:  "Your verification Email",
	}

	// TODO: Simpan token ke redis dengan TTL 5 menit

	err = helpers.Mail(&configMail).Send()
	if err != nil {
		return errors.New("failed send email")
	}

	return nil

}

// VerifyEmail implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) VerifyEmail(ctx context.Context, email, token string) error {
	// TODO: Periksa apakah token sudah lebih dari 5 menit menggunakan redis
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return errors.New("email is not found")
	}

	currentTime := time.Now()
	user.EmailVerificationAt = &currentTime

	err = u.userRepository.UpdateUser(ctx, *user)
	if err != nil {
		return err
	}

	return nil

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

	token := helpers.GenerateRandomOTP()
	tokenString := strconv.Itoa(token)

	configMail := helpers.DataMail{
		Username: newUser.Username,
		Email:    newUser.Email,
		Token:    tokenString,
		Subject:  "Your verification Email",
	}
	// TODO: Simpan token ke redis dengan TTL 5 menit

	err = helpers.Mail(&configMail).Send()
	if err != nil {
		return &newUser, errors.New("failed send email")
	}

	return &newUser, nil
}

func (u *authenticationUsecaseImpl) Login(ctx context.Context, payload request.UserLogin) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByUsername(ctx, payload.Username)
	if err != nil {
		return &user, err
	}

	if user.EmailVerificationAt == nil {
		return &user, errors.New("email not verified. Please verif your email first")
	}

	comparePassword := helpers.ComparePass([]byte(user.Password), []byte(payload.Password))
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
