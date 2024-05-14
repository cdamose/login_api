package domain

import (
	"context"
	"fmt"
	"login_api/internal/auth/model/dao"
	"login_api/internal/auth/repository"
	"login_api/internal/common/config"
	messaging_broker "login_api/internal/common/messaging_broker"
	"login_api/internal/common/utils"

	"github.com/kataras/iris/v12/x/errors"
	"github.com/sirupsen/logrus"
)

type AuthDomain struct {
	logger     logrus.Entry
	config     config.Config
	repository repository.AuthRepository
	broker     messaging_broker.RappitMQBroker
}

func NewAuthDomain(logger logrus.Entry, config config.Config, repository repository.AuthRepository) AuthDomain {
	return AuthDomain{
		logger:     logger,
		config:     config,
		repository: repository,
	}
}

func (ad *AuthDomain) GetUserProfile(ctx context.Context, phone_number string) (*dao.UserProfile, error) {
	profile, err := ad.repository.GetUserProfile(ctx, phone_number)
	if err != nil {
		return nil, err
	}
	return profile, err

}
func (ad *AuthDomain) CreateUserProfile(ctx context.Context, phone_number string) (*dao.UserProfile, error) {
	is_account_already_exist, err := ad.repository.CheckMobileNumberAlredayExists(ctx, phone_number)
	if err != nil {
		return nil, err
	}
	if !is_account_already_exist {
		profile, err := ad.repository.CreateUserProfile(ctx, phone_number)
		if err != nil {
			return nil, err
		}
		return profile, err
	} else {
		return nil, errors.New("phone number alreday assoiated with account")
	}

}

func (ad *AuthDomain) VerifyAccount(ctx context.Context, user_id string, otp string) (bool, error) {
	_, err := ad.repository.GetValidOTPDetails(ctx, user_id, otp)
	if err != nil {
		return false, err
	}
	result, err := ad.repository.UpdateUserVerfiedStatus(ctx, user_id, true)
	if err != nil {
		return false, err
	}
	result, err = ad.repository.UpdateOTPUsedStatus(ctx, user_id, otp, true)
	if err != nil {
		return false, err
	}

	return result, err

}

func (ad *AuthDomain) GenerateOTP(ctx context.Context, phone_number string) (bool, error) {
	otp := utils.GenerateOTP()
	//publish otp and phone number to verification_topic
	messaging_broker.Publish("verification_topic", otp, phone_number)
	result, err := ad.repository.GenerateOTP(ctx, phone_number, otp)
	fmt.Println(err)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ad *AuthDomain) Login(ctx context.Context, phone_number string) (*string, error) {
	otp := utils.GenerateOTP()
	//publish otp and phone number to verification_topic
	messaging_broker.Publish("verification_topic", otp, phone_number)
	result, err := ad.repository.Login(ctx, phone_number, otp)
	ad.repository.RecordUserEvents(ctx, *result, "Login")
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return result, nil
}
