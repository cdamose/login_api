package container

import (
	"login_api/internal/communication_svc/app"
)

type Application struct {
	PingApplication app.PingApp
	AuthApplication app.AuthApp
}
