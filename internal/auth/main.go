package main

import (
	"log"

	"login_api/internal/auth/container"
	"login_api/internal/auth/ports"
	"login_api/internal/auth/repository/adapters"

	//"login_api/internal/auth/ports"
	"login_api/internal/common/config"
	"login_api/internal/common/database/migrations"

	//"login_api/internal/common/server"
	"net/http"

	"login_api/internal/common/genproto/auth/api/protobuf/authv1connect"

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
	migrations.ExecMigration(config.MYSQLUser, config.MYSQLPassword, config.MYSQLHost, config.MYSQLDatabase)
	auther := ports.NewAuthServer(application)
	mux := http.NewServeMux()
	path, handler := authv1connect.NewAuthServiceHandler(auther)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
