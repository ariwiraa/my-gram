package usecase

import (
	"context"
	"mime/multipart"
)

type UploadFileUsecase interface {
	Upload(ctx context.Context, fileHeader *multipart.FileHeader, userId uint) (string, error)
}
