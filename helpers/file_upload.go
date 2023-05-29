package helpers

import (
	"mime/multipart"
	"net/http"
)

func IsImageFile(file multipart.File) bool {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return false
	}

	fileType := http.DetectContentType(buffer)
	return fileType == "image/jpeg" || fileType == "image/png"
}
