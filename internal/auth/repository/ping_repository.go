package repository

import (
	"context"
	"login_api/internal/auth/model/dao"
)

type Repository interface {
	Ping(ctx context.Context) (*dao.Ping, error)
}
