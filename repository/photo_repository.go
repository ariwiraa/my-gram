package repository

import (
	"errors"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo domain.Photo) (domain.Photo, error)
	FindById(id uint) (domain.Photo, error)
	FindAll() ([]domain.Photo, error)
	Update(photo domain.Photo, id uint) (domain.Photo, error)
	Delete(id uint)
	IsPhotoExist(id uint) error
}

type photoRepository struct {
	db *gorm.DB
}

// IsPhotoExist implements PhotoRepository
func (r *photoRepository) IsPhotoExist(id uint) error {
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
func (r *photoRepository) Delete(id uint) {
	var photo domain.Photo

	err := r.db.Debug().Where("id = ?", id).Delete(&photo).Error
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
func (r *photoRepository) FindById(id uint) (domain.Photo, error) {
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
func (r *photoRepository) Update(photo domain.Photo, id uint) (domain.Photo, error) {

	err := r.db.Debug().Model(&photo).Where("id = ?", id).Updates(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db: db}
}
