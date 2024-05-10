package dao

import "time"

type UserProfile struct {
	UserID       string     `db:"user_id"`
	IsVerified   bool       `db:"is_verified"`
	CreatedAt    time.Time  `db:"created_at"`
	VerifiedAt   *time.Time `db:"verified_at"`
	MobileNumber string     `db:"phone_number"`
}

type OTPDetails struct {
	OTPID      string     `db:"otp_id"`
	UserID     string     `db:"user_id"`
	OTPCode    string     `db:"otp_code"`
	IsUsed     bool       `db:"is_used"`
	CreatedAt  time.Time  `db:"created_at"`
	VerifiedAt *time.Time `db:"modified_at"`
}
