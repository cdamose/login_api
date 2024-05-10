package repository

import (
	"context"
	"login_api/internal/auth/model/dao"
)

type AuthRepository interface {
	CheckMobileNumberAlredayExists(ctx context.Context, mobile_number string) (bool, error)
	CreateUserProfile(ctx context.Context, mobile_number string) (*dao.UserProfile, error)
	GetValidOTPDetails(ctx context.Context, user_id string, otp_code string) (*dao.OTPDetails, error)
	UpdateUserVerfiedStatus(ctx context.Context, user_id string, status bool) (bool, error)
	GenerateOTP(ctx context.Context, phone_number string, otp_code string) (bool, error)
	//PersistOTP(ctx context.Context, userId string, mobile_number string, OTP string) (bool, error)
	//VerifyOTP(ctx context.Context, userId string, mobile_number string) (bool, error)
	//GetProfile(ctx context.Context, mobile_number string) (dao.UserProfile, error)
}
