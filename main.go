package main

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/config"
	"github.com/Screen17/catalog/logger"
	"github.com/Screen17/catalog/middleware"
	"github.com/Screen17/catalog/migration"
	"github.com/Screen17/catalog/repository"
	"github.com/Screen17/catalog/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	conf, env := config.Load()
	logger := logger.NewLogger(env)
	logger.GetZapLogger().Infof("Loaded this configuration : application." + env + ".yml")

	rep := repository.NewBookRepository(logger, conf)
	context := appcontext.NewContext(rep, conf, logger)

	migration.CreateDatabase(context)
	migration.InitMasterData(context)

	routes.Init(e, context)
	middleware.InitLoggerMiddleware(e, context)
	middleware.InitSessionMiddleware(e, context)

	if conf.StaticContents.Path != "" {
		e.Static("/", conf.StaticContents.Path)
		logger.GetZapLogger().Infof("Served the static contents. path: " + conf.StaticContents.Path)
	}

	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}

	defer rep.Close()
}
