package domain

import (
	"time"
)

// Comment represents the model for an Comment
type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Message   string `gorm:"not null" json:"message"`
	PhotoId   string `json:"photo_id"`
	UserId    uint   `json:"user_id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	User      User  `gorm:"foreignKey:UserId" json:"-"`
	Photo     Photo `gorm:"foreignKey:PhotoId" json:"-"`
}
