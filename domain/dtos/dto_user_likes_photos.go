package dtos

type UserLikesPhotoRequest struct {
	PhotoId uint `validate:"required" json:"photo_id"`
}
