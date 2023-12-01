package usecase

import "github.com/ariwiraa/my-gram/domain/dtos/request"

type FollowUsecase interface {
	FollowUser(followRequest request.FollowRequest) (string, error)
}
