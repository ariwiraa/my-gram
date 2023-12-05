package impl

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/domain/dtos/response"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/google/uuid"
	"log"
)

type photoUsecase struct {
	photoRepository          repository.PhotoRepository
	commentRepository        repository.CommentRepository
	tagRepository            repository.TagRepository
	photoTagsRepository      repository.PhotoTagsRepository
	userLikesPhotoRepository repository.UserLikesPhotoRepository
	userRepository           repository.UserRepository
}

// Create implements PhotoUsecase
func (u *photoUsecase) Create(payload request.PhotoRequest, userId uint, fileLocation string) (*response.PhotoResponse, error) {
	photo := domain.Photo{
		ID:       uuid.NewString(),
		Caption:  payload.Caption,
		PhotoUrl: fileLocation,
		UserId:   userId,
	}

	usernameCh := make(chan string)
	go u.fetchUsername(userId, usernameCh)

	newPhoto, err := u.photoRepository.Create(photo)
	if err != nil {
		log.Println("error create photo: ", err)
		return &response.PhotoResponse{}, err
	}

	responsePhoto := response.PhotoResponse{
		Id:        newPhoto.ID,
		Caption:   newPhoto.Caption,
		PhotoUrl:  newPhoto.PhotoUrl,
		CreatedAt: newPhoto.CreatedAt,
		Username:  <-usernameCh,
	}

	for _, photoTag := range payload.Tags {
		tag := domain.Tag{Name: photoTag}

		newTag, err := u.tagRepository.AddTagIfNotExists(tag.Name)
		if err != nil {
			log.Println("error create tag: ", err)
			return &response.PhotoResponse{}, err
		}

		photoTags := domain.PhotoTags{
			PhotoId: newPhoto.ID,
			TagId:   newTag.ID,
		}

		err = u.photoTagsRepository.Add(photoTags)
		if err != nil {
			log.Println("Failed to associate tag with photo: ", err)
			return &response.PhotoResponse{}, err
		}
		// Hubungkan tag dengan foto
		responsePhoto.PhotoTags = append(responsePhoto.PhotoTags, photoTag)
		log.Println("Tag added succesfully")
	}

	log.Println("photo create succesfully")
	return &responsePhoto, nil
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

func (u *photoUsecase) GetById(id string) (*response.PhotoResponse, error) {
	photo, err := u.photoRepository.FindById(id)
	if err != nil {
		log.Printf("Error fetching photo by ID: %v", err)
		return &response.PhotoResponse{}, err
	}

	// Menggunakan goroutine untuk mengambil total comments dan total likes secara bersamaan
	totalCommentsCh := make(chan int64)
	totalLikesCh := make(chan int64)

	go u.calculateTotalComments(photo.ID, totalCommentsCh)
	go u.calculateTotalLikes(photo.ID, totalLikesCh)

	totalComments := <-totalCommentsCh
	totalLikes := <-totalLikesCh

	photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(id)
	if err != nil {
		log.Printf("Error fetching photo tags: %v", err)
		return &response.PhotoResponse{}, err
	}

	responsePhoto := response.PhotoResponse{
		Id:            photo.ID,
		PhotoUrl:      photo.PhotoUrl,
		Caption:       photo.Caption,
		CreatedAt:     photo.CreatedAt,
		Username:      photo.User.Username,
		TotalComments: totalComments,
		TotalLikes:    totalLikes,
	}

	u.processPhotoTags(photoTags, &responsePhoto)

	return &responsePhoto, nil
}

// Update implements PhotoUsecase
func (u *photoUsecase) Update(payload request.UpdatePhotoRequest, id string, userId uint) (*response.PhotoResponse, error) {
	photo, err := u.photoRepository.FindByIdAndByUserId(id, userId)
	if err != nil {
		log.Printf("Error fetching photo by id and user id %v", err)
		return &response.PhotoResponse{}, err
	}

	photo.Caption = payload.Caption

	err = u.photoTagsRepository.Delete(id)
	if err != nil {
		log.Printf("Error deleting association photo tag %v", err)
		return &response.PhotoResponse{}, err
	}

	updatedPhoto, err := u.photoRepository.Update(*photo, id)
	if err != nil {
		log.Printf("Error updating photo %v", err)
		return &response.PhotoResponse{}, err
	}

	totalCommentsCh := make(chan int64)
	totalLikesCh := make(chan int64)
	usernameCh := make(chan string)

	go u.fetchUsername(userId, usernameCh)
	go u.calculateTotalComments(photo.ID, totalCommentsCh)
	go u.calculateTotalLikes(photo.ID, totalLikesCh)

	totalComments := <-totalCommentsCh
	totalLikes := <-totalLikesCh
	username := <-usernameCh

	responsePhoto := response.PhotoResponse{
		Id:            updatedPhoto.ID,
		PhotoUrl:      updatedPhoto.PhotoUrl,
		Caption:       updatedPhoto.Caption,
		CreatedAt:     updatedPhoto.CreatedAt,
		Username:      username,
		TotalComments: totalComments,
		TotalLikes:    totalLikes,
	}

	for _, photoTag := range payload.Tags {
		tag := domain.Tag{Name: photoTag}

		newTag, err := u.tagRepository.AddTagIfNotExists(tag.Name)
		if err != nil {
			log.Printf("Error adding tag %v", err)
			return &responsePhoto, err
		}

		photoTags := domain.PhotoTags{
			PhotoId: updatedPhoto.ID,
			TagId:   newTag.ID,
		}

		err = u.photoTagsRepository.Add(photoTags)
		if err != nil {
			log.Printf("Failed to associate tag with photo %v", err)
			return &responsePhoto, err
		}
		// Hubungkan tag dengan foto
		updatedPhoto.Tags = append(updatedPhoto.Tags, *newTag)

	}

	return &responsePhoto, nil
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

func (u *photoUsecase) calculateTotalComments(photoID string, resultCh chan<- int64) {
	totalComments, err := u.commentRepository.CountCommentsByPhotoId(photoID)
	if err != nil {
		log.Printf("Error fetching total comments: %v", err)
		resultCh <- 0
		return
	}

	resultCh <- totalComments
}

func (u *photoUsecase) calculateTotalLikes(photoID string, resultCh chan<- int64) {
	totalLikes, err := u.userLikesPhotoRepository.CountUsersWhoLikedPhotoByPhotoId(photoID)
	if err != nil {
		log.Printf("Error fetching total likes: %v", err)
		resultCh <- 0
		return
	}

	resultCh <- totalLikes
}

func (u *photoUsecase) fetchUsername(userId uint, resultCh chan<- string) {
	user, err := u.userRepository.FindById(userId)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return
	}

	resultCh <- user.Username
}

func (u *photoUsecase) processPhotoTags(photoTags []domain.PhotoTags, responsePhoto *response.PhotoResponse) {
	for _, photoTag := range photoTags {
		tag, err := u.tagRepository.FindById(photoTag.TagId)
		if err != nil {
			log.Printf("Error fetching tag: %v", err)
			return
		}

		responsePhoto.PhotoTags = append(responsePhoto.PhotoTags, tag.Name)
	}
}

func NewPhotoUsecase(photo repository.PhotoRepository,
	comment repository.CommentRepository,
	tag repository.TagRepository,
	photoTags repository.PhotoTagsRepository,
	userLikesPhotoRepository repository.UserLikesPhotoRepository,
	userRepository repository.UserRepository) usecase.PhotoUsecase {
	return &photoUsecase{
		photoRepository:          photo,
		commentRepository:        comment,
		tagRepository:            tag,
		photoTagsRepository:      photoTags,
		userLikesPhotoRepository: userLikesPhotoRepository,
		userRepository:           userRepository,
	}
}
