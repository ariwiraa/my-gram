package impl

import (
	"context"
	"errors"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type authenticationRepositoryImpl struct {
	db *gorm.DB
}

// ExistsByRefreshToken implements repository.AuthenticationRepository.
func (r *authenticationRepositoryImpl) FindByRefreshToken(ctx context.Context, token string) (*domain.Authentication, error) {
	authentication := new(domain.Authentication)
	err := r.db.WithContext(ctx).First(&authentication, "refresh_token = ?", token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return authentication, helpers.ErrRefreshTokenNotFound
		}
		log.Printf("[FindByRefreshToken] with error detail %v", err.Error())
		return authentication, helpers.ErrRepository
	}

	return authentication, nil
}

// Add implements repository.AuthenticationRepository.
func (r *authenticationRepositoryImpl) Add(ctx context.Context, authentication domain.Authentication) error {
	err := r.db.WithContext(ctx).Create(&authentication).Error
	if err != nil {
		log.Printf("[Add] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil
}

// Delete implements repository.AuthenticationRepository.
func (r *authenticationRepositoryImpl) Delete(ctx context.Context, authentication domain.Authentication) error {
	err := r.db.WithContext(ctx).Where("refresh_token = ?", authentication.RefreshToken).Delete(&authentication).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ErrRefreshTokenNotFound
		}
		log.Printf("[Delete] with error detail %v", err.Error())
		return helpers.ErrRepository
	}

	return nil
}

func NewAuthenticationRepositoryImpl(db *gorm.DB) repository.AuthenticationRepository {
	return &authenticationRepositoryImpl{
		db: db,
	}
}
