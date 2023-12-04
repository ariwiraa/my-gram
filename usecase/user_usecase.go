package usecase

import "github.com/ariwiraa/my-gram/domain/dtos/response"

type UserUsecase interface {
	GetUserProfileByUsername(username string) (*response.UserProfileResponse, error)
}
