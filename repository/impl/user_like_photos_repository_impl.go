package impl

import (
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type userLikesPhotoRepository struct {
	db *gorm.DB
}

func (r *userLikesPhotoRepository) FindUserWhoLiked(userId uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("LikedPhotos").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *userLikesPhotoRepository) FindPhotoWhoLiked(photoId string) (*domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Debug().Preload("LikedBy").Where("id = ?", photoId).First(&photo).Error
	if err != nil {
		return &photo, err
	}

	return &photo, err
}

// CountLikesPhotoById implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) CountUsersWhoLikedPhotoByPhotoId(photoId string) (int64, error) {
	var totalLikes int64
	err := r.db.Model(&domain.UserLikesPhoto{}).Where("photo_id = ?", photoId).Count(&totalLikes).Error
	if err != nil {
		return 0, err
	}

	return totalLikes, nil
}

// DeleteLike implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) DeleteLike(photoId string, userId uint) {
	var userLikesPhoto domain.UserLikesPhoto

	err := r.db.Debug().Where("photo_id = ? AND user_id = ?", photoId, userId).Delete(&userLikesPhoto).Error
	if err != nil {
		return
	}
}

// InsertLike implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) InsertLike(userLikesPhoto domain.UserLikesPhoto) error {
	err := r.db.Debug().Create(&userLikesPhoto).Error
	if err != nil {
		return err
	}

	return nil
}

// VerifyUserLike implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) VerifyUserLike(photoId string, userId uint) (bool, error) {
	var userLikesPhoto domain.UserLikesPhoto

	err := r.db.Debug().Where("photo_id = ? AND user_id = ?", photoId, userId).First(&userLikesPhoto).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewUserLikesPhotoRepository(db *gorm.DB) repository.UserLikesPhotoRepository {
	return &userLikesPhotoRepository{db: db}
}
