package dtos

type UserRequest struct {
	Username string `validate:"required,min=3" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}

type UserLogin struct {
	Username string `validate:"required,min=3" json:"username"`
	Password string `validate:"required,min=8" json:"password"`
}
