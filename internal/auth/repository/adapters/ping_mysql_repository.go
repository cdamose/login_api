package adapters

import (
	"context"

	"login_api/internal/auth/model/dao"
	"login_api/internal/common/config"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type mysqlPing struct {
	Pingable bool `db:"pingabale"`
}

type MySQLPingRepository struct {
	db     *sqlx.DB
	logger logrus.Entry
	config config.Config
}

func NewMySQLPINGRepository(db *sqlx.DB, logger logrus.Entry, config config.Config) *MySQLPingRepository {
	if db == nil {
		panic("missing db")
	}
	return &MySQLPingRepository{db: db, logger: logger, config: config}
}

// sqlContextGetter is an interface provided both by transaction and standard db connection
type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (m MySQLPingRepository) Ping(ctx context.Context) (*dao.Ping, error) {
	var ping_obj dao.Ping
	_, err := m.db.Exec("select 1")
	if err == nil {
		ping_obj = dao.Ping{
			Message: "i'm able to connect " + m.config.MYSQLDatabase,
		}
	}

	return &ping_obj, err
}

func NewMySQLConnection() (*sqlx.DB, error) {
	config := mysql.NewConfig()

	config.Net = "tcp"
	config.Addr = os.Getenv("MYSQL_ADDR")
	config.User = os.Getenv("MYSQL_USER")
	config.Passwd = os.Getenv("MYSQL_PASSWORD")
	config.DBName = os.Getenv("MYSQL_DATABASE")
	config.ParseTime = true // with that parameter, we can use time.Time in mysqlHour.Hour

	db, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to MySQL")
	}

	return db, nil
}
