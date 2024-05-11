package app

import (
	"context"
	"login_api/internal/communication_svc/model/dto"
)

type AuthApp interface {
	SignUp(ctx context.Context, phone_number string) (*dto.UserProfile, error)
	VerifyAccount(ctx context.Context, user_id string, otp string) (*dto.VerifiedAccountResp, error)
	GenerateOTP(ctx context.Context, phone_number string) (*dto.CommonResponse, error)
	Login(ctx context.Context, phone_number string) (*dto.LoginResponse, error)
	GetUserProfile(ctx context.Context, phone_number string) (*dto.UserProfile, error)
}
