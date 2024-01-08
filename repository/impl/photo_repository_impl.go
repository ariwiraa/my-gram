package impl

import (
	"context"
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

func (r *photoRepository) FindPhotosByIDList(ctx context.Context, photoIds []string) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.WithContext(ctx).Preload("User").Preload("Comments").Find(&photos, "id IN ?", photoIds).Error
	return photos, err
}

func (r *photoRepository) FindByUserId(ctx context.Context, id uint) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.WithContext(ctx).Preload("Comments").Find(&photos, "user_id = ?", id).Error
	if err != nil {
		return photos, err
	}

	return photos, nil
}

type photoRepository struct {
	db *gorm.DB
}

func (r *photoRepository) CountPhotoByUserId(ctx context.Context, userId uint) (int64, error) {
	var totalPosts int64
	err := r.db.WithContext(ctx).Model(&domain.Photo{}).Where("user_id = ?", userId).Count(&totalPosts).Error
	if err != nil {
		return 0, err
	}

	return totalPosts, nil
}

// IsPhotoExist implements PhotoRepository
func (r *photoRepository) IsPhotoExist(ctx context.Context, id string) error {
	var photo domain.Photo
	err := r.db.WithContext(ctx).First(&photo, "id = ?", id).Error
	if err != nil {
		return errors.New("id photo doesn't exists")
	}

	return nil
}

// Create implements PhotoRepository
func (r *photoRepository) Create(ctx context.Context, photo domain.Photo) (domain.Photo, error) {
	err := r.db.WithContext(ctx).Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

// Delete implements PhotoRepository
func (r *photoRepository) Delete(ctx context.Context, photo domain.Photo) {
	err := r.db.WithContext(ctx).Where("id = ?", photo.ID).Delete(&photo).Error
	if err != nil {
		return
	}
}

// FindAll implements PhotoRepository
func (r *photoRepository) FindAll(ctx context.Context) ([]domain.Photo, error) {
	var photos []domain.Photo

	err := r.db.WithContext(ctx).Find(&photos).Error
	if err != nil {
		return photos, err
	}
	return photos, nil
}

// FindById implements PhotoRepository
func (r *photoRepository) FindById(ctx context.Context, id string) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.WithContext(ctx).Preload("User").Preload("Comments").First(&photo, "id = ?", id).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

// Update implements PhotoRepository
func (r *photoRepository) Update(ctx context.Context, photo domain.Photo, id string) (domain.Photo, error) {

	err := r.db.WithContext(ctx).Model(&photo).Where("id = ?", id).Updates(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func NewPhotoRepository(db *gorm.DB) repository.PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) FindByIdAndByUserId(ctx context.Context, id string, userId uint) (*domain.Photo, error) {
	var photo domain.Photo
	err := r.db.WithContext(ctx).Preload("Comments").First(&photo, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		return &photo, err
	}

	return &photo, nil
}
