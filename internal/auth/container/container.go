package container

import (
	"login_api/internal/common/config"

	"github.com/jmoiron/sqlx"
)

func InitApplication(config config.Config, db *sqlx.DB) (Application, error) {
	ping_app, err := InitializePingApplication(config, db)
	//urlshortner_app, err := InitializeURLShortnerApplication(config, db)
	return Application{
		PingApplication: ping_app,
		//URLShortenApplication: urlshortner_app,
	}, err

}
