package domain

import (
	"login_api/internal/auth/repository"
	"login_api/internal/common/config"

	"github.com/sirupsen/logrus"
)

type AuthDomain struct {
	logger     logrus.Entry
	config     config.Config
	repository repository.AuthRepository
}

func NewAuthDomain(logger logrus.Entry, config config.Config, repository repository.AuthRepository) AuthDomain {
	return AuthDomain{
		logger:     logger,
		config:     config,
		repository: repository,
	}
}
