package impl

import (
	"context"
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

// CountCommentsByPhotoId implements CommentRepository
func (r *commentRepository) CountCommentsByPhotoId(ctx context.Context, photoId string) (int64, error) {
	var totalComment int64
	err := r.db.WithContext(ctx).Model(&domain.Comment{}).Where("photo_id = ?", photoId).Count(&totalComment).Error
	if err != nil {
		return 0, err
	}

	return totalComment, nil
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db: db}
}

// Create implements CommentRepository
func (r *commentRepository) Create(ctx context.Context, comment domain.Comment) (*domain.Comment, error) {
	err := r.db.WithContext(ctx).Create(&comment).Error
	if err != nil {
		return &comment, err
	}

	return &comment, nil
}

// Delete implements CommentRepository
func (r *commentRepository) Delete(ctx context.Context, id uint) {
	var comment domain.Comment

	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return
	}
}

// FindAll implements CommentRepository
func (r *commentRepository) FindAllCommentsByPhotoId(ctx context.Context, photoId string) ([]domain.Comment, error) {
	var comments []domain.Comment

	err := r.db.WithContext(ctx).Find(&comments, "photo_id = ?", photoId).Error
	if err != nil {
		return comments, errors.New("photo doesn't exists")
	}
	return comments, nil
}

// FindById implements CommentRepository
func (r *commentRepository) FindById(ctx context.Context, id uint) (*domain.Comment, error) {
	var comment domain.Comment
	err := r.db.WithContext(ctx).First(&comment, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &comment, errors.New("comment doesn't exists")
		}
	}

	return &comment, err
}

// Update implements CommentRepository
func (r *commentRepository) Update(ctx context.Context, comment domain.Comment, id uint) (*domain.Comment, error) {

	err := r.db.WithContext(ctx).Model(&comment).Where("id = ?", id).Updates(&comment).Error
	if err != nil {
		return &comment, err
	}

	return &comment, nil
}
