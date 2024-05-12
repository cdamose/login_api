package adapters

import (
	"context"

	"login_api/internal/common/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"
)

type PostgresCommunicatinRepository struct {
	db     *sqlx.DB
	logger logrus.Entry
	config config.Config
}

func NewPostgresCommunicationRepository(db *sqlx.DB, logger logrus.Entry, config config.Config) *PostgresCommunicatinRepository {
	if db == nil {
		panic("missing db")
	}
	return &PostgresCommunicatinRepository{db: db, logger: logger, config: config}
}
func (m PostgresCommunicatinRepository) SendSMS(ctx context.Context, phone_number string, message string) error {
	//write code to make databas entry
	return nil
}
