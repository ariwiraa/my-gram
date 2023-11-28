package domain

import (
	"time"
)

type User struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Username    string     `gorm:"not null" json:"username"`
	Email       string     `gorm:"not null" json:"email"`
	Password    string     `gorm:"not null" json:"-"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
	Photos      []Photo    `gorm:"foreignKey:UserId;" json:"-"`
	Comments    []Comment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	LikedPhotos []Photo    `gorm:"many2many:user_likes_photos" json:"-"`
}
