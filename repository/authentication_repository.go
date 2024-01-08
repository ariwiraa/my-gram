package repository

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type AuthenticationRepository interface {
	Add(ctx context.Context, authentication domain.Authentication) error
	FindByRefreshToken(ctx context.Context, token string) (*domain.Authentication, error)
	Delete(ctx context.Context, authentication domain.Authentication) error
}
