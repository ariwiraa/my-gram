package repository

import "github.com/ariwiraa/my-gram/domain"

type PhotoTagsRepository interface {
	Add(photoTags domain.PhotoTags) error
	Delete(photoId string) error
	FindPhotoTagsByPhotoId(photoId string) ([]domain.PhotoTags, error)
}
