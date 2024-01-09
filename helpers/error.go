package helpers

import (
	"errors"
	"net/http"
)

var (
	// base error
	ErrBadRequest          = errors.New("bad request")
	ErrNotFound            = errors.New("not found")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("general error")
	ErrForbiddenAccess     = errors.New("forbidden error")

	// field error bad request
	ErrEmailInvalid          = errors.New("email invalid")
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailAlreadyUserd     = errors.New("email is already used")
	ErrEmailNotFound         = errors.New("email not found")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password minimum length is 8")
	ErrPasswordNotMatch      = errors.New("password not match")
	ErrUsernameInvalidLength = errors.New("username minimum length is 3")
	ErrUsernameRequired      = errors.New("username is required")
	ErrLinkExpired           = errors.New("your link is expired")
	ErrUsernameAlreadyUsed   = errors.New("username already used")
	ErrEmailNotVerified      = errors.New("email not verified. Please verif your email first")
	ErrRefreshTokenNotFound  = errors.New("refresh token not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrPhotoNotFound         = errors.New("photo not found")
	ErrCommentNotFound       = errors.New("comment not found")
	ErrTagNotFound           = errors.New("tag not found")
	ErrFileNotSupported      = errors.New("file not supported")

	ErrCommentMessageRequired = errors.New("message is required")

	// general
	ErrFailedSendEmail = errors.New("failed send email")
	ErrRepository      = errors.New("error repository")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral         = NewError(ErrInternalServerError.Error(), "99999", http.StatusInternalServerError)
	ErrorBadRequest      = NewError(ErrBadRequest.Error(), "40000", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorUnauthorized    = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), "40100", http.StatusForbidden)
)

var (
	// bad request
	ErrorEmailInvalid           = NewError(ErrEmailInvalid.Error(), "40001", http.StatusBadRequest)
	ErrorEmailRequired          = NewError(ErrEmailRequired.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired       = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength  = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorUsernameRequired       = NewError(ErrUsernameRequired.Error(), "40005", http.StatusBadRequest)
	ErrorUsernameInvalidLength  = NewError(ErrUsernameInvalidLength.Error(), "40006", http.StatusBadRequest)
	ErrorCommentMessageRequired = NewError(ErrCommentMessageRequired.Error(), "40007", http.StatusBadRequest)
	ErrorLinkExpired            = NewError(ErrLinkExpired.Error(), "40008", http.StatusBadRequest)
	ErrorFileNotSupported       = NewError(ErrFileNotSupported.Error(), "40009", http.StatusBadRequest)

	// conflict
	ErrorEmailAlreadyUsed    = NewError(ErrEmailAlreadyUserd.Error(), "40901", http.StatusConflict)
	ErrorUsernameAlreadyUsed = NewError(ErrUsernameAlreadyUsed.Error(), "40902", http.StatusConflict)

	// not found
	ErrorEmailNotFound        = NewError(ErrEmailNotFound.Error(), "40401", http.StatusNotFound)
	ErrorRefreshTokenNotFound = NewError(ErrRefreshTokenNotFound.Error(), "40402", http.StatusNotFound)
	ErrorUserNotFound         = NewError(ErrUserNotFound.Error(), "40403", http.StatusNotFound)
	ErrorPhotoNotFound        = NewError(ErrPhotoNotFound.Error(), "40404", http.StatusNotFound)
	ErrorCommentNotFound      = NewError(ErrCommentNotFound.Error(), "40405", http.StatusNotFound)
	ErrorTagNotFound          = NewError(ErrTagNotFound.Error(), "40406", http.StatusNotFound)

	// unauthorized
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
	ErrorEmailNotVerified = NewError(ErrEmailNotVerified.Error(), "40102", http.StatusUnauthorized)

	// internal server error
	ErrorRepository      = NewError(ErrRepository.Error(), "50001", http.StatusInternalServerError)
	ErrorFailedSendEmail = NewError(ErrFailedSendEmail.Error(), "50002", http.StatusInternalServerError)
)

var (
	ErrorMapping = map[string]Error{
		ErrEmailInvalid.Error():           ErrorEmailInvalid,
		ErrEmailRequired.Error():          ErrorEmailRequired,
		ErrPasswordRequired.Error():       ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error():  ErrorPasswordInvalidLength,
		ErrPasswordNotMatch.Error():       ErrorPasswordNotMatch,
		ErrUsernameRequired.Error():       ErrorUsernameRequired,
		ErrUsernameInvalidLength.Error():  ErrorUsernameInvalidLength,
		ErrCommentMessageRequired.Error(): ErrorCommentMessageRequired,
		ErrEmailNotFound.Error():          ErrorEmailNotFound,
		ErrFailedSendEmail.Error():        ErrorFailedSendEmail,
		ErrRepository.Error():             ErrorRepository,
		ErrLinkExpired.Error():            ErrorLinkExpired,
		ErrEmailAlreadyUserd.Error():      ErrorEmailAlreadyUsed,
		ErrUsernameAlreadyUsed.Error():    ErrorUsernameAlreadyUsed,
		ErrEmailNotVerified.Error():       ErrorEmailNotVerified,
		ErrRefreshTokenNotFound.Error():   ErrorRefreshTokenNotFound,
		ErrUserNotFound.Error():           ErrorUserNotFound,
		ErrPhotoNotFound.Error():          ErrorPhotoNotFound,
		ErrCommentNotFound.Error():        ErrorCommentNotFound,
		ErrTagNotFound.Error():            ErrorTagNotFound,
		ErrFileNotSupported.Error():       ErrorFileNotSupported,
	}
)
