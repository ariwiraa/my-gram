package dtos

type CommentRequest struct {
	Message string `validate:"required" json:"message"`
	PhotoId uint   `validate:"required" json:"photo_id"`
}
