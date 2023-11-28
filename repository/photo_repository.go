package repository

import (
	"github.com/ariwiraa/my-gram/domain"
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
