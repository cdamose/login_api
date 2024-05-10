package app

import (
	"context"
	"login_api/internal/auth/domain"
	"login_api/internal/auth/model/conversion"
	"login_api/internal/auth/model/dto"
	"login_api/internal/common/config"

	"github.com/sirupsen/logrus"
)

type AuthApplication struct {
	logger logrus.Entry
	config config.Config
	domain domain.AuthDomain
}

func NewAuthApplication(logger logrus.Entry, config config.Config, domain domain.AuthDomain) AuthApplication {
	return AuthApplication{
		logger: logger,
		config: config,
		domain: domain,
	}
}

func (app AuthApplication) SignUp(ctx context.Context, phone_number string) (*dto.UserProfile, error) {
	domain_obj, err := app.domain.CreateUserProfile(ctx, phone_number)
	if err != nil {
		return nil, err
	}
	dto_obj := conversion.ConvertToUpdatedUserProfile(*domain_obj)
	return &dto_obj, nil
}

func (app AuthApplication) VerifyAccount(ctx context.Context, user_id string, otp string) (*dto.VerifiedAccountResp, error) {
	var resp = &dto.VerifiedAccountResp{}
	domain_obj, err := app.domain.VerifyAccount(ctx, user_id, otp)

	if err != nil {
		resp.Message = "Account Not Able to verified , pls try again"
	}
	if domain_obj {
		resp.Message = "Account Verified Successfully..!"
	}
	return resp, err

}

func (app AuthApplication) GenerateOTP(ctx context.Context, phone_number string) (*dto.CommonResponse, error) {
	var resp = &dto.CommonResponse{}
	domain_obj, err := app.domain.GenerateOTP(ctx, phone_number)
	if err != nil {
		resp.Message = "Something went wrong as of now , try again later "
	}
	if domain_obj {
		resp.Message = "OTP Sent Successfully..!"
	} else {
		resp.Message = "Seems to be given mobile number not register with our system"
	}
	return resp, err

}

func (app AuthApplication) Login(ctx context.Context, phone_number string) (*dto.LoginResponse, error) {
	var resp = &dto.LoginResponse{}
	domain_obj, err := app.domain.Login(ctx, phone_number)
	if err != nil {
		resp.Message = "Something went wrong,please try again"
	}
	resp.UserID = *domain_obj
	return resp, err

}
