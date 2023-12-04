package impl

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
	"gorm.io/gorm"
)

type followRepositoryImpl struct {
	db *gorm.DB
}

func (r *followRepositoryImpl) FindFollowingByUserId(userId uint) ([]domain.User, error) {
	var followings []domain.User

	err := r.db.Debug().Table("follows").
		Select("users.*").
		Joins("INNER JOIN users ON follows.following_id = users.id").
		Where("follows.follower_id = ?", userId).
		Find(&followings).
		Error

	if err != nil {
		return followings, err
	}

	return followings, nil
}

func (r *followRepositoryImpl) FindFollowersByUserId(userId uint) ([]domain.User, error) {
	var followers []domain.User

	err := r.db.Debug().Table("follows").
		Select("users.*").
		Joins("INNER JOIN users ON follows.follower_id = users.id").
		Where("follows.following_id = ?", userId).
		Find(&followers).
		Error

	if err != nil {
		return followers, err
	}

	return followers, nil
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
