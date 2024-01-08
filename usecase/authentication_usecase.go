package usecase

import (
	"context"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
)

type AuthenticationUsecase interface {
	Add(ctx context.Context, token string) error
	ExistsByRefreshToken(ctx context.Context, token string) error
	Delete(ctx context.Context, token string) error
	Register(ctx context.Context, payload request.UserRegister) (*domain.User, error)
	Login(ctx context.Context, payload request.UserLogin) (*domain.User, error)
	VerifyEmail(ctx context.Context, email, token string) error
	ResendEmail(ctx context.Context, email string) error
}
