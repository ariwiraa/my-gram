package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
)

type CommentUsecase interface {
	Create(payload request.CommentRequest) (*domain.Comment, error)
	GetById(id uint, photoId string) (*domain.Comment, error)
	GetAllCommentsByPhotoId(photoId string) ([]domain.Comment, error)
	Update(payload request.CommentRequest, id uint) (*domain.Comment, error)
	Delete(id uint, photoId string)
}
