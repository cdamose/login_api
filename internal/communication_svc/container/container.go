package container

import (
	"login_api/internal/common/config"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func InitApplication(config config.Config, db *sqlx.DB) (Application, error) {
	ping_app, err := InitializePingApplication(config, db)
	log.Info(err)
	communication_app, err := InitializeCommunicationApplication(config, db)
	log.Info(err)
	return Application{
		PingApplication:          ping_app,
		CommunicationApplication: communication_app,
	}, err

}
