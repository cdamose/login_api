package app

import (
	"context"

	"login_api/internal/communication_svc/domain"
	"login_api/internal/communication_svc/model/dto"
	"login_api/internal/common/config"

	"github.com/sirupsen/logrus"
)

type PingApplication struct {
	logger logrus.Entry
	config config.Config
	domain domain.PingDomain
}

func NewPingApplication(logger logrus.Entry, config config.Config, domain domain.PingDomain) PingApplication {
	return PingApplication{
		logger: logger,
		config: config,
		domain: domain,
	}
}
func (p PingApplication) Ping(ctx context.Context) (*dto.Ping, error) {
	return p.domain.Ping(ctx)

}
