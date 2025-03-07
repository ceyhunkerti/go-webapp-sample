package routes

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo, context appcontext.Context) {
	conf := context.GetConfig()
	if conf.Extension.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}))
	}

	errorHandler := controller.NewErrorController(context)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())

	dataset := controller.NewDatasetController(context)
	lineage := controller.NewLineageController(context)
	account := controller.NewAccountController(context)
	health := controller.NewHealthController(context)

	e.GET(controller.APIDatasets, func(c echo.Context) error { return dataset.GetDatasetList(c) })

	e.GET(controller.APILineages, func(c echo.Context) error { return lineage.GetLineageList(c) })


	e.GET(controller.APIAccountLoginStatus, func(c echo.Context) error { return account.GetLoginStatus(c) })
	e.GET(controller.APIAccountLoginAccount, func(c echo.Context) error { return account.GetLoginAccount(c) })

	if conf.Extension.SecurityEnabled {
		e.POST(controller.APIAccountLogin, func(c echo.Context) error { return account.Login(c) })
		e.POST(controller.APIAccountLogout, func(c echo.Context) error { return account.Logout(c) })
	}

	e.GET(controller.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })

}
