package main

import (
	"fmt"
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
	// broker := messaging_broker.NewRappitMQBroker()

	// messageHandler := func(body []byte) {
	// 	// Process the incoming message from RabbitMQ here
	// 	fmt.Printf("Received message: %s\n", body)
	// }

	// // Subscribe to the RabbitMQ queue/topic
	// go broker.Subscribe("verification_topic", messageHandler)

	config := config.InitConfig()
	db, err := adapters.NewPostgreSQLConnection()
	if err != nil {
		log.Fatalln(err)
	}
	application, err := container.InitApplication(config, db)

	fmt.Println(err)
	//execute database migration files
	migrations.ExecMigration(config.MYSQLUser, config.MYSQLPassword, config.MYSQLHost, config.MYSQLDatabase, "file://./db/migrations/auth")
	auther := ports.NewAuthServer(application)
	mux := http.NewServeMux()
	path, handler := authv1connect.NewAuthServiceHandler(auther)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	// will move into common module
	//startRapitmqConsumenr()
	// fmt.Println("test  35")

	// fmt.Println("test  36")

	//broker.Subscribe("")

	// Wait for termination signal to gracefully shutdown
	// exitSignal := make(chan os.Signal, 1)
	// signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	// <-exitSignal

	// // Stop consuming messages from RabbitMQ
	// broker.Stop()

	// fmt.Println("Shutting down...")
}
