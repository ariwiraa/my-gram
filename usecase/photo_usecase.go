package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
)

type PhotoUsecase interface {
	Create(payload request.PhotoRequest, userId uint, fileLocation string) (domain.Photo, error)
	GetById(id string) (domain.Photo, error)
	GetAll() ([]domain.Photo, error)
	GetAllPhotosByUserId(userId uint) ([]domain.Photo, error)
	Update(payload request.UpdatePhotoRequest, id string, userId uint) (domain.Photo, error)
	Delete(id string)
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
