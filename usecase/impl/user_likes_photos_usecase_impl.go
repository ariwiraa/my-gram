package impl

import (
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type userLikesPhotosUsecase struct {
	likesRepository repository.UserLikesPhotoRepository
	photoRepository repository.PhotoRepository
	userRepository  repository.UserRepository
}

func (u *userLikesPhotosUsecase) GetPhotosLikedByUserId(userId uint) ([]domain.Photo, error) {
	user, err := u.likesRepository.FindUserWhoLiked(userId)
	if err != nil {
		return []domain.Photo{}, err
	}

	likedPhotos := user.LikedPhotos

	var photos []domain.Photo
	for _, likedPhoto := range likedPhotos {
		photo, err := u.photoRepository.FindById(likedPhoto.ID)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return photos, nil
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

func NewUserLikesPhotosUsecase(likesRepository repository.UserLikesPhotoRepository, photoRepository repository.PhotoRepository, userRepository repository.UserRepository) usecase.UserLikesPhotosUsecase {
	return &userLikesPhotosUsecase{likesRepository: likesRepository, photoRepository: photoRepository, userRepository: userRepository}
}

func (u *userLikesPhotosUsecase) GetUsersWhoLikedPhotoByPhotoId(photoId string) ([]domain.User, error) {
	err := u.photoRepository.IsPhotoExist(photoId)
	if err != nil {
		return []domain.User{}, err
	}

	photo, err := u.likesRepository.FindPhotoWhoLiked(photoId)
	if err != nil {
		return []domain.User{}, err
	}

	likedUsers := photo.LikedBy

	var users []domain.User
	for _, likedUser := range likedUsers {
		user, err := u.userRepository.FindById(likedUser.ID)
		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return users, nil
}
