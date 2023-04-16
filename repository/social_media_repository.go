package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	FindById(id uint) (domain.SocialMedia, error)
	FindAll() ([]domain.SocialMedia, error)
	Update(socialMedia domain.SocialMedia, id uint) (domain.SocialMedia, error)
	Delete(id uint)
}

type socialMediaRepository struct {
	db *gorm.DB
}

// Create implements SocialMediaRepository
func (r *socialMediaRepository) Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Debug().Create(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

// Delete implements SocialMediaRepository
func (r *socialMediaRepository) Delete(id uint) {
	var socialMedia domain.SocialMedia

	err := r.db.Debug().Where("id = ?", id).Delete(&socialMedia).Error
	if err != nil {
		log.Fatalln("error deleting data", err)
		return
	}
}

// FindAll implements SocialMediaRepository
func (r *socialMediaRepository) FindAll() ([]domain.SocialMedia, error) {
	var socialMedias []domain.SocialMedia

	err := r.db.Debug().Find(&socialMedias).Error
	if err != nil {
		log.Fatal("error getting all social medias data: ", err)
	}
	return socialMedias, nil
}

// FindById implements SocialMediaRepository
func (r *socialMediaRepository) FindById(id uint) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Debug().First(&socialMedia, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("social media not found")
		}
		log.Fatal("error getting data :", err)
	}

	return socialMedia, err
}

// Update implements SocialMediaRepository
func (r *socialMediaRepository) Update(socialMedia domain.SocialMedia, id uint) (domain.SocialMedia, error) {

	err := r.db.Debug().Model(&socialMedia).Where("id = ?", id).Updates(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}

	fmt.Println("repo", socialMedia)

	return socialMedia, nil
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db: db}
}
