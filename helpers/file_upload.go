package helpers

import (
	"io"
	"mime/multipart"
	"net/http"
)

const (
	MAX_FILE_SIZE = 2
)

func IsImageFile(file multipart.File) (bool, error) {
	// Baca maksimal 512 byte untuk mendeteksi tipe file
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return false, err
	}

	// Set posisi kembali ke awal file setelah membaca buffer
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}

	// Deteksi tipe file
	fileType := http.DetectContentType(buffer)

	// Periksa tipe file yang didukung
	supportedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	if supportedTypes[fileType] {
		return true, nil
	}

	// Tipe file tidak didukung
	return false, ErrFileNotSupported
}

func IsFileSizeValid(file *multipart.FileHeader) error {
	// Batas maksimum ukuran file dalam byte
	maxSize := int64(MAX_FILE_SIZE) * 1024 * 1024

	// Baca ukuran file
	fileSize := file.Size

	// Periksa apakah ukuran file melebihi batas
	if fileSize > maxSize {
		return errFileSizeNotValid
	}

	return nil
}
