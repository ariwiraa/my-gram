package repository

import (
	"errors"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo domain.Photo) (domain.Photo, error)
	FindById(id string) (domain.Photo, error)
	FindAll() ([]domain.Photo, error)
	FindByUserId(id uint) ([]domain.Photo, error)
	FindByIdAndByUserId(id string, userId uint) (*domain.Photo, error)
	Update(photo domain.Photo, id string) (domain.Photo, error)
	Delete(photo domain.Photo)
	IsPhotoExist(id string) error
}

type photoRepository struct {
	db *gorm.DB
}

// IsPhotoExist implements PhotoRepository
func (r *photoRepository) IsPhotoExist(id string) error {
	var photo domain.Photo
	err := r.db.Debug().First(&photo, "id = ?", id).Error
	if err != nil {
		return err
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
		log.Fatalln("error deleting data", err)
		return
	}
}

// FindAll implements PhotoRepository
func (r *photoRepository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo

	err := r.db.Debug().Find(&photos).Error
	if err != nil {
		log.Fatal("error getting all photos data: ", err)
	}
	return photos, nil
}

// FindById implements PhotoRepository
func (r *photoRepository) FindById(id string) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Debug().Preload("Comments").First(&photo, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("photo not found")
		}
		log.Fatal("error getting data :", err)
	}

	return photo, err
}

// Update implements PhotoRepository
func (r *photoRepository) Update(photo domain.Photo, id string) (domain.Photo, error) {

	err := r.db.Debug().Model(&photo).Where("id = ?", id).Updates(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
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

func (r *photoRepository) FindByUserId(id uint) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Debug().Preload("Comments").Find(&photos, "user_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("photo not found")
		}
		log.Fatal("error getting data :", err)
	}

	return photos, err
}
