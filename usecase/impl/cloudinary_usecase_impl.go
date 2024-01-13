package impl

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type cloudinaryImpl struct {
	cloud cloudinary.Cloudinary
}

func NewCloudinaryImpl(cloud cloudinary.Cloudinary) usecase.CloudinaryUsecase {
	return &cloudinaryImpl{cloud: cloud}
}

// Remove implements usecase.CloudinaryUsecase.
func (u *cloudinaryImpl) Remove(ctx context.Context, urlString string, userId uint) (err error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		log.Printf("[Remove, Parse] with error detail %v", err.Error())
		return
	}

	imagePath := path.Base(parsedURL.Path)
	filename := strings.TrimSuffix(imagePath, filepath.Ext(imagePath)) // Ambil nama file tanpa ekstensi

	pathDestination := fmt.Sprintf("%d-images", userId)

	res, err := u.cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: getPublicId(pathDestination, filename),
	})

	if err != nil {
		log.Printf("[Remove, Destroy] with error detail %v", err.Error())
		return err
	}

	if strings.Contains(res.Result, "not found") {
		return helpers.ErrPhotoNotFound
	}

	return err
}

// Upload implements usecase.CloudinaryUsecase.
func (u *cloudinaryImpl) Upload(ctx context.Context, file interface{}, pathDestination string) (uri string, err error) {
	filename := uuid.NewString()

	res, err := u.cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: getPublicId(pathDestination, filename),
		Eager:    "q_10",
	})

	if err != nil {
		log.Printf("[Upload, Upload] with error detail %v", err.Error())
		return "", err
	}

	// check if there are any eager in response
	if len(res.Eager) > 0 {
		// will return secure url with transformation
		return res.Eager[0].SecureURL, nil
	}

	// if no, will use secure url (without transformation)
	url := res.SecureURL

	return url, nil
}

func getPublicId(pathDestination, filename string) string {
	return "mygram-image/" + pathDestination + "/" + filename
}
