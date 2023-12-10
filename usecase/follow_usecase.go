package usecase

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
)

type FollowUsecase interface {
	FollowUser(ctx context.Context, followRequest request.FollowRequest) (string, error)
	GetFollowersByUsername(ctx context.Context, username string) ([]domain.User, error)
	GetFollowingsByUsername(ctx context.Context, username string) ([]domain.User, error)
}
