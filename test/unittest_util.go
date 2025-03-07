package test

import (
	"encoding/json"
	"fmt"

	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/config"
	"github.com/Screen17/catalog/logger"
	"github.com/Screen17/catalog/middleware"
	"github.com/Screen17/catalog/migration"
	"github.com/Screen17/catalog/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Prepare func is to prepare for unit test.
func Prepare() (*echo.Echo, appcontext.Context) {
	e := echo.New()

	conf := &config.Config{}
	conf.Database.Dialect = "sqlite3"
	conf.Database.Host = "file::memory:?cache=shared"
	conf.Database.Migration = true
	conf.Extension.MasterGenerator = true
	conf.Extension.SecurityEnabled = false
	conf.Log.RequestLogFormat = "${remote_ip} ${account_name} ${uri} ${method} ${status}"

	logger := initTestLogger()
	rep := repository.NewCatalogRepository(logger, conf)
	context := appcontext.NewContext(rep, conf, logger)

	middleware.InitLoggerMiddleware(e, context)

	migration.CreateDatabase(context)
	migration.InitMasterData(context)

	middleware.InitSessionMiddleware(e, context)
	return e, context
}

func initTestLogger() *logger.Logger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	myConfig := zap.Config{
		Level:       level,
		Encoding:    "console",
		Development: true,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zap, err := myConfig.Build()
	if err != nil {
		fmt.Printf("Error")
	}
	sugar := zap.Sugar()
	// set package varriable logger.
	logger := &logger.Logger{Zap: sugar}
	logger.GetZapLogger().Infof("Success to read zap logger configuration")
	_ = zap.Sync()
	return logger
}

// ConvertToString func is convert model to string.
func ConvertToString(model interface{}) string {
	bytes, _ := json.Marshal(model)
	return string(bytes)
}
