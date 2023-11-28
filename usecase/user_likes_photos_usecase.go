package usecase

import (
	"errors"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
)

type UserLikesPhotosUsecase interface {
	LikeThePhoto(photoId string, userId uint) (string, error)
}

type userLikesPhotosUsecase struct {
	likesRepository repository.UserLikesPhotoRepository
	photoRepository repository.PhotoRepository
}

// LikeThePhoto implements UserLikesPhotosUsecase
func (u *userLikesPhotosUsecase) LikeThePhoto(photoId string, userId uint) (string, error) {
	err := u.photoRepository.IsPhotoExist(photoId)
	if err != nil {
		return "", errors.New("foto tidak tersedia")
	}

	userLike, _ := u.likesRepository.VerifyUserLike(photoId, userId)

	likes := domain.UserLikesPhoto{
		PhotoId: photoId,
		UserId:  userId,
	}

	var message string
	if !userLike {
		u.likesRepository.InsertLike(likes)
		message = "Berhasil menyukai foto"
	} else {
		u.likesRepository.DeleteLike(likes.PhotoId, likes.UserId)
		message = "Gagal menyukai foto"
	}

	return message, nil

}

func NewUserLikesPhotosUsecase(likesRepository repository.UserLikesPhotoRepository, photoRepository repository.PhotoRepository) UserLikesPhotosUsecase {
	return &userLikesPhotosUsecase{likesRepository: likesRepository, photoRepository: photoRepository}
}
