package impl

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/google/uuid"
)

type photoUsecase struct {
	photoRepository     repository.PhotoRepository
	commentRepository   repository.CommentRepository
	tagRepository       repository.TagRepository
	photoTagsRepository repository.PhotoTagsRepository
}

// Create implements PhotoUsecase
func (u *photoUsecase) Create(payload request.PhotoRequest, userId uint, fileLocation string) (domain.Photo, error) {
	photo := domain.Photo{
		ID:       uuid.NewString(),
		Caption:  payload.Caption,
		PhotoUrl: fileLocation,
		UserId:   userId,
	}

	newPhoto, err := u.photoRepository.Create(photo)
	if err != nil {
		return newPhoto, err
	}

	for _, photoTag := range payload.Tags {
		tag := domain.Tag{Name: photoTag}

		newTag, err := u.tagRepository.AddTagIfNotExists(tag.Name)
		if err != nil {
			return newPhoto, err
		}

		photoTags := domain.PhotoTags{
			PhotoId: newPhoto.ID,
			TagId:   newTag.ID,
		}

		err = u.photoTagsRepository.Add(photoTags)
		if err != nil {
			return newPhoto, err
		}
		// Hubungkan tag dengan foto
		newPhoto.Tags = append(newPhoto.Tags, *newTag)

	}

	return newPhoto, nil
}

// Delete implements PhotoUsecase
func (u *photoUsecase) Delete(id string) {
	photo, err := u.photoRepository.FindById(id)
	if err != nil {
		return
	}

	u.photoRepository.Delete(photo)
	err = u.photoTagsRepository.Delete(photo.ID)
	if err != nil {
		return
	}
}

// GetAll implements PhotoUsecase
func (u *photoUsecase) GetAll() ([]domain.Photo, error) {
	photos, err := u.photoRepository.FindAll()
	if err != nil {
		return photos, err
	}

	for _, photo := range photos {
		totalComments, _ := u.commentRepository.CountCommentsByPhotoId(photo.ID)
		photo.TotalComment = totalComments

		photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(photo.ID)
		// Jika tidak ada tag, maka return photo
		if err != nil {
			return photos, nil
		}

		for _, photoTag := range photoTags {
			tag, err := u.tagRepository.FindById(photoTag.TagId)
			if err != nil {
				return photos, err
			}

			photo.Tags = append(photo.Tags, *tag)
		}
	}

	return photos, nil
}

func (u *photoUsecase) GetById(id string) (domain.Photo, error) {
	photo, err := u.photoRepository.FindById(id)
	if err != nil {
		return photo, err
	}
	//TODO: Nanti gunakan goroutine untuk mengambil total comments dan tags.
	totalComments, _ := u.commentRepository.CountCommentsByPhotoId(photo.ID)
	photo.TotalComment = totalComments

	photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(id)
	// Jika tidak ada tag, maka return photo
	if err != nil {
		return photo, nil
	}

	for _, photoTag := range photoTags {
		tag, err := u.tagRepository.FindById(photoTag.TagId)
		if err != nil {
			return photo, err
		}

		photo.Tags = append(photo.Tags, *tag)

	}

	return photo, nil
}

// Update implements PhotoUsecase
func (u *photoUsecase) Update(payload request.UpdatePhotoRequest, id string, userId uint) (domain.Photo, error) {
	photo, err := u.photoRepository.FindByIdAndByUserId(id, userId)
	if err != nil {
		panic(err)
	}

	photo.Caption = payload.Caption

	err = u.photoTagsRepository.Delete(id)
	if err != nil {
		return domain.Photo{}, err
	}

	updatedPhoto, err := u.photoRepository.Update(*photo, id)
	if err != nil {
		return updatedPhoto, err
	}

	for _, photoTag := range payload.Tags {
		tag := domain.Tag{Name: photoTag}

		newTag, err := u.tagRepository.AddTagIfNotExists(tag.Name)
		if err != nil {
			return updatedPhoto, err
		}

		photoTags := domain.PhotoTags{
			PhotoId: updatedPhoto.ID,
			TagId:   newTag.ID,
		}

		err = u.photoTagsRepository.Add(photoTags)
		if err != nil {
			return updatedPhoto, err
		}
		// Hubungkan tag dengan foto
		updatedPhoto.Tags = append(updatedPhoto.Tags, *newTag)

	}

	return updatedPhoto, nil
}

func (u *photoUsecase) GetAllPhotosByUserId(userId uint) ([]domain.Photo, error) {
	photos, err := u.photoRepository.FindByUserId(userId)
	if err != nil {
		return photos, err
	}

	for _, photo := range photos {
		totalComments, _ := u.commentRepository.CountCommentsByPhotoId(photo.ID)
		photo.TotalComment = totalComments

		photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(photo.ID)
		// Jika tidak ada tag, maka return photo
		if err != nil {
			return photos, nil
		}

		for _, photoTag := range photoTags {
			tag, err := u.tagRepository.FindById(photoTag.TagId)
			if err != nil {
				return photos, err
			}

			photo.Tags = append(photo.Tags, *tag)

		}
	}

	return photos, nil
}

func NewPhotoUsecase(photo repository.PhotoRepository, comment repository.CommentRepository, tag repository.TagRepository, photoTags repository.PhotoTagsRepository) usecase.PhotoUsecase {
	return &photoUsecase{
		photoRepository:     photo,
		commentRepository:   comment,
		tagRepository:       tag,
		photoTagsRepository: photoTags,
	}
}
