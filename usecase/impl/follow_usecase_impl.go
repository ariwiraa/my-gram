package impl

import (
	"context"
	"log"
	"time"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type followUsecaseImpl struct {
	followRepository repository.FollowRepository
	userRepository   repository.UserRepository
}

func NewFollowUsecaseImpl(followRepository repository.FollowRepository, userRepository repository.UserRepository) usecase.FollowUsecase {
	return &followUsecaseImpl{followRepository: followRepository, userRepository: userRepository}
}

func (u *followUsecaseImpl) GetFollowingsByUsername(ctx context.Context, username string) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		log.Printf("[GetFollowingByUsername, FindByUsername] with error detail %v", err.Error())
		return []domain.User{}, err
	}

	followings, err := u.followRepository.FindFollowingByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("[GetFollowingByUsername, FindFollowingByUserId] with error detail %v", err.Error())
		return followings, err
	}

	return followings, nil
}

func (u *followUsecaseImpl) GetFollowersByUsername(ctx context.Context, username string) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		log.Printf("[GetFollowersByUsername, FindByUsername] with error detail %v", err.Error())
		return []domain.User{}, err
	}

	followers, err := u.followRepository.FindFollowersByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("[GetFollowingByUsername, FindFollowersByUserId] with error detail %v", err.Error())
		return followers, err
	}

	return followers, nil
}

func (u *followUsecaseImpl) FollowUser(ctx context.Context, followRequest request.FollowRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := u.userRepository.IsUserExists(ctx, followRequest.UserIdFollowing)
	if err != nil {
		log.Printf("[FollowUser, IsUserExists] with error detail %v", err.Error())
		return "", err
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
			log.Printf("[FollowUser, Save] with error detail %v", err.Error())
			return "", err
		}
		log.Printf("id %d succesfully follow id %d", followRequest.UserIdFollowing, followRequest.UserIdFollower)
		message = "successfully followed"
	} else {
		err := u.followRepository.Delete(ctx, follow)
		if err != nil {
			log.Printf("[FollowUser, Delete] with error detail %v", err.Error())
			return "", err
		}
		log.Printf("id %d succesfully unfollow id %d", followRequest.UserIdFollowing, followRequest.UserIdFollower)
		message = "successfully unfollowed"
	}

	return message, nil
}
