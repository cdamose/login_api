package domain

import (
	"context"
	"login_api/internal/common/config"
	"login_api/internal/communication_svc/repository"

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
	err := ad.repository.SendSMS(ctx, phone_number, message)
	if err != nil {
		return err
	}
	return nil

}
