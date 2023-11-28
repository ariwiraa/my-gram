package impl

import (
	"errors"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type tagRepositoryImpl struct {
	db *gorm.DB
}

// FindById implements repository.TagRepository.
func (r *tagRepositoryImpl) FindById(id uint) (*domain.Tag, error) {
	var tag domain.Tag
	err := r.db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, errors.New("tag not found")
	}

	return &tag, nil
}

func (r *tagRepositoryImpl) AddTagIfNotExists(name string) (*domain.Tag, error) {
	var existingTag domain.Tag
	err := r.db.Where("name = ?", name).First(&existingTag).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Tag dengan name tersebut belum ada, tambahkan tag baru
			newTag := domain.Tag{Name: name}
			newTagSaved, err := r.Add(newTag)
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
func (r *tagRepositoryImpl) Add(tag domain.Tag) (*domain.Tag, error) {
	err := r.db.Create(&tag).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// FindByName implements repository.TagRepository.
func (r *tagRepositoryImpl) FindByName(name string) ([]domain.Tag, error) {
	var tags []domain.Tag
	err := r.db.Where("lower(name) LIKE lower(?)", "%"+name+"%").Find(&tags).Error
	if err != nil {
		return nil, errors.New("tag Not Found")
	}

	return tags, nil
}

func NewTagRepositoryImpl(db *gorm.DB) repository.TagRepository {
	return &tagRepositoryImpl{db: db}
}
