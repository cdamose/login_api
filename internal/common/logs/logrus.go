package logs

import (
	"os"
	"strconv"

	"login_api/internal/common/config"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Init(config config.Config) logrus.Entry {
	SetFormatter(logrus.StandardLogger())

	logrus.SetLevel(logrus.DebugLevel)
	logger := logrus.WithFields(
		logrus.Fields{
			"extra_field_one": "extra_value_one",
		})
	return *logger
}

func SetFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if isLocalEnv, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocalEnv {
		logger.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})
	}
}
