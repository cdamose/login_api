package repository

import (
	"context"
	"login_api/internal/communication_svc/model/dao"
)

type Repository interface {
	Ping(ctx context.Context) (*dao.Ping, error)
}
