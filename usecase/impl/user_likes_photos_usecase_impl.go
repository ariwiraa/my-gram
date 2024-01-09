package impl

import (
	"context"
	"log"
	"time"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type userLikesPhotosUsecase struct {
	likesRepository repository.UserLikesPhotoRepository
	photoRepository repository.PhotoRepository
	userRepository  repository.UserRepository
}

func NewUserLikesPhotosUsecase(likesRepository repository.UserLikesPhotoRepository, photoRepository repository.PhotoRepository, userRepository repository.UserRepository) usecase.UserLikesPhotosUsecase {
	return &userLikesPhotosUsecase{likesRepository: likesRepository, photoRepository: photoRepository, userRepository: userRepository}
}

func (u *userLikesPhotosUsecase) GetPhotosLikedByUserId(ctx context.Context, userId uint) ([]domain.Photo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.likesRepository.FindUserWhoLiked(ctx, userId)
	if err != nil {
		log.Printf("[GetPhotosLikedByUserId, FindUserWhoLiked] with error detail %v", err.Error())
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
	photos, err := u.photoRepository.FindPhotosByIDList(ctx, photoIds)
	if err != nil {
		log.Printf("[GetPhotosLikedByUserId, FindPhotosByIDList] with error detail %v", err.Error())
		return photos, err
	}

	return photos, nil
}

// LikeThePhoto implements UserLikesPhotosUsecase
func (u *userLikesPhotosUsecase) LikeThePhoto(ctx context.Context, photoId string, userId uint) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := u.photoRepository.IsPhotoExist(ctx, photoId)
	if err != nil {
		log.Printf("[LikeThePhoto, IsPhotoExist] with error detail %v", err.Error())
		return "", err
	}

	userLike, _ := u.likesRepository.VerifyUserLike(ctx, photoId, userId)

	likes := domain.UserLikesPhoto{
		PhotoId: photoId,
		UserId:  userId,
	}

	var message string
	if !userLike {
		u.likesRepository.InsertLike(ctx, likes)
		message = "Berhasil menyukai foto"
	} else {
		u.likesRepository.DeleteLike(ctx, likes.PhotoId, likes.UserId)
		message = "Gagal menyukai foto"
	}

	return message, nil

}

func (u *userLikesPhotosUsecase) GetUsersWhoLikedPhotoByPhotoId(ctx context.Context, photoId string) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := u.photoRepository.IsPhotoExist(ctx, photoId)
	if err != nil {
		log.Printf("[GetUsersWhoLikedPhotoByPhotoId, IsPhotoExist] with error detail %v", err.Error())
		return []domain.User{}, err
	}

	photo, err := u.likesRepository.FindPhotoWhoLiked(ctx, photoId)
	if err != nil {
		log.Printf("[GetUsersWhoLikedPhotoByPhotoId, FindPhotoWhoLiked] with error detail %v", err.Error())
		return []domain.User{}, err
	}

	likedUsers := photo.LikedBy

	var userIds []uint
	for _, likedUser := range likedUsers {
		userIds = append(userIds, likedUser.ID)
	}

	users, err := u.userRepository.FindUsersByIDList(ctx, userIds)
	if err != nil {
		log.Printf("[GetUsersWhoLikedPhotoByPhotoId, FindUsersByIDList] with error detail %v", err.Error())
		return users, err
	}

	return users, nil
}
