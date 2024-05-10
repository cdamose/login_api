package container

import (
	"login_api/internal/auth/app"
)

type Application struct {
	PingApplication app.PingApp
	AuthApplication app.AuthApp
}
