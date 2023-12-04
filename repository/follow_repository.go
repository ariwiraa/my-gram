package repository

import (
	"github.com/ariwiraa/my-gram/domain"
)

type FollowRepository interface {
	Save(follow domain.Follow) error
	Delete(follow domain.Follow) error
	VerifyUserFollow(follow domain.Follow) (bool, error)
	CountFollowerByUserId(userId uint) (int64, error)
	CountFollowingByUserId(userId uint) (int64, error)
	FindFollowersByUserId(userId uint) ([]domain.User, error)
	FindFollowingByUserId(userId uint) ([]domain.User, error)
}
