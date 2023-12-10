package repository

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, comment domain.Comment) (*domain.Comment, error)
	FindById(ctx context.Context, id uint) (*domain.Comment, error)
	FindAllCommentsByPhotoId(ctx context.Context, photoId string) ([]domain.Comment, error)
	Update(ctx context.Context, comment domain.Comment, id uint) (*domain.Comment, error)
	Delete(ctx context.Context, id uint)
	CountCommentsByPhotoId(ctx context.Context, photoId string) (int64, error)
}
