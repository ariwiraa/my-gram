package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
)

type UserLikesPhotosUsecase interface {
	LikeThePhoto(photoId string, userId uint) (string, error)
	GetUsersWhoLikedPhotoByPhotoId(photoId string) ([]domain.User, error)
	GetPhotosLikedByUserId(userId uint) ([]domain.Photo, error)
}
