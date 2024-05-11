package app

import (
	"context"
	"login_api/internal/communication_svc/model/dto"
)

type PingApp interface {
	Ping(ctx context.Context) (*dto.Ping, error)
}
