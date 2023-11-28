package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
)

type CommentUsecase interface {
	Create(payload request.CommentRequest, userId uint) (domain.Comment, error)
	GetById(id uint) (domain.Comment, error)
	GetAll() ([]domain.Comment, error)
	Update(payload request.CommentRequest, id uint, userId uint) (domain.Comment, error)
	Delete(comment domain.Comment)
}
