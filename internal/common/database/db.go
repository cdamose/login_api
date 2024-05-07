package database

import (
	"context"
	"database/sql"
)

type DatabaseOperation interface {
	ExecTX(context.Context, string, ...interface{}) (*sql.Result, error)
	Exec(context.Context, string, ...interface{}) (*sql.Result, error)
	Query(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRow(context.Context, string, ...interface{}) (*sql.Row, error)
}
