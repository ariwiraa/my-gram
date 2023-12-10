package usecase

import (
	"context"
	"github.com/ariwiraa/my-gram/domain/dtos/response"
)

type UserUsecase interface {
	GetUserProfileByUsername(ctx context.Context, username string) (*response.UserProfileResponse, error)
}
