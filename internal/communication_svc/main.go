package main

import (
	"log"

	"login_api/internal/communication_svc/container"
	"login_api/internal/communication_svc/ports"
	"login_api/internal/communication_svc/repository/adapters"

	//"login_api/internal/communication_svc/ports"
	"login_api/internal/common/config"
	"login_api/internal/common/database/migrations"

	//"login_api/internal/common/server"
	"net/http"

	"login_api/internal/common/genproto/communication/api/protobuf/communicationv1connect"

	//"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {

	config := config.InitConfig()
	db, err := adapters.NewPostgreSQLConnection()
	if err != nil {
		log.Fatalln(err)
	}
	application, err := container.InitApplication(config, db)
	//execute database migration files
	migrations.ExecMigration(config.MYSQLUser, config.MYSQLPassword, config.MYSQLHost, config.MYSQLDatabase, "file://./db/migrations/communication_svc")
	auther := ports.NewCommunicationServer(application)
	mux := http.NewServeMux()
	path, handler := communicationv1connect.NewCommunicationServiceHandler(auther)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":"+config.Port,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
