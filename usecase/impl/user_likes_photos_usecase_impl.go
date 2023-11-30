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

	var photoIds []string
	for _, likedPhoto := range likedPhotos {
		photoIds = append(photoIds, likedPhoto.ID)
	}
	// Function FindPhotosByIDList menggunakan IN bukan WHERE
	// Karena IN bisa mengambil semua id dengan satu kali call database daripada where yg harus berkali kali
	// Jadi harus dihindari call database di dalam loop
	photos, err := u.photoRepository.FindPhotosByIDList(photoIds)

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

	var userIds []uint
	for _, likedUser := range likedUsers {
		userIds = append(userIds, likedUser.ID)
	}

	users, err := u.userRepository.FindUsersByIDList(userIds)
	if err != nil {
		return users, err
	}

	return users, nil
}
