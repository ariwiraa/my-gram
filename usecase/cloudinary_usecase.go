package usecase

import "context"

type CloudinaryUsecase interface {
	// @Param File refer to file buffer
	// @Param pathDestination refer to target directory/bucket in cloud provider
	Upload(ctx context.Context, file interface{}, pathDestination string) (uri string, err error)
	Remove(ctx context.Context, urlString string, userId uint) (err error)
}
