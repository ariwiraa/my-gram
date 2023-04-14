package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an Comment
type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Message   string `gorm:"not null" valid:"required~your message is required" json:"message"`
	PhotoId   uint   `json:"photo_id"`
	UserId    uint   `json:"user_id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type CommentRequest struct {
	Message string `json:"message"`
	PhotoId uint   `json:"photo_id"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
