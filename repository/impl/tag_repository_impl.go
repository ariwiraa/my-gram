package impl

import (
	"context"
	"errors"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type tagRepositoryImpl struct {
	db *gorm.DB
}

// FindById implements repository.TagRepository.
func (r *tagRepositoryImpl) FindById(ctx context.Context, id uint) (*domain.Tag, error) {
	var tag domain.Tag
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, errors.New("tag not found")
	}

	return &tag, nil
}

func (r *tagRepositoryImpl) AddTagIfNotExists(ctx context.Context, name string) (*domain.Tag, error) {
	var existingTag domain.Tag
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&existingTag).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Tag dengan name tersebut belum ada, tambahkan tag baru
			newTag := domain.Tag{Name: name}
			newTagSaved, err := r.Add(ctx, newTag)
			if err != nil {
				return nil, err
			}
			// Kembalikan tag baru yang telah dibuat
			return newTagSaved, nil
		}
	}
	// Tag sudah ada dalam database, kembalikan tag yang ada
	return &existingTag, nil
}

// Add implements repository.TagRepository.
func (r *tagRepositoryImpl) Add(ctx context.Context, tag domain.Tag) (*domain.Tag, error) {
	err := r.db.WithContext(ctx).Create(&tag).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// FindByName implements repository.TagRepository.
func (r *tagRepositoryImpl) FindByName(ctx context.Context, name string) ([]domain.Tag, error) {
	var tags []domain.Tag
	err := r.db.WithContext(ctx).Where("lower(name) LIKE lower(?)", "%"+name+"%").Find(&tags).Error
	if err != nil {
		return nil, errors.New("tag Not Found")
	}

	return tags, nil
}

func NewTagRepositoryImpl(db *gorm.DB) repository.TagRepository {
	return &tagRepositoryImpl{db: db}
}
