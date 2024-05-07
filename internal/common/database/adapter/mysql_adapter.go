package adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"login_api/internal/common/database"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLAdapter struct {
	DB *sql.DB
}

func NewMySQLAdapter(db *sql.DB) database.DatabaseOperation {
	return &MySQLAdapter{DB: db}
}
func (adapter *MySQLAdapter) ExecTX(ctx context.Context, query string, param ...interface{}) (*sql.Result, error) {
	tx, err := adapter.DB.Begin()

	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(query, param...)
	fmt.Println(err)
	if err != nil {

		tx.Rollback()

		return nil, errors.New("Failed to execute transaction , transaction rolled back. Reason: " + err.Error() + "\r\n")

	}
	inserted_id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(inserted_id)
	}

	return &result, nil
}
func (adpter *MySQLAdapter) Exec(ctx context.Context, query string, param ...interface{}) (*sql.Result, error) {
	result, err := adpter.DB.Exec(query, param...)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (adapter *MySQLAdapter) Query(ctx context.Context, query string, param ...interface{}) (*sql.Rows, error) {
	result, _ := adapter.DB.Query(query, param...)
	return result, nil
}
func (adapter *MySQLAdapter) QueryRow(ctx context.Context, query string, param ...interface{}) (*sql.Row, error) {
	result := adapter.DB.QueryRow(query, param...)
	return result, nil
}
