package utils

import (
	"math/rand"
	"strconv"
)

func GenerateOTP() string {
	otp := rand.Intn(999999)

	return strconv.Itoa(otp)
}
