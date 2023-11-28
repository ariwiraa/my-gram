package repository

import (
	"github.com/ariwiraa/my-gram/domain"
)

type CommentRepository interface {
	Create(comment domain.Comment) (domain.Comment, error)
	FindById(id uint) (domain.Comment, error)
	FindAll() ([]domain.Comment, error)
	Update(comment domain.Comment, id uint) (domain.Comment, error)
	Delete(id uint)
	CountCommentsByPhotoId(photoId string) (int64, error)
}
