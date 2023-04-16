package domain

import (
	"time"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Username    string `gorm:"not null;uniqueIndex" valid:"required~Your username is required " json:"username"`
	Email       string `gorm:"not null;uniqueIndex" valid:"required~Your email is required,email~Invalid email format" json:"email"`
	Password    string `gorm:"not null" valid:"required,minstringlength(6)~Password has to have a minimum length of 6 characters" json:"-"`
	Age         uint   `gorm:"not null" valid:"required,range(8|100)~Age must be over 8" json:"age"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	SocialMedia SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Photos      []Photo     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments    []Comment   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint   `json:"age"`
}

type UserLogin struct {
	Email    string `valid:"required~Your email is required,email~Invalid email format" json:"email"`
	Password string `valid:"required,minstringlength(6)~Password has to have a minimum length of 6 characters" json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
