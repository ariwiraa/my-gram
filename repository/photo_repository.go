package repository

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type PhotoRepository interface {
	Create(ctx context.Context, photo domain.Photo) (domain.Photo, error)
	FindById(ctx context.Context, id string) (domain.Photo, error)
	FindAll(ctx context.Context) ([]domain.Photo, error)
	FindByUserId(ctx context.Context, id uint) ([]domain.Photo, error)
	FindByIdAndByUserId(ctx context.Context, id string, userId uint) (*domain.Photo, error)
	Update(ctx context.Context, photo domain.Photo, id string) (domain.Photo, error)
	Delete(ctx context.Context, photo domain.Photo)
	IsPhotoExist(ctx context.Context, id string) error
	FindPhotosByIDList(ctx context.Context, photoIds []string) ([]domain.Photo, error)
	CountPhotoByUserId(ctx context.Context, userId uint) (int64, error)
}
