//go:build wireinject
// +build wireinject

package container

import (
	"login_api/internal/auth/app"
	"login_api/internal/auth/domain"
	"login_api/internal/auth/repository"
	"login_api/internal/auth/repository/adapters"
	"login_api/internal/common/config"
	"login_api/internal/common/logs"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

func InitializeDomain(config config.Config, db *sqlx.DB) (app.PingApp, error) {
	wire.Build(app.NewPingApplication, domain.NewPingDomain, adapters.NewPostgressPingRepository, logs.Init,
		wire.Bind(new(app.PingApp), new(app.PingApplication)),
		wire.Bind(new(repository.Repository), new(*adapters.MySQLPingRepository)))
	//wire.Build(config.InitConfig)
	return app.PingApplication{}, nil
}

func InitializePingApplication(config config.Config, db *sqlx.DB) (app.PingApp, error) {
	wire.Build(app.NewPingApplication, domain.NewPingDomain, adapters.NewPostgressPingRepository, logs.Init,
		wire.Bind(new(app.PingApp), new(app.PingApplication)),
		wire.Bind(new(repository.Repository), new(*adapters.MySQLPingRepository)))
	//wire.Build(config.InitConfig)
	return app.PingApplication{}, nil
}
