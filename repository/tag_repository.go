package repository

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type TagRepository interface {
	Add(ctx context.Context, tag domain.Tag) (*domain.Tag, error)
	FindByName(ctx context.Context, name string) ([]domain.Tag, error)
	FindById(ctx context.Context, id uint) (*domain.Tag, error)
	AddTagIfNotExists(ctx context.Context, name string) (*domain.Tag, error)
}
