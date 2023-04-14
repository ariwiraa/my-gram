package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) string {
	salt := 8
	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(password), salt)

	return string(hashingPassword)
}

func ComparePass(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
