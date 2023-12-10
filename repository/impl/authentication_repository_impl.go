package impl

import (
	"context"
	"github.com/ariwiraa/my-gram/domain"
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
		return authentication, err
	}
	return authentication, nil
}

// Add implements repository.AuthenticationRepository.
func (r *authenticationRepositoryImpl) Add(ctx context.Context, authentication domain.Authentication) error {
	err := r.db.WithContext(ctx).Create(&authentication).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete implements repository.AuthenticationRepository.
func (r *authenticationRepositoryImpl) Delete(ctx context.Context, authentication domain.Authentication) error {
	err := r.db.WithContext(ctx).Where("refresh_token = ?", authentication.RefreshToken).Delete(&authentication).Error
	if err != nil {
		return err
	}

	return nil
}

func NewAuthenticationRepositoryImpl(db *gorm.DB) repository.AuthenticationRepository {
	return &authenticationRepositoryImpl{
		db: db,
	}
}
