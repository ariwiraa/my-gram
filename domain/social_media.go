package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" json:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url"`
	UserId         uint
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type SocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
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
