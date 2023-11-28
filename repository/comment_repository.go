package repository

import (
	"errors"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment domain.Comment) (domain.Comment, error)
	FindById(id uint) (domain.Comment, error)
	FindAll() ([]domain.Comment, error)
	Update(comment domain.Comment, id uint) (domain.Comment, error)
	Delete(id uint)
	CountCommentsByPhotoId(photoId string) (int64, error)
}

type commentRepository struct {
	db *gorm.DB
}

// CountCommentsByPhotoId implements CommentRepository
func (r *commentRepository) CountCommentsByPhotoId(photoId string) (int64, error) {
	var totalComment int64
	err := r.db.Model(&domain.Comment{}).Where("photo_id = ?", photoId).Count(&totalComment).Error
	if err != nil {
		return 0, err
	}

	return totalComment, nil
}

// Create implements CommentRepository
func (r *commentRepository) Create(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Debug().Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

// Delete implements CommentRepository
func (r *commentRepository) Delete(id uint) {
	var comment domain.Comment

	err := r.db.Debug().Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		log.Fatalln("error deleting data", err)
		return
	}
}

// FindAll implements CommentRepository
func (r *commentRepository) FindAll() ([]domain.Comment, error) {
	var comments []domain.Comment

	err := r.db.Debug().Find(&comments).Error
	if err != nil {
		log.Fatal("error getting all comments data: ", err)
	}
	return comments, nil
}

// FindById implements CommentRepository
func (r *commentRepository) FindById(id uint) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Debug().First(&comment, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("comment not found")
		}
		log.Fatal("error getting data :", err)
	}

	return comment, err
}

// Update implements CommentRepository
func (r *commentRepository) Update(comment domain.Comment, id uint) (domain.Comment, error) {

	err := r.db.Debug().Model(&comment).Where("id = ?", id).Updates(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}
