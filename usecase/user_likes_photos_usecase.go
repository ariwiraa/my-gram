package usecase

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type UserLikesPhotosUsecase interface {
	LikeThePhoto(ctx context.Context, photoId string, userId uint) (string, error)
	GetUsersWhoLikedPhotoByPhotoId(ctx context.Context, photoId string) ([]domain.User, error)
	GetPhotosLikedByUserId(ctx context.Context, userId uint) ([]domain.Photo, error)
}
