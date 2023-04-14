package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for an Photo
type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"not null" valid:"required~your title is required" json:"title"`
	Caption   string `gorm:"not null" json:"caption"`
	PhotoUrl  string `gorm:"not null" valid:"required~your photo is required" json:"photo_url"`
	UserId    uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	User      *User `json:"user"`
}

type PhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
