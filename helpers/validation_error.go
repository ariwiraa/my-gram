package helpers

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) string {
	var errorMessage string

	for _, fieldError := range err.(validator.ValidationErrors) {
		errorMessage = messageForTag(fieldError.Tag(), fieldError.Field())
	}

	return errorMessage
}

func messageForTag(tag string, field string) string {
	switch tag {
	case "required":
		return field + " Field is Required"
	case "email":
		return "Invalid email"
	}
	return ""
}
