package utils

import (
	"strconv"
	"testing"
)

func TestGenerateOTP(t *testing.T) {
	otp := GenerateOTP()
	if len(otp) != 6 {
		t.Errorf("Generated OTP has length %d, expected 6", len(otp))
	}
	_, err := strconv.Atoi(otp)
	if err != nil {
		t.Errorf("Generated OTP is not a valid integer: %v", err)
	}
}
