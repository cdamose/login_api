package ports

import (
	"context"
	communicationv1 "login_api/internal/common/genproto/communication/api/protobuf"
	"login_api/internal/communication_svc/container"

	"connectrpc.com/connect"
)

type CommunicationServer struct {
	Application container.Application
}

func NewCommunicationServer(application container.Application) *CommunicationServer {
	return &CommunicationServer{Application: application}

}
func (av *CommunicationServer) Ping(ctx context.Context, req *connect.Request[communicationv1.Request]) (*connect.Response[communicationv1.Response], error) {
	return nil, nil
}
