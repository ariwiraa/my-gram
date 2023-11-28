package domain

import "time"

type UserLikesPhoto struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserId    uint   `gorm:"not null" json:"user_id"`
	PhotoId   string `gorm:"not null" json:"photo_id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
