package repository

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type FollowRepository interface {
	Save(ctx context.Context, follow domain.Follow) error
	Delete(ctx context.Context, follow domain.Follow) error
	VerifyUserFollow(ctx context.Context, follow domain.Follow) (bool, error)
	CountFollowerByUserId(ctx context.Context, userId uint) (int64, error)
	CountFollowingByUserId(ctx context.Context, userId uint) (int64, error)
	FindFollowersByUserId(ctx context.Context, userId uint) ([]domain.User, error)
	FindFollowingByUserId(ctx context.Context, userId uint) ([]domain.User, error)
}
