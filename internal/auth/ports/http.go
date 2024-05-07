package ports

import (
	"fmt"

	"login_api/internal/auth/container"
	"net/http"

	"github.com/go-chi/render"
)

type HttpServer struct {
	Application container.Application
}

func NewHttpServer(application container.Application) HttpServer {

	return HttpServer{Application: application}
}

func (h HttpServer) GetPing(w http.ResponseWriter, r *http.Request) {
	dto_obj, _ := h.Application.PingApplication.Ping(r.Context())
	fmt.Println(dto_obj)
	response := PingResponse{}
	message := dto_obj.Message
	response.Message = &message
	render.Respond(w, r, response)
}
