package impl

import (
	"errors"
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
	"log"
)

type followUsecaseImpl struct {
	followRepository repository.FollowRepository
	userRepository   repository.UserRepository
}

func (u *followUsecaseImpl) GetFollowingsByUsername(username string) ([]domain.User, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		log.Printf("Error fetching user by username: %v", err)
		return []domain.User{}, err
	}

	followings, err := u.followRepository.FindFollowingByUserId(user.ID)
	if err != nil {
		log.Printf("Error fetching followings by id: %v", err)
		return followings, err
	}

	log.Println("User followings fetched successfully")
	return followings, nil
}

func (u *followUsecaseImpl) GetFollowersByUsername(username string) ([]domain.User, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		log.Printf("Error fetching user by username: %v", err)
		return []domain.User{}, err
	}

	followers, err := u.followRepository.FindFollowersByUserId(user.ID)
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

func (u *followUsecaseImpl) FollowUser(followRequest request.FollowRequest) (string, error) {
	err := u.userRepository.IsUserExists(followRequest.UserIdFollowing)
	if err != nil {
		log.Printf("Error checking user by id: %v", err)
		return "", errors.New("user doesn't exists")
	}

	follow := domain.Follow{
		FollowerId:  followRequest.UserIdFollower,
		FollowingId: followRequest.UserIdFollowing,
	}

	followed, _ := u.followRepository.VerifyUserFollow(follow)

	var message string
	if !followed {
		err := u.followRepository.Save(follow)
		if err != nil {
			return "", errors.New("failed follow user")
		}
		log.Printf("id %d succesfully follow id %d", followRequest.UserIdFollowing, followRequest.UserIdFollower)
		message = "successfully followed"
	} else {
		err := u.followRepository.Delete(follow)
		if err != nil {
			return "", errors.New("failed unfollow user")
		}
		log.Printf("id %d succesfully unfollow id %d", followRequest.UserIdFollowing, followRequest.UserIdFollower)
		message = "successfully unfollowed"
	}

	return message, nil
}
