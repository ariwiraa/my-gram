package impl

import (
	"context"
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
	"log"
	"time"
)

type followUsecaseImpl struct {
	followRepository repository.FollowRepository
	userRepository   repository.UserRepository
}

func (u *followUsecaseImpl) GetFollowingsByUsername(ctx context.Context, username string) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		log.Printf("Error fetching user by username: %v", err)
		return []domain.User{}, err
	}

	followings, err := u.followRepository.FindFollowingByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("Error fetching followings by id: %v", err)
		return followings, err
	}

	log.Println("User followings fetched successfully")
	return followings, nil
}

func (u *followUsecaseImpl) GetFollowersByUsername(ctx context.Context, username string) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		log.Printf("Error fetching user by username: %v", err)
		return []domain.User{}, err
	}

	followers, err := u.followRepository.FindFollowersByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("Error fetching followers by id: %v", err)
		return followers, err
	}

	log.Println("User followers fetched successfully")
	return followers, nil
}

func NewFollowUsecaseImpl(followRepository repository.FollowRepository, userRepository repository.UserRepository) usecase.FollowUsecase {
	return &followUsecaseImpl{followRepository: followRepository, userRepository: userRepository}
}

func (u *followUsecaseImpl) FollowUser(ctx context.Context, followRequest request.FollowRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := u.userRepository.IsUserExists(ctx, followRequest.UserIdFollowing)
	if err != nil {
		log.Printf("Error checking user by id: %v", err)
		return "", errors.New("user doesn't exists")
	}

	follow := domain.Follow{
		FollowerId:  followRequest.UserIdFollower,
		FollowingId: followRequest.UserIdFollowing,
	}

	followed, _ := u.followRepository.VerifyUserFollow(ctx, follow)

	var message string
	if !followed {
		err := u.followRepository.Save(ctx, follow)
		if err != nil {
			return "", errors.New("failed follow user")
		}
		log.Printf("id %d succesfully follow id %d", followRequest.UserIdFollowing, followRequest.UserIdFollower)
		message = "successfully followed"
	} else {
		err := u.followRepository.Delete(ctx, follow)
		if err != nil {
			return "", errors.New("failed unfollow user")
		}
		log.Printf("id %d succesfully unfollow id %d", followRequest.UserIdFollowing, followRequest.UserIdFollower)
		message = "successfully unfollowed"
	}

	return message, nil
}
