//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	"login_api/internal/auth/app"
	"login_api/internal/auth/domain"
	"login_api/internal/auth/repository"
	"login_api/internal/auth/repository/adapters"
	"login_api/internal/common/config"
	"login_api/internal/common/logs"
)

func InitializeDomain(config config.Config, db *sqlx.DB) (app.PingApplication, error) {
	wire.Build(app.NewPingApplication, domain.NewPingDomain, adapters.NewMySQLPINGRepository, logs.Init,
		wire.Bind(new(repository.Repository), new(*adapters.MySQLPingRepository)))
	//wire.Build(config.InitConfig)
	return app.PingApplication{}, nil
}
