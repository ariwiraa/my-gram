package repository

import "github.com/ariwiraa/my-gram/domain"

type AuthenticationRepository interface {
	Add(authentication domain.Authentication) error
	FindByRefreshToken(token string) (*domain.Authentication, error)
	Delete(authentication domain.Authentication) error
}
