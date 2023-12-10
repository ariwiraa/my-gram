package impl

import (
	"context"
	"fmt"
	"github.com/ariwiraa/my-gram/domain/dtos/response"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
	"log"
	"time"
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
		log.Printf("Error fetching user by username: %v", err)
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	// TODO: Terapkan goroutine dan channel
	follower, err := u.followRepository.CountFollowerByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("Error counting followers: %v", err)
		return nil, fmt.Errorf("failed to count followers: %w", err)
	}

	following, err := u.followRepository.CountFollowingByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("Error counting following: %v", err)
		return nil, fmt.Errorf("failed to count following: %w", err)
	}

	totalPosts, err := u.photoRepository.CountPhotoByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("Error counting total posts: %v", err)
		return nil, fmt.Errorf("failed to count total posts: %w", err)
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
