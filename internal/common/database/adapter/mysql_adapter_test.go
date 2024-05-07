package adapter_test

import (
	"context"
	"testing"

	"login_api/internal/common/database/adapter"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMySQLAdapter_ExecTX(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(1, 1)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WithArgs("value1", "value2").WillReturnResult(mockResult)
	mock.ExpectCommit()

	mySQLAdapter := adapter.NewMySQLAdapter(db)
	result, err := mySQLAdapter.ExecTX(context.Background(), "INSERT INTO table_name (column1, column2) VALUES (?, ?)", "value1", "value2")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestMySQLAdapter_Exec(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec("UPDATE").WithArgs("value1", "value2").WillReturnResult(mockResult)

	mySQLAdapter := adapter.NewMySQLAdapter(db)
	result, err := mySQLAdapter.Exec(context.Background(), "UPDATE table_name SET column1 = ? WHERE column2 = ?", "value1", "value2")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestMySQLAdapter_Query(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"column1", "column2"}).
		AddRow("value1", "value2")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mySQLAdapter := adapter.NewMySQLAdapter(db)
	result, err := mySQLAdapter.Query(context.Background(), "SELECT * FROM table_name")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestMySQLAdapter_QueryRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	row := sqlmock.NewRows([]string{"column1", "column2"}).
		AddRow("value1", "value2")

	mock.ExpectQuery("SELECT").WillReturnRows(row)

	mySQLAdapter := adapter.NewMySQLAdapter(db)
	result, err := mySQLAdapter.QueryRow(context.Background(), "SELECT * FROM table_name")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
