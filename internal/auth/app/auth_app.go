package app

import (
	"login_api/internal/auth/domain"
	"login_api/internal/common/config"

	"github.com/sirupsen/logrus"
)

type AuthApplication struct {
	logger logrus.Entry
	config config.Config
	domain domain.AuthDomain
}

func NewURLShortnerApplication(logger logrus.Entry, config config.Config, domain domain.AuthDomain) AuthApplication {
	return AuthApplication{
		logger: logger,
		config: config,
		domain: domain,
	}
}
