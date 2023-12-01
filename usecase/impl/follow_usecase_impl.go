package impl

import (
	"errors"
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

func (u *followUsecaseImpl) FollowUser(followRequest request.FollowRequest) (string, error) {
	err := u.userRepository.IsUserExists(followRequest.UserIdFollowing)
	if err != nil {
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
		message = "successfully followed"
	} else {
		err := u.followRepository.Delete(follow)
		if err != nil {
			return "", errors.New("failed unfollow user")
		}
		message = "successfully unfollowed"
	}

	return message, nil
}
