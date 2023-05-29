package usecase

import (
	"errors"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos"
	"github.com/ariwiraa/my-gram/repository"
)

type CommentUsecase interface {
	Create(payload dtos.CommentRequest, userId uint) (domain.Comment, error)
	GetById(id uint) (domain.Comment, error)
	GetAll() ([]domain.Comment, error)
	Update(payload dtos.CommentRequest, id uint, userId uint) (domain.Comment, error)
	Delete(comment domain.Comment)
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

// Create implements CommentUsecase
func (u *commentUsecase) Create(payload dtos.CommentRequest, userId uint) (domain.Comment, error) {
	var comment domain.Comment

	err := u.photoRepository.IsPhotoExist(payload.PhotoId)
	if err != nil {
		return comment, errors.New("photo tidak ada")
	}

	comment = domain.Comment{
		Message: payload.Message,
		PhotoId: payload.PhotoId,
		UserId:  userId,
	}

	newComment, err := u.commentRepository.Create(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

// Delete implements CommentUsecase
func (u *commentUsecase) Delete(comment domain.Comment) {
	u.commentRepository.Delete(comment.ID)
}

// GetAll implements CommentUsecase
func (u *commentUsecase) GetAll() ([]domain.Comment, error) {
	comments, err := u.commentRepository.FindAll()
	if err != nil {
		return comments, err
	}

	return comments, nil
}

// GetById implements CommentUsecase
func (u *commentUsecase) GetById(id uint) (domain.Comment, error) {
	comment, err := u.commentRepository.FindById(id)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

// Update implements CommentUsecase
func (u *commentUsecase) Update(payload dtos.CommentRequest, id uint, userId uint) (domain.Comment, error) {
	comment, err := u.commentRepository.FindById(id)
	if err != nil {
		panic(err)
	}

	comment.Message = payload.Message
	comment.PhotoId = payload.PhotoId
	comment.UserId = userId

	updatedComment, err := u.commentRepository.Update(comment, id)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func NewCommentUsecase(comment repository.CommentRepository, photoRepository repository.PhotoRepository) CommentUsecase {
	return &commentUsecase{
		commentRepository: comment,
		photoRepository:   photoRepository,
	}
}
