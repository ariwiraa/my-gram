package impl

import (
	"context"
	"log"
	"time"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/domain/dtos/response"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/google/uuid"
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
func (u *photoUsecase) Create(ctx context.Context, payload request.PhotoRequest, userId uint, fileLocation string) (*response.PhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	photo := domain.Photo{
		ID:       uuid.NewString(),
		Caption:  payload.Caption,
		PhotoUrl: fileLocation,
		UserId:   userId,
	}

	usernameCh := make(chan string)
	go u.fetchUsername(ctx, userId, usernameCh)

	newPhoto, err := u.photoRepository.Create(ctx, photo)
	if err != nil {
		log.Printf("[Create, Create] with error detail %v", err.Error())
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

		newTag, err := u.tagRepository.AddTagIfNotExists(ctx, tag.Name)
		if err != nil {
			log.Printf("[Create, AddTagIfNotExists] with error detail %v", err.Error())
			return &response.PhotoResponse{}, err
		}

		photoTags := domain.PhotoTags{
			PhotoId: newPhoto.ID,
			TagId:   newTag.ID,
		}

		err = u.photoTagsRepository.Add(ctx, photoTags)
		if err != nil {
			log.Printf("[Create, Add] with error detail %v", err.Error())
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
func (u *photoUsecase) Delete(ctx context.Context, id string) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	photo, err := u.photoRepository.FindById(ctx, id)
	if err != nil {
		log.Printf("[Delete, FindById] with error detail %v", err.Error())
		return
	}

	u.photoRepository.Delete(ctx, photo)
	err = u.photoTagsRepository.Delete(ctx, photo.ID)
	if err != nil {
		log.Printf("[Delete, Delete] with error detail %v", err.Error())
		return
	}
}

// GetAll implements PhotoUsecase
func (u *photoUsecase) GetAll(ctx context.Context) ([]domain.Photo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	photos, err := u.photoRepository.FindAll(ctx)
	if err != nil {
		return photos, err
	}

	for _, photo := range photos {
		totalComments, _ := u.commentRepository.CountCommentsByPhotoId(ctx, photo.ID)
		photo.TotalComment = totalComments

		photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(ctx, photo.ID)
		// Jika tidak ada tag, maka return photo
		if err != nil {
			return photos, nil
		}

		for _, photoTag := range photoTags {
			tag, err := u.tagRepository.FindById(ctx, photoTag.TagId)
			if err != nil {
				return photos, err
			}

			photo.Tags = append(photo.Tags, *tag)
		}
	}

	return photos, nil
}

func (u *photoUsecase) GetById(ctx context.Context, id string) (*response.PhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	photo, err := u.photoRepository.FindById(ctx, id)
	if err != nil {
		log.Printf("[GetById, FindById] with error detail %v", err.Error())
		return &response.PhotoResponse{}, err
	}

	// Menggunakan goroutine untuk mengambil total comments dan total likes secara bersamaan
	totalCommentsCh := make(chan int64)
	totalLikesCh := make(chan int64)

	go u.calculateTotalComments(ctx, photo.ID, totalCommentsCh)
	go u.calculateTotalLikes(ctx, photo.ID, totalLikesCh)

	totalComments := <-totalCommentsCh
	totalLikes := <-totalLikesCh

	photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(ctx, id)
	if err != nil {
		log.Printf("[GetById, FindPhotoTagsByPhotoId] with error detail %v", err.Error())
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

	u.processPhotoTags(ctx, photoTags, &responsePhoto)

	return &responsePhoto, nil
}

// Update implements PhotoUsecase
func (u *photoUsecase) Update(ctx context.Context, payload request.UpdatePhotoRequest, id string, userId uint) (*response.PhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	photo, err := u.photoRepository.FindByIdAndByUserId(ctx, id, userId)
	if err != nil {
		log.Printf("[Update, FIndByIdAndByUserId] with error detail %v", err.Error())
		return &response.PhotoResponse{}, err
	}

	photo.Caption = payload.Caption

	err = u.photoTagsRepository.Delete(ctx, id)
	if err != nil {
		log.Printf("[Update, Delete] with error detail %v", err.Error())
		return &response.PhotoResponse{}, err
	}

	updatedPhoto, err := u.photoRepository.Update(ctx, *photo, id)
	if err != nil {
		log.Printf("[Update, Update] with error detail %v", err.Error())
		return &response.PhotoResponse{}, err
	}

	totalCommentsCh := make(chan int64)
	totalLikesCh := make(chan int64)
	usernameCh := make(chan string)

	go u.fetchUsername(ctx, userId, usernameCh)
	go u.calculateTotalComments(ctx, photo.ID, totalCommentsCh)
	go u.calculateTotalLikes(ctx, photo.ID, totalLikesCh)

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

		newTag, err := u.tagRepository.AddTagIfNotExists(ctx, tag.Name)
		if err != nil {
			log.Printf("[Update, AddTagIfNotExists] with error detail %v", err.Error())
			return &responsePhoto, err
		}

		photoTags := domain.PhotoTags{
			PhotoId: updatedPhoto.ID,
			TagId:   newTag.ID,
		}

		err = u.photoTagsRepository.Add(ctx, photoTags)
		if err != nil {
			log.Printf("[Update, Add] with error detail %v", err.Error())
			return &responsePhoto, err
		}
		// Hubungkan tag dengan foto
		updatedPhoto.Tags = append(updatedPhoto.Tags, *newTag)

	}

	return &responsePhoto, nil
}

func (u *photoUsecase) GetAllPhotosByUserId(ctx context.Context, userId uint) ([]domain.Photo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	photos, err := u.photoRepository.FindByUserId(ctx, userId)
	if err != nil {
		return photos, err
	}

	for _, photo := range photos {
		totalComments, _ := u.commentRepository.CountCommentsByPhotoId(ctx, photo.ID)
		photo.TotalComment = totalComments

		photoTags, err := u.photoTagsRepository.FindPhotoTagsByPhotoId(ctx, photo.ID)
		// Jika tidak ada tag, maka return photo
		if err != nil {
			log.Printf("[GetAllPhotosByUserId, FindPhotoTagsByPhotoId] with error detail %v", err.Error())
			return photos, nil
		}

		for _, photoTag := range photoTags {
			tag, err := u.tagRepository.FindById(ctx, photoTag.TagId)
			if err != nil {
				log.Printf("[GetAllPhotosByUserId, FindById] with error detail %v", err.Error())
				return photos, err
			}

			photo.Tags = append(photo.Tags, *tag)

		}
	}

	return photos, nil
}

func (u *photoUsecase) calculateTotalComments(ctx context.Context, photoID string, resultCh chan<- int64) {
	totalComments, err := u.commentRepository.CountCommentsByPhotoId(ctx, photoID)
	if err != nil {
		log.Printf("[calculateTotalComments, CountCommentsByPhotoId] with error detail %v", err.Error())
		resultCh <- 0
		return
	}

	resultCh <- totalComments
}

func (u *photoUsecase) calculateTotalLikes(ctx context.Context, photoID string, resultCh chan<- int64) {
	totalLikes, err := u.userLikesPhotoRepository.CountUsersWhoLikedPhotoByPhotoId(ctx, photoID)
	if err != nil {
		log.Printf("[calculateTotalLikes, CountUsersWhoLikedPhotoByPhotoId] with error detail %v", err.Error())
		resultCh <- 0
		return
	}

	resultCh <- totalLikes
}

func (u *photoUsecase) fetchUsername(ctx context.Context, userId uint, resultCh chan<- string) {
	user, err := u.userRepository.FindById(ctx, userId)
	if err != nil {
		log.Printf("[fetchUsername, FindById] with error detail %v", err.Error())
		return
	}

	resultCh <- user.Username
}

func (u *photoUsecase) processPhotoTags(ctx context.Context, photoTags []domain.PhotoTags, responsePhoto *response.PhotoResponse) {
	for _, photoTag := range photoTags {
		tag, err := u.tagRepository.FindById(ctx, photoTag.TagId)
		if err != nil {
			log.Printf("[processPhotoTags, FindById] with error detail %v", err.Error())
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
