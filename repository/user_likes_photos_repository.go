package repository

import (
	"github.com/ariwiraa/my-gram/domain"
)

type UserLikesPhotoRepository interface {
	InsertLike(userLikesPhoto domain.UserLikesPhoto) error
	DeleteLike(photoId string, userId uint)
	VerifyUserLike(photoId string, userId uint) (bool, error)
	FindPhotoWhoLiked(photoId string) (*domain.Photo, error)
	FindUserWhoLiked(userId uint) (*domain.User, error)
	CountUsersWhoLikedPhotoByPhotoId(photoId string) (int64, error)
}
