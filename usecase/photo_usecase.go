package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/domain/dtos/response"
)

type PhotoUsecase interface {
	Create(payload request.PhotoRequest, userId uint, fileLocation string) (*response.PhotoResponse, error)
	GetById(id string) (*response.PhotoResponse, error)
	GetAll() ([]domain.Photo, error)
	GetAllPhotosByUserId(userId uint) ([]domain.Photo, error)
	Update(payload request.UpdatePhotoRequest, id string, userId uint) (*response.PhotoResponse, error)
	Delete(id string)
}
