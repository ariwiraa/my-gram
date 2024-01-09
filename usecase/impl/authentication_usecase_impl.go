package impl

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type authenticationUsecaseImpl struct {
	repo            repository.AuthenticationRepository
	userRepository  repository.UserRepository
	redisRepository repository.RedisRepository
}

func NewAuthenticationUsecaseImpl(repo repository.AuthenticationRepository, userRepository repository.UserRepository, redisRepository repository.RedisRepository) usecase.AuthenticationUsecase {
	return &authenticationUsecaseImpl{
		repo:            repo,
		userRepository:  userRepository,
		redisRepository: redisRepository,
	}
}

// ResendEmail implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) ResendEmail(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		log.Printf("[ResendEmail, FindByEmail] with error detail %v", err.Error())
		return err
	}

	token := helpers.GenerateRandomOTP()
	tokenString := strconv.Itoa(token)

	configMail := helpers.DataMail{
		Username: user.Username,
		Email:    user.Email,
		Token:    tokenString,
		Subject:  "Your verification Email",
	}

	err = helpers.Mail(&configMail).Send()
	if err != nil {
		log.Printf("[ResendEmail, Mail] with error detail %v", err.Error())
		return helpers.ErrorFailedSendEmail
	}

	err = u.redisRepository.Set(ctx, user.Email, tokenString, 5*time.Minute)
	if err != nil {
		log.Printf("[ResendEmail, Set] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil

}

// VerifyEmail implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) VerifyEmail(ctx context.Context, email, token string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		log.Printf("[VerifyEmail, FindByEmail] with error detail %v", err.Error())
		return err
	}

	value, err := u.redisRepository.Get(ctx, user.Email)
	if err != nil {
		log.Printf("[VerifyEmail, Get] with error detail %v", err.Error())
		return helpers.ErrLinkExpired
	}

	if value == token && token != "" {
		currentTime := time.Now()
		user.EmailVerificationAt = &currentTime
	}

	err = u.userRepository.UpdateUser(ctx, *user)
	if err != nil {
		log.Printf("[VerifyEmail, UpdateUser] with error detail %v", err.Error())
		return err
	}

	return nil

}

func (u *authenticationUsecaseImpl) Register(ctx context.Context, payload request.UserRegister) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	isEmailUsed, err := u.userRepository.IsEmailExists(ctx, payload.Email)
	if err != nil {
		log.Printf("[Register, IsEmailExists] with error detail %v", err.Error())
		return &user, err
	}

	if isEmailUsed {
		return &user, helpers.ErrEmailAlreadyUserd
	}

	isUsernameUsed, err := u.userRepository.IsUsernameExists(ctx, payload.Username)
	if err != nil {
		log.Printf("[Register, IsUsernameExist] with error detail %v", err.Error())
		return &user, nil
	}

	if isUsernameUsed {
		return &user, helpers.ErrorUsernameAlreadyUsed
	}

	hashingPassword := helpers.HashPass(payload.Password)

	user = domain.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashingPassword,
	}

	newUser, err := u.userRepository.AddUser(ctx, user)
	if err != nil {
		log.Printf("[VerifyEmail, AddUser] with error detail %v", err.Error())
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

	err = helpers.Mail(&configMail).Send()
	if err != nil {
		log.Printf("[Register, Mail] with error detail %v", err.Error())
		return &newUser, helpers.ErrFailedSendEmail
	}

	// TODO: Simpan token ke redis dengan TTL 5 menit
	err = u.redisRepository.Set(ctx, newUser.Email, tokenString, 5*time.Minute)
	if err != nil {
		log.Printf("[VerifyEmail, Set] with error detail %v", err.Error())
		return &newUser, helpers.ErrRepository
	}

	return &newUser, nil
}

func (u *authenticationUsecaseImpl) Login(ctx context.Context, payload request.UserLogin) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByUsername(ctx, payload.Username)
	if err != nil {
		log.Printf("[Login, FindByUsername] with error detail %v", err.Error())
		return &user, err
	}

	if user.EmailVerificationAt == nil {
		return &user, helpers.ErrEmailNotVerified
	}

	comparePassword := helpers.ComparePass([]byte(user.Password), []byte(payload.Password))
	if !comparePassword {
		log.Printf("[Login, ComparePass] with error detail %v", err.Error())
		return &user, helpers.ErrPasswordNotMatch
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
		log.Printf("[Add, Add] with error detail %v", err.Error())
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
		log.Printf("[Delete, FindByRefreshToken] with error detail %v", err.Error())
		return err
	}

	err = u.repo.Delete(ctx, *authentication)
	if err != nil {
		log.Printf("[Delete, Delete] with error detail %v", err.Error())
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) ExistsByRefreshToken(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := u.repo.FindByRefreshToken(ctx, token)
	if err != nil {
		log.Printf("[ExistsByRefreshToken, FindByRefreshToken] with error detail %v", err.Error())
		return err
	}

	return nil
}
