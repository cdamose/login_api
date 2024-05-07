package adapter

import (
	"context"
	"database/sql"
	"errors"
	"login_api/internal/common/database"
)

type PostgreSQLAdapter struct {
	DB *sql.DB
}

func NewPostgreSQLAdapter(db *sql.DB) database.DatabaseOperation {
	return &PostgreSQLAdapter{DB: db}
}

func (adapter *PostgreSQLAdapter) ExecTX(ctx context.Context, query string, param ...interface{}) (*sql.Result, error) {
	tx, err := adapter.DB.Begin()

	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(query, param...)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Failed to execute transaction, transaction rolled back. Reason: " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (adapter *PostgreSQLAdapter) Exec(ctx context.Context, query string, param ...interface{}) (*sql.Result, error) {
	result, err := adapter.DB.Exec(query, param...)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (adapter *PostgreSQLAdapter) Query(ctx context.Context, query string, param ...interface{}) (*sql.Rows, error) {
	result, err := adapter.DB.Query(query, param...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (adapter *PostgreSQLAdapter) QueryRow(ctx context.Context, query string, param ...interface{}) (*sql.Row, error) {
	result := adapter.DB.QueryRow(query, param...)
	return result, nil
}
