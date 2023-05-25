package domain

import (
	"time"
)

// Photo represents the model for an Photo
type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Caption   string `gorm:"not null" json:"caption"`
	PhotoUrl  string `gorm:"not null" json:"photo_url"`
	UserId    uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	User      *User `json:"user"`
}
