package impl

import (
	"context"
	"log"
	"time"

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
func (u *commentUsecase) Create(ctx context.Context, payload request.CommentRequest) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var comment domain.Comment

	err := u.photoRepository.IsPhotoExist(ctx, payload.PhotoId)
	if err != nil {
		log.Printf("[Create, IsPhotoExist] with error detail %v", err.Error())
		return &comment, err
	}

	comment = domain.Comment{
		Message: payload.Message,
		PhotoId: payload.PhotoId,
		UserId:  payload.UserId,
	}

	newComment, err := u.commentRepository.Create(ctx, comment)
	if err != nil {
		log.Printf("[Create, Create] with error detail %v", err.Error())
		return newComment, err
	}

	return newComment, nil
}

// Delete implements CommentUsecase
func (u *commentUsecase) Delete(ctx context.Context, id uint, photoId string) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := u.photoRepository.IsPhotoExist(ctx, photoId)
	if err != nil {
		log.Printf("[Delete, IsPhotoExist] with error detail %v", err.Error())
		return
	}

	comment, err := u.commentRepository.FindById(ctx, id)
	if err != nil {
		log.Printf("[Delete, FindById] with error detail %v", err.Error())
		return
	}

	u.commentRepository.Delete(ctx, comment.ID)
}

// GetAll implements CommentUsecase
func (u *commentUsecase) GetAllCommentsByPhotoId(ctx context.Context, photoId string) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := u.photoRepository.IsPhotoExist(ctx, photoId)
	if err != nil {
		log.Printf("[GetAllCommentsByPhotoId, IsPhotoExist] with error detail %v", err.Error())
		return nil, err
	}

	comments, err := u.commentRepository.FindAllCommentsByPhotoId(ctx, photoId)
	if err != nil {
		log.Printf("[GetAllCommentsByPhotoId, FindAllCommentsByPhotoId] with error detail %v", err.Error())
		return comments, err
	}

	return comments, nil
}

// GetById implements CommentUsecase
func (u *commentUsecase) GetById(ctx context.Context, id uint, photoId string) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := u.photoRepository.IsPhotoExist(ctx, photoId)
	if err != nil {
		log.Printf("[GetById, IsPhotoExist] with error detail %v", err.Error())
		return &domain.Comment{}, err
	}

	comment, err := u.commentRepository.FindById(ctx, id)
	if err != nil {
		log.Printf("[GetById, FindById] with error detail %v", err.Error())
		return comment, err
	}

	return comment, nil
}

// Update implements CommentUsecase
func (u *commentUsecase) Update(ctx context.Context, payload request.CommentRequest, id uint) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	comment, err := u.commentRepository.FindById(ctx, id)
	if err != nil {
		log.Printf("[Update, FindById] with error detail %v", err.Error())
		return comment, err
	}

	comment.Message = payload.Message

	updatedComment, err := u.commentRepository.Update(ctx, *comment, id)
	if err != nil {
		log.Printf("[Update, Update] with error detail %v", err.Error())
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
