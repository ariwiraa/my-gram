package repository

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type PhotoTagsRepository interface {
	Add(ctx context.Context, photoTags domain.PhotoTags) error
	Delete(ctx context.Context, photoId string) error
	FindPhotoTagsByPhotoId(ctx context.Context, photoId string) ([]domain.PhotoTags, error)
}
