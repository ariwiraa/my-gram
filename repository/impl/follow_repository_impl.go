package impl

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type followRepositoryImpl struct {
	db *gorm.DB
}

func (r *followRepositoryImpl) CountFollowerByUserId(userId uint) (int64, error) {
	var totalFollower int64
	err := r.db.Model(&domain.Follow{}).Where("following_id = ?", userId).Count(&totalFollower).Error
	if err != nil {
		return totalFollower, err
	}

	return totalFollower, nil
}

func (r *followRepositoryImpl) CountFollowingByUserId(userId uint) (int64, error) {
	var totalFollower int64
	err := r.db.Model(&domain.Follow{}).Where("follower_id = ?", userId).Count(&totalFollower).Error
	if err != nil {
		return totalFollower, err
	}

	return totalFollower, nil
}

func (r *followRepositoryImpl) Save(follow domain.Follow) error {
	err := r.db.Create(&follow).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *followRepositoryImpl) Delete(follow domain.Follow) error {
	err := r.db.
		Where("following_id = ? AND follower_id = ?", follow.FollowingId, follow.FollowerId).
		Delete(&follow).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *followRepositoryImpl) VerifyUserFollow(follow domain.Follow) (bool, error) {
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
