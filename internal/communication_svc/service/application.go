package auth

// import (
// 	"context"
// 	"fmt"
// 	"login_api/internal/auth/adapters"
// 	"login_api/internal/auth/app"
// 	"login_api/internal/auth/domain"
// 	"login_api/internal/common/metrics"

// 	"github.com/sirupsen/logrus"
// )

// func NewApplication(ctx context.Context) app.Application {

// 	logger := logrus.NewEntry(logrus.StandardLogger())
// 	metricsClient := metrics.NoOp{}
// 	//setup depedency injection here
// 	db, err := adapters.NewMySQLConnection()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ping_domain := domain.PingDomain{}
// 	health_check_repo := adapters.NewMySQLPINGRepository(db, ping_domain)

// 	return app.Application{
// 		Queries: app.Queries{
// 			HealthCheck: query.NewHealthCheckHandler(health_check_repo, logger, metricsClient)},
// 	}
// }
