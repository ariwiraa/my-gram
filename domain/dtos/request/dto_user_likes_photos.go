package request

type UserLikesPhotoRequest struct {
	PhotoId uint `validate:"required" json:"photo_id"`
}
