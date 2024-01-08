package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomOTP() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	otp := random.Intn(10000)
	if otp < 1000 {
		otp += 1000
	}

	return otp
}
