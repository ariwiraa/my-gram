package dtos

type CommentRequest struct {
	Message string `validate:"required" json:"message"`
	PhotoId string `validate:"required" json:"photo_id"`
}
