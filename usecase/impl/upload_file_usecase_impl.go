package impl

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
)

type uploadFileUsecaseImpl struct {
	cloudinary usecase.CloudinaryUsecase
}

func NewUploadFileImpl(cloudinary usecase.CloudinaryUsecase) usecase.UploadFileUsecase {
	return &uploadFileUsecaseImpl{
		cloudinary: cloudinary,
	}
}

// Upload implements usecase.UploadFile.
func (u *uploadFileUsecaseImpl) Upload(ctx context.Context, fileHeader *multipart.FileHeader, userId uint) (string, error) {

	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("[UploadFile, Upload, Open] error with detail %v", err.Error())
		return "", err
	}
	defer file.Close()

	fileSupported, err := helpers.IsImageFile(file)
	if err != nil {
		log.Printf("[UploadFile, Upload, IsImageFile] error with detail %v", err.Error())
		return "", err
	}

	if !fileSupported {
		return "", helpers.ErrFileNotSupported
	}

	err = helpers.IsFileSizeValid(fileHeader)
	if err != nil {
		return "", helpers.ErrorFileSizeNotValid
	}

	buffer := bytes.NewBuffer(nil)

	_, err = io.Copy(buffer, file)
	if err != nil {
		log.Printf("[UploadFileUsecase, Upload, Copy] error with detail %v", err.Error())
		return "", err
	}

	pathDestination := fmt.Sprintf("%d-images", userId)

	url, err := u.cloudinary.Upload(ctx, buffer, pathDestination)
	if err != nil {
		log.Printf("[UploadFile, Upload, Upload] error with detail %v", err.Error())
		return url, err
	}

	return url, err
}
