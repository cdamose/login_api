package app

import (
	"context"
	"login_api/internal/common/config"
	"login_api/internal/communication_svc/domain"

	"github.com/sirupsen/logrus"
)

type CommunicationApplication struct {
	logger logrus.Entry
	config config.Config
	domain domain.CommunicationDomain
}

func NewCommunicationApplication(logger logrus.Entry, config config.Config, domain domain.CommunicationDomain) CommunicationApplication {
	return CommunicationApplication{
		logger: logger,
		config: config,
		domain: domain,
	}
}

func (app CommunicationApplication) SendSMS(ctx context.Context, phone_number string, message string) (bool, error) {
	err := app.domain.SendSMS(ctx, phone_number, message)
	if err != nil {
		return false, err
	}
	return true, nil
}
