package domain

import (
	"time"
)

// Photo represents the model for an Photo
type Photo struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Caption      string     `json:"caption"`
	PhotoUrl     string     `gorm:"not null" json:"photo_url"`
	UserId       uint       `json:"user_id"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	User         User       `gorm:"foreignKey:UserId" json:"-"`
	TotalComment int64      `gorm:"-" json:"total_comment"`
	Comments     []Comment  `gorm:"foreignKey:PhotoId" json:"comments,omitempty"`
	LikedBy      []User     `gorm:"many2many:user_likes_photos" json:"liked_by,omitempty"`
	Tags         []Tag      `gorm:"many2many:photo_tags" json:"tags,omitempty"`
}
