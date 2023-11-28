package impl

import (
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type commentUsecase struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

// Create implements CommentUsecase
func (u *commentUsecase) Create(payload request.CommentRequest) (*domain.Comment, error) {
	var comment domain.Comment

	err := u.photoRepository.IsPhotoExist(payload.PhotoId)
	if err != nil {
		return &comment, errors.New("photo tidak ada")
	}

	comment = domain.Comment{
		Message: payload.Message,
		PhotoId: payload.PhotoId,
		UserId:  payload.UserId,
	}

	newComment, err := u.commentRepository.Create(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

// Delete implements CommentUsecase
func (u *commentUsecase) Delete(id uint, photoId string) {
	err := u.photoRepository.IsPhotoExist(photoId)
	if err != nil {
		return
	}

	comment, err := u.commentRepository.FindById(id)
	if err != nil {
		return
	}

	u.commentRepository.Delete(comment.ID)
}

// GetAll implements CommentUsecase
func (u *commentUsecase) GetAllCommentsByPhotoId(photoId string) ([]domain.Comment, error) {
	err := u.photoRepository.IsPhotoExist(photoId)
	if err != nil {
		return nil, err
	}

	comments, err := u.commentRepository.FindAllCommentsByPhotoId(photoId)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

// GetById implements CommentUsecase
func (u *commentUsecase) GetById(id uint, photoId string) (*domain.Comment, error) {
	err := u.photoRepository.IsPhotoExist(photoId)
	if err != nil {
		return &domain.Comment{}, err
	}

	comment, err := u.commentRepository.FindById(id)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

// Update implements CommentUsecase
func (u *commentUsecase) Update(payload request.CommentRequest, id uint) (*domain.Comment, error) {
	comment, err := u.commentRepository.FindById(id)
	if err != nil {
		return comment, err
	}

	comment.Message = payload.Message

	updatedComment, err := u.commentRepository.Update(*comment, id)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func NewCommentUsecase(comment repository.CommentRepository, photoRepository repository.PhotoRepository) usecase.CommentUsecase {
	return &commentUsecase{
		commentRepository: comment,
		photoRepository:   photoRepository,
	}
}
