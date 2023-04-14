package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" valid:"required~your name is required" json:"name"`
	SocialMediaUrl string `gorm:"not null" valid:"required~your socialmedia is required" json:"social_media_url"`
	UserId         uint
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type SocialMediaRequest struct {
	Name           string `gorm:"not null" valid:"required~your name is required" json:"name"`
	SocialMediaUrl string `gorm:"not null" valid:"required~your socialmedia is required" json:"social_media_url"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
