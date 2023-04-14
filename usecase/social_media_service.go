package usecase

import (
	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/repository"
)

type SocialMediaUsecase interface {
	Create(payload domain.SocialMediaRequest, userId uint) (domain.SocialMedia, error)
	GetById(id uint) (domain.SocialMedia, error)
	GetAll() ([]domain.SocialMedia, error)
	Update(payload domain.SocialMediaRequest, id uint, userId uint) (domain.SocialMedia, error)
	Delete(socialMedia domain.SocialMedia)
}

type socialMediaUsecase struct {
	socialMediaRepository repository.SocialMediaRepository
}

// Create implements SocialMediaUsecase
func (u *socialMediaUsecase) Create(payload domain.SocialMediaRequest, userId uint) (domain.SocialMedia, error) {
	socialMedia := domain.SocialMedia{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:         userId,
	}

	newSocialMedia, err := u.socialMediaRepository.Create(socialMedia)
	if err != nil {
		return newSocialMedia, err
	}

	return newSocialMedia, nil
}

// Delete implements SocialMediaUsecase
func (u *socialMediaUsecase) Delete(socialMedia domain.SocialMedia) {
	u.socialMediaRepository.Delete(socialMedia.ID)
}

// GetAll implements SocialMediaUsecase
func (u *socialMediaUsecase) GetAll() ([]domain.SocialMedia, error) {
	socialMedias, err := u.socialMediaRepository.FindAll()
	if err != nil {
		return socialMedias, err
	}

	return socialMedias, nil
}

// GetById implements SocialMediaUsecase
func (u *socialMediaUsecase) GetById(id uint) (domain.SocialMedia, error) {
	socialMedia, err := u.socialMediaRepository.FindById(id)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

// Update implements SocialMediaUsecase
func (u *socialMediaUsecase) Update(payload domain.SocialMediaRequest, id uint, userId uint) (domain.SocialMedia, error) {
	socialMedia, err := u.socialMediaRepository.FindById(id)
	if err != nil {
		panic(err)
	}

	// var socialMedia domain.SocialMedia

	socialMedia.Name = payload.Name
	socialMedia.SocialMediaUrl = payload.SocialMediaUrl
	socialMedia.UserId = userId

	updatedSocialMedia, err := u.socialMediaRepository.Update(socialMedia, id)
	if err != nil {
		return updatedSocialMedia, err
	}

	return updatedSocialMedia, nil
}

func NewSocialMediaUsecase(socialMedia repository.SocialMediaRepository) SocialMediaUsecase {
	return &socialMediaUsecase{socialMediaRepository: socialMedia}
}
