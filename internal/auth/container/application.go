package container

import (
	"login_api/internal/auth/app"
)

type Application struct {
	PingApplication       app.PingApp
	URLShortenApplication app.AuthApp
}
