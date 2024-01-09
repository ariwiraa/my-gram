package impl

import (
	"context"
	"errors"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type userLikesPhotoRepository struct {
	db *gorm.DB
}

func (r *userLikesPhotoRepository) FindUserWhoLiked(ctx context.Context, userId uint) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Preload("LikedPhotos").Where("id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, helpers.ErrUserNotFound
		}
		log.Printf("[FindUserWhoLiked] with error detail %v", err.Error())
		return &user, helpers.ErrRepository
	}

	return &user, nil
}

func (r *userLikesPhotoRepository) FindPhotoWhoLiked(ctx context.Context, photoId string) (*domain.Photo, error) {
	var photo domain.Photo
	err := r.db.WithContext(ctx).Preload("LikedBy").Where("id = ?", photoId).First(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &photo, helpers.ErrPhotoNotFound
		}
		log.Printf("[FindPhotoWhoLiked] with error detail %v", err.Error())
		return &photo, helpers.ErrRepository
	}

	return &photo, err
}

// CountLikesPhotoById implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) CountUsersWhoLikedPhotoByPhotoId(ctx context.Context, photoId string) (int64, error) {
	var totalLikes int64
	err := r.db.WithContext(ctx).Model(&domain.UserLikesPhoto{}).Where("photo_id = ?", photoId).Count(&totalLikes).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return totalLikes, helpers.ErrPhotoNotFound
		}
		log.Printf("[FindUserWhoLiked] with error detail %v", err.Error())
		return totalLikes, helpers.ErrRepository
	}

	return totalLikes, nil
}

// DeleteLike implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) DeleteLike(ctx context.Context, photoId string, userId uint) {
	var userLikesPhoto domain.UserLikesPhoto

	err := r.db.WithContext(ctx).Where("photo_id = ? AND user_id = ?", photoId, userId).Delete(&userLikesPhoto).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		log.Printf("[Delete] with error detail %v", err.Error())
		return
	}
}

// InsertLike implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) InsertLike(ctx context.Context, userLikesPhoto domain.UserLikesPhoto) error {
	err := r.db.WithContext(ctx).Create(&userLikesPhoto).Error
	if err != nil {
		log.Printf("[InsertLike] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil
}

// VerifyUserLike implements UserLikesPhotoRepository
func (r *userLikesPhotoRepository) VerifyUserLike(ctx context.Context, photoId string, userId uint) (bool, error) {
	var userLikesPhoto domain.UserLikesPhoto

	err := r.db.WithContext(ctx).Where("photo_id = ? AND user_id = ?", photoId, userId).First(&userLikesPhoto).Error
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
