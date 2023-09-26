package repository

import (
	"errors"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type UserLikesPhotoRepository interface {
	InsertLike(userLikesPhoto domain.UserLikesPhoto) error
	DeleteLike(photoId, userId uint)
	VerifyUserLike(photoId, userId uint) (bool, error)
	CountLikesPhotoById(photoId uint) (uint, error)
	GetLikedPhotosByUserId(userId uint) ([]domain.Photo, error)
}

type userLikesPhotoRepository struct {
	db *gorm.DB
}

// CountLikesPhotoById implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) CountLikesPhotoById(photoId uint) (uint, error) {
	var totalLikes int64
	err := r.db.Model(&domain.UserLikesPhoto{}).Where("photo_id = ?", photoId).Count(&totalLikes).Error
	if err != nil {
		return 0, err
	}

	return uint(totalLikes), nil
}

// DeleteLike implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) DeleteLike(photoId uint, userId uint) {
	var userLikesPhoto domain.UserLikesPhoto

	err := r.db.Debug().Where("photo_id = ? AND user_id = ?", photoId, userId).Delete(&userLikesPhoto).Error
	if err != nil {
		log.Fatal("error deleting data", err)
		return
	}
}

// GetLikedPhotoByUserId implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) GetLikedPhotosByUserId(userId uint) ([]domain.Photo, error) {
	var user domain.User

	err := r.db.Preload("LikedPhotos").Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []domain.Photo{}, nil
		}

		return nil, err
	}

	likedPhotos := user.LikedPhotos

	var photos []domain.Photo
	for _, likedPhoto := range likedPhotos {
		var photo domain.Photo

		err := r.db.Preload("User").Preload("Comments").Where("id = ?", likedPhoto.ID).First(&photo).Error
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return photos, nil
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
func (r *userLikesPhotoRepository) VerifyUserLike(photoId uint, userId uint) (bool, error) {
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

func NewUserLikesPhotoRepository(db *gorm.DB) UserLikesPhotoRepository {
	return &userLikesPhotoRepository{db: db}
}
