package usecase

import (
	"context"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/domain/dtos/response"
)

type PhotoUsecase interface {
	Create(ctx context.Context, payload request.PhotoRequest, userId uint) (*response.PhotoResponse, error)
	GetById(ctx context.Context, id string) (*response.PhotoResponse, error)
	GetAll(ctx context.Context) ([]domain.Photo, error)
	GetAllPhotosByUserId(ctx context.Context, userId uint) ([]domain.Photo, error)
	Update(ctx context.Context, payload request.UpdatePhotoRequest, id string, userId uint) (*response.PhotoResponse, error)
	Delete(ctx context.Context, id string) error
}
