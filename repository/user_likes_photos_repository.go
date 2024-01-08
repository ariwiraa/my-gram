package repository

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type UserLikesPhotoRepository interface {
	InsertLike(ctx context.Context, userLikesPhoto domain.UserLikesPhoto) error
	DeleteLike(ctx context.Context, photoId string, userId uint)
	VerifyUserLike(ctx context.Context, photoId string, userId uint) (bool, error)
	FindPhotoWhoLiked(ctx context.Context, photoId string) (*domain.Photo, error)
	FindUserWhoLiked(ctx context.Context, userId uint) (*domain.User, error)
	CountUsersWhoLikedPhotoByPhotoId(ctx context.Context, photoId string) (int64, error)
}
