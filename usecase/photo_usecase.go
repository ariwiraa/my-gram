package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/google/uuid"
)

type PhotoUsecase interface {
	Create(payload dtos.PhotoRequest, userId uint, fileLocation string) (domain.Photo, error)
	GetById(id string) (domain.Photo, error)
	GetAll() ([]domain.Photo, error)
	GetAllPhotosByUserId(userId uint) ([]domain.Photo, error)
	Update(payload dtos.UpdatePhotoRequest, id string, userId uint) (domain.Photo, error)
	Delete(id string)
}

type photoUsecase struct {
	photoRepository     repository.PhotoRepository
	commentRepository   repository.CommentRepository
	tagRepository       repository.TagRepository
	photoTagsRepository repository.PhotoTagsRepository
}

// Create implements PhotoUsecase
func (u *photoUsecase) Create(payload dtos.PhotoRequest, userId uint, fileLocation string) (domain.Photo, error) {
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

/*// GetById implements PhotoUsecase
func (u *photoUsecase) GetById(id string) (domain.Photo, error) {
	photo, err := u.photoRepository.FindById(id)
	if err != nil {
		return photo, err
	}

	totalCommentsCh := make(chan int64)
	photoTagsCh := make(chan []domain.PhotoTags)

	var wg sync.WaitGroup

	// Goroutine untuk mengambil total komentar
	wg.Add(1)
	go func() {
		defer wg.Done()
		totalComments, _ := u.commentRepository.CountCommentsByPhotoId(photo.ID)
		totalCommentsCh <- totalComments
	}()

	// Goroutine untuk mengambil tag
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(id)
	//	if err != nil {
	//		photoTagsCh <- []domain.PhotoTags{}
	//	} else {
	//		photoTagsCh <- photoTags
	//	}
	//}()

	// Tunggu hingga kedua goroutine selesai
	wg.Wait()

	// Ambil hasil dari channel totalCommentsCh dan set ke photo.TotalComment
	totalComments := <-totalCommentsCh
	photo.TotalComment = totalComments

	// Ambil hasil dari channel photoTagsCh dan tambahkan ke photo.Tags
	photoTags := <-photoTagsCh
	for _, photoTag := range photoTags {
		tag, err := u.tagRepository.FindById(photoTag.TagId)
		if err != nil {
			return photo, err
		}
		photo.Tags = append(photo.Tags, *tag)
	}
	log.Println(photo.Tags)

	return photo, nil
}*/

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
func (u *photoUsecase) Update(payload dtos.UpdatePhotoRequest, id string, userId uint) (domain.Photo, error) {
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

func NewPhotoUsecase(photo repository.PhotoRepository, comment repository.CommentRepository, tag repository.TagRepository, photoTags repository.PhotoTagsRepository) PhotoUsecase {
	return &photoUsecase{
		photoRepository:     photo,
		commentRepository:   comment,
		tagRepository:       tag,
		photoTagsRepository: photoTags,
	}
}
