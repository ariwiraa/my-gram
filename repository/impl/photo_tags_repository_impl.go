package impl

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type photoTagsRepositoryImpl struct {
	db *gorm.DB
}

// FindPhotoTagsByPhotoId implements repository.PhotoTagsRepository.
func (r *photoTagsRepositoryImpl) FindPhotoTagsByPhotoId(photoId string) ([]domain.PhotoTags, error) {
	var photoTags []domain.PhotoTags
	err := r.db.Find(&photoTags, "photo_id = ?", photoId).Error
	if err != nil {
		return nil, err
	}

	return photoTags, nil
}

// Add implements repository.PhotoTagsRepository.
func (r *photoTagsRepositoryImpl) Add(photoTags domain.PhotoTags) error {
	err := r.db.Create(&photoTags).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete implements repository.PhotoTagsRepository.
func (r *photoTagsRepositoryImpl) Delete(photoId string) error {
	var photoTags domain.PhotoTags
	err := r.db.Where("photo_id = ?", photoId).Delete(&photoTags).Error
	if err != nil {
		return err
	}

	return nil
}

func NewPhotoTagsRepositoryImpl(db *gorm.DB) repository.PhotoTagsRepository {
	return &photoTagsRepositoryImpl{db: db}
}
