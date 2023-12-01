package impl

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type followRepositoryImpl struct {
	db *gorm.DB
}

func (r followRepositoryImpl) Save(follow domain.Follow) error {
	err := r.db.Create(&follow).Error
	if err != nil {
		return err
	}
	return nil
}

func (r followRepositoryImpl) Delete(follow domain.Follow) error {
	err := r.db.
		Where("following_id = ? AND follower_id = ?", follow.FollowingId, follow.FollowerId).
		Delete(&follow).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r followRepositoryImpl) VerifyUserFollow(follow domain.Follow) (bool, error) {
	err := r.db.
		Where("following_id = ? AND follower_id = ?", follow.FollowingId, follow.FollowerId).
		First(&follow).
		Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func NewFollowRepositoryImpl(db *gorm.DB) repository.FollowRepository {
	return &followRepositoryImpl{db: db}
}
