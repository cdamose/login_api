package domain

import (
	"context"
	"fmt"
	"login_api/internal/common/config"
	"login_api/internal/communication_svc/repository"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/sirupsen/logrus"
)

type CommunicationDomain struct {
	logger     logrus.Entry
	config     config.Config
	repository repository.CommunicationRepository
}

func NewCommunicationDomain(logger logrus.Entry, config config.Config, repository repository.CommunicationRepository) CommunicationDomain {
	return CommunicationDomain{
		logger:     logger,
		config:     config,
		repository: repository,
	}
}

func (ad *CommunicationDomain) SendSMS(ctx context.Context, phone_number string, message string) error {
	client := twilio.NewRestClient()
	params := &api.CreateMessageParams{}
	params.SetBody(message)
	params.SetFrom("+12176694094")
	params.SetTo(phone_number)
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
	err = ad.repository.SendSMS(ctx, phone_number, message)
	if err != nil {
		return err
	}

	return nil

}
