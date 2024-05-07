package main

import (
	"context"
	"fmt"
	"log"

	"login_api/internal/auth/container"
	//"login_api/internal/auth/ports"
	"login_api/internal/common/config"
	"login_api/internal/common/database/migrations"
	"login_api/internal/common/logs"

	//"login_api/internal/common/server"
	"net/http"

	authv1 "login_api/internal/common/genproto/auth/api/protobuf"
	"login_api/internal/common/genproto/auth/api/protobuf/authv1connect"

	"connectrpc.com/connect"
	//"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type AuthServer struct {
	Application container.Application
}

func (av *AuthServer) SignupWithPhoneNumber(context.Context, *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.Response], error) {
	res := connect.NewResponse(&authv1.Response{
		Message: "success",
	})
	return res, nil
}
func (av *AuthServer) VerifyAccount(context.Context, *connect.Request[authv1.OTP]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}
func (av *AuthServer) Login(context.Context, *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}
func (av *AuthServer) OTPGenerate(context.Context, *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.Response], error) {
	res := connect.NewResponse(&authv1.Response{
		Message: "success",
	})
	return res, nil
}
func (av *AuthServer) GetProfile(context.Context, *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}

func main() {

	config := config.InitConfig()
	// currently we are establishing connection without connection pool , need to enchance with connection pool
	db, err := sqlx.Connect("mysql", config.MYSQLUser+":"+config.MYSQLPassword+"@("+config.MYSQLHost+":3306)/"+config.MYSQLDatabase)
	if err != nil {
		log.Fatalln(err)
	}

	application, err := container.InitApplication(config, db)

	fmt.Println(err)

	logger := logs.Init(config)
	fmt.Println(logger)
	fmt.Println(config.MYSQLHost)
	migrations.ExecMigration(config.MYSQLUser, config.MYSQLPassword, config.MYSQLHost, config.MYSQLDatabase)

	auther := &AuthServer{Application: application}
	mux := http.NewServeMux()
	path, handler := authv1connect.NewAuthServiceHandler(auther)
	fmt.Println("+++++++++++++++++++++++++")
	fmt.Println(path)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	// server.RunHTTPServer(func(router chi.Router) http.Handler {
	// 	return ports.HandlerFromMux(
	// 		ports.NewHttpServer(application),
	// 		router,
	// 	)
	// })

	// logger := logrus.NewEntry(logrus.StandardLogger())

	// ctx := context.Background()

	// application := service.NewApplication(ctx)
}
