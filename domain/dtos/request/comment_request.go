package request

type CommentRequest struct {
	Message string `validate:"required" json:"message"`
	PhotoId string
	UserId  uint
}
