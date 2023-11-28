package repository

import "github.com/ariwiraa/my-gram/domain"

type TagRepository interface {
	Add(tag domain.Tag) (*domain.Tag, error)
	FindByName(name string) ([]domain.Tag, error)
	FindById(id uint) (*domain.Tag, error)
	AddTagIfNotExists(name string) (*domain.Tag, error)
}
