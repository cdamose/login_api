package ports

import (
	"login_api/internal/auth/container"
)

type HttpServer struct {
	Application container.Application
}

func NewHttpServer(application container.Application) HttpServer {

	return HttpServer{Application: application}
}
