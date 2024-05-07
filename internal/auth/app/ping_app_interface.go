package app

import (
	"context"
	"login_api/internal/auth/model/dto"
)

type PingApp interface {
	Ping(ctx context.Context) (*dto.Ping, error)
}
