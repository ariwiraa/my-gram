package impl

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/usecase"
)

type authenticationUsecaseImpl struct {
	repo repository.AuthenticationRepository
}

// Add implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) Add(token string) error {
	authentication := new(domain.Authentication)

	authentication.RefreshToken = token

	err := u.repo.Add(*authentication)
	if err != nil {
		return err
	}

	return nil

}

// Delete implements usecase.AuthenticationUsecase.
func (u *authenticationUsecaseImpl) Delete(token string) error {
	authentication, err := u.repo.FindByRefreshToken(token)
	if err != nil {
		return err
	}

	err = u.repo.Delete(*authentication)
	if err != nil {
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) ExistsByRefreshToken(token string) error {
	_, err := u.repo.FindByRefreshToken(token)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthenticationUsecaseImpl(repo repository.AuthenticationRepository) usecase.AuthenticationUsecase {
	return &authenticationUsecaseImpl{
		repo: repo,
	}
}
