package impl

import (
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

func (r *photoRepository) FindPhotosByIDList(photoIds []string) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Debug().Preload("User").Preload("Comments").Find(&photos, "id IN ?", photoIds).Error
	return photos, err
}

func (r *photoRepository) FindByUserId(id uint) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Debug().Preload("Comments").Find(&photos, "user_id = ?", id).Error
	if err != nil {
		return photos, err
	}

	return photos, nil
}

type photoRepository struct {
	db *gorm.DB
}

func (r *photoRepository) CountPhotoByUserId(userId uint) (int64, error) {
	var totalPosts int64
	err := r.db.Model(&domain.Photo{}).Where("user_id = ?", userId).Count(&totalPosts).Error
	if err != nil {
		return 0, err
	}

	return totalPosts, nil
}

// IsPhotoExist implements PhotoRepository
func (r *photoRepository) IsPhotoExist(id string) error {
	var photo domain.Photo
	err := r.db.Debug().First(&photo, "id = ?", id).Error
	if err != nil {
		return errors.New("id photo doesn't exists")
	}

	return nil
}

// Create implements PhotoRepository
func (r *photoRepository) Create(photo domain.Photo) (domain.Photo, error) {
	err := r.db.Debug().Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

// Delete implements PhotoRepository
func (r *photoRepository) Delete(photo domain.Photo) {
	err := r.db.Debug().Where("id = ?", photo.ID).Delete(&photo).Error
	if err != nil {
		return
	}
}

// FindAll implements PhotoRepository
func (r *photoRepository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo

	err := r.db.Debug().Find(&photos).Error
	if err != nil {
		return photos, err
	}
	return photos, nil
}

// FindById implements PhotoRepository
func (r *photoRepository) FindById(id string) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Debug().Preload("User").Preload("Comments").First(&photo, "id = ?", id).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

// Update implements PhotoRepository
func (r *photoRepository) Update(photo domain.Photo, id string) (domain.Photo, error) {

	err := r.db.Debug().Model(&photo).Where("id = ?", id).Updates(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func NewPhotoRepository(db *gorm.DB) repository.PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) FindByIdAndByUserId(id string, userId uint) (*domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Preload("Comments").First(&photo, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		return &photo, err
	}

	return &photo, nil
}
