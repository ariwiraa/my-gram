package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
)

type AuthenticationUsecase interface {
	Add(token string) error
	ExistsByRefreshToken(token string) error
	Delete(token string) error
	Register(payload request.UserRegister) (*domain.User, error)
	Login(payload request.UserLogin) (*domain.User, error)
}
