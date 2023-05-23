package domain

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"not null" json:"username"`
	Email     string `gorm:"not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
