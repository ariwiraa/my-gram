package request

import "mime/multipart"

type PhotoRequest struct {
	Caption  string                `form:"caption"`
	PhotoUrl *multipart.FileHeader `form:"photo_url" binding:"required"`
	Tags     []string              `form:"tags" `
}

type UpdatePhotoRequest struct {
	Caption string   `form:"caption"`
	Tags    []string `form:"tags"`
}
