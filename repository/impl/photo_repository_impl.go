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

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repository.PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) FindPhotosByIDList(ctx context.Context, photoIds []string) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.WithContext(ctx).Preload("User").Preload("Comments").Find(&photos, "id IN ?", photoIds).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photos, helpers.ErrPhotoNotFound
		}
		log.Printf("[FindPhotosByIdList] with error detail %v", err.Error())
		return photos, helpers.ErrRepository
	}
	return photos, err
}

func (r *photoRepository) FindByUserId(ctx context.Context, id uint) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.WithContext(ctx).Preload("Comments").Find(&photos, "user_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photos, helpers.ErrUserNotFound
		}
		log.Printf("[FindById] with error detail %v", err.Error())
		return photos, helpers.ErrRepository
	}

	return photos, nil
}

func (r *photoRepository) CountPhotoByUserId(ctx context.Context, userId uint) (int64, error) {
	var totalPosts int64
	err := r.db.WithContext(ctx).Model(&domain.Photo{}).Where("user_id = ?", userId).Count(&totalPosts).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return totalPosts, helpers.ErrUserNotFound
		}
		log.Printf("[CountPhotoByUserId] with error detail %v", err.Error())
		return totalPosts, helpers.ErrRepository
	}

	return totalPosts, nil
}

// IsPhotoExist implements PhotoRepository
func (r *photoRepository) IsPhotoExist(ctx context.Context, id string) error {
	var photo domain.Photo
	err := r.db.WithContext(ctx).First(&photo, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ErrPhotoNotFound
		}
		log.Printf("[IsPhotoExist] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil
}

// Create implements PhotoRepository
func (r *photoRepository) Create(ctx context.Context, photo domain.Photo) (domain.Photo, error) {
	err := r.db.WithContext(ctx).Create(&photo).Error
	if err != nil {
		log.Printf("[Create] with error detail %v", err.Error())
		return photo, helpers.ErrRepository
	}

	return photo, nil
}

// Delete implements PhotoRepository
func (r *photoRepository) Delete(ctx context.Context, photo domain.Photo) error {
	err := r.db.WithContext(ctx).Where("id = ?", photo.ID).Delete(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ErrPhotoNotFound
		}
		log.Printf("[IsPhotoExist] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil
}

// FindAll implements PhotoRepository
func (r *photoRepository) FindAll(ctx context.Context) ([]domain.Photo, error) {
	var photos []domain.Photo

	err := r.db.WithContext(ctx).Find(&photos).Error
	if err != nil {
		log.Printf("[FindAll] with error detail %v", err.Error())
		return photos, helpers.ErrRepository
	}
	return photos, nil
}

// FindById implements PhotoRepository
func (r *photoRepository) FindById(ctx context.Context, id string) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.WithContext(ctx).Preload("User").Preload("Comments").First(&photo, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photo, err
		}
		log.Printf("[FindById] with error detail %v", err.Error())
		return photo, helpers.ErrRepository
	}

	return photo, nil
}

// Update implements PhotoRepository
func (r *photoRepository) Update(ctx context.Context, photo domain.Photo, id string) (domain.Photo, error) {

	err := r.db.WithContext(ctx).Model(&photo).Where("id = ?", id).Updates(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photo, helpers.ErrPhotoNotFound
		}
		log.Printf("[IsPhotoExist] with error detail %v", err.Error())
		return photo, helpers.ErrRepository
	}

	return photo, nil
}

func (r *photoRepository) FindByIdAndByUserId(ctx context.Context, id string, userId uint) (*domain.Photo, error) {
	var photo domain.Photo
	err := r.db.WithContext(ctx).Preload("Comments").First(&photo, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &photo, helpers.ErrPhotoNotFound
		}
		log.Printf("[IsPhotoExist] with error detail %v", err.Error())
		return &photo, helpers.ErrRepository
	}

	return &photo, nil
}
