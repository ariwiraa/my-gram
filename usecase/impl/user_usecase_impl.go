package impl

import (
	"context"
	"log"
	"time"

	"github.com/ariwiraa/my-gram/domain/dtos/response"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type userUsecaseImpl struct {
	userRepository   repository.UserRepository
	photoRepository  repository.PhotoRepository
	followRepository repository.FollowRepository
}

func (u *userUsecaseImpl) GetUserProfileByUsername(ctx context.Context, username string) (*response.UserProfileResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Printf("Fetching user profile for username: %s", username)
	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		log.Printf("[GetUserProfileByUsername, FindByUsername] with error detail %v", err.Error())
		return nil, err
	}

	// TODO: Terapkan goroutine dan channel
	follower, err := u.followRepository.CountFollowerByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("[GetUserProfileByUsername, CountFollowerByUserId] with error detail %v", err.Error())
		return nil, err
	}

	following, err := u.followRepository.CountFollowingByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("[GetUserProfileByUsername, CountFollowingByUserId] with error detail %v", err.Error())
		return nil, err
	}

	totalPosts, err := u.photoRepository.CountPhotoByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("[GetUserProfileByUsername, CountPhotoByUserId] with error detail %v", err.Error())
		return nil, err
	}

	log.Println("User profile fetched successfully")
	return &response.UserProfileResponse{
		Username:   user.Username,
		Following:  following,
		Follower:   follower,
		PostsCount: totalPosts,
		Posts:      user.Photos,
	}, nil
}

func NewUserUsecaseImpl(userRepository repository.UserRepository, photoRepository repository.PhotoRepository, followRepository repository.FollowRepository) usecase.UserUsecase {
	return &userUsecaseImpl{
		userRepository:   userRepository,
		photoRepository:  photoRepository,
		followRepository: followRepository,
	}
}
