package helpers

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) error {
	var errorMessage error

	for _, fieldError := range err.(validator.ValidationErrors) {
		errorMessage = messageForTag(fieldError.Tag(), fieldError.Field(), fieldError.Param())
	}

	return errorMessage
}

func messageForTag(tag string, field string, param string) error {
	switch tag {
	case "email":
		return ErrorEmailInvalid
	case "required":
		return ErrorFieldRequired(field)
	case "min":
		return ErrorFieldMinimum(field)
	}
	return ErrBadRequest
}

func ErrorFieldRequired(field string) error {
	switch field {
	case "Email":
		return ErrEmailRequired
	case "Password":
		return ErrPasswordRequired
	case "Username":
		return ErrUsernameRequired
	case "Message":
		return ErrCommentMessageRequired
	}

	return ErrBadRequest
}

func ErrorFieldMinimum(field string) error {
	switch field {
	case "Password":
		return ErrPasswordInvalidLength
	case "Username":
		return ErrUsernameInvalidLength
	}

	return ErrBadRequest
}
