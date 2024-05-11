package ports

import (
	"login_api/internal/communication_svc/container"
)

type HttpServer struct {
	Application container.Application
}

func NewHttpServer(application container.Application) HttpServer {

	return HttpServer{Application: application}
}
