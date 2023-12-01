package domain

import "time"

type Follow struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	FollowerId   uint       `json:"follower_id"`
	FollowingId  uint       `json:"following_id"`
	DateFollowed *time.Time `json:"date_followed"`
	Follower     User       `gorm:"foreignKey:FollowerId" json:"follower"`
	Following    User       `gorm:"foreignKey:FollowingId" json:"following"`
}
