package impl

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type photoTagsRepositoryImpl struct {
	db *gorm.DB
}

// FindPhotoTagsByPhotoId implements repository.PhotoTagsRepository.
func (r *photoTagsRepositoryImpl) FindPhotoTagsByPhotoId(ctx context.Context, photoId string) ([]domain.PhotoTags, error) {
	var photoTags []domain.PhotoTags
	err := r.db.WithContext(ctx).Find(&photoTags, "photo_id = ?", photoId).Error
	if err != nil {
		return nil, err
	}

	return photoTags, nil
}

// Add implements repository.PhotoTagsRepository.
func (r *photoTagsRepositoryImpl) Add(ctx context.Context, photoTags domain.PhotoTags) error {
	err := r.db.WithContext(ctx).Create(&photoTags).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete implements repository.PhotoTagsRepository.
func (r *photoTagsRepositoryImpl) Delete(ctx context.Context, photoId string) error {
	var photoTags domain.PhotoTags
	err := r.db.WithContext(ctx).Where("photo_id = ?", photoId).Delete(&photoTags).Error
	if err != nil {
		return err
	}

	return nil
}

func NewPhotoTagsRepositoryImpl(db *gorm.DB) repository.PhotoTagsRepository {
	return &photoTagsRepositoryImpl{db: db}
}
