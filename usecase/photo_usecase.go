package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
)

type PhotoUsecase interface {
	Create(payload domain.PhotoRequest, userId uint) (domain.Photo, error)
	GetById(id uint) (domain.Photo, error)
	GetAll() ([]domain.Photo, error)
	Update(payload domain.PhotoRequest, id uint, userId uint) (domain.Photo, error)
	Delete(photo domain.Photo)
}

type photoUsecase struct {
	photoRepository repository.PhotoRepository
}

// Create implements PhotoUsecase
func (u *photoUsecase) Create(payload domain.PhotoRequest, userId uint) (domain.Photo, error) {
	photo := domain.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserId:   userId,
	}

	newPhoto, err := u.photoRepository.Create(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}

// Delete implements PhotoUsecase
func (u *photoUsecase) Delete(photo domain.Photo) {
	u.photoRepository.Delete(photo.ID)
}

// GetAll implements PhotoUsecase
func (u *photoUsecase) GetAll() ([]domain.Photo, error) {
	photos, err := u.photoRepository.FindAll()
	if err != nil {
		return photos, err
	}

	return photos, nil
}

// GetById implements PhotoUsecase
func (u *photoUsecase) GetById(id uint) (domain.Photo, error) {
	photo, err := u.photoRepository.FindById(id)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

// Update implements PhotoUsecase
func (u *photoUsecase) Update(payload domain.PhotoRequest, id uint, userId uint) (domain.Photo, error) {
	photo, err := u.photoRepository.FindById(id)
	if err != nil {
		panic(err)
	}

	// var photo domain.Photo

	photo.Title = payload.Title
	photo.Caption = payload.Caption
	photo.PhotoUrl = payload.PhotoUrl
	photo.UserId = userId

	updatedPhoto, err := u.photoRepository.Update(photo, id)
	if err != nil {
		return updatedPhoto, err
	}

	return updatedPhoto, nil
}

func NewPhotoUsecase(photo repository.PhotoRepository) PhotoUsecase {
	return &photoUsecase{photoRepository: photo}
}
