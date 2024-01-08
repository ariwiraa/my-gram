package usecase

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
)

type CommentUsecase interface {
	Create(ctx context.Context, payload request.CommentRequest) (*domain.Comment, error)
	GetById(ctx context.Context, id uint, photoId string) (*domain.Comment, error)
	GetAllCommentsByPhotoId(ctx context.Context, photoId string) ([]domain.Comment, error)
	Update(ctx context.Context, payload request.CommentRequest, id uint) (*domain.Comment, error)
	Delete(ctx context.Context, id uint, photoId string)
}
