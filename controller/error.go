package controller

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/labstack/echo/v4"
)

// APIError has a error code and a message.
type APIError struct {
	Code    int
	Message string
}

// ErrorController is a controller for handling errors.
type ErrorController struct {
	context appcontext.Context
}

// NewErrorController is constructor.
func NewErrorController(context appcontext.Context) *ErrorController {
	return &ErrorController{context: context}
}

// JSONError is cumstomize error handler
func (controller *ErrorController) JSONError(err error, c echo.Context) {
	logger := controller.context.GetLogger()
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}

	var apierr APIError
	apierr.Code = code
	apierr.Message = msg

	if !c.Response().Committed {
		if reserr := c.JSON(code, apierr); reserr != nil {
			logger.GetZapLogger().Errorf(reserr.Error())
		}
	}
	logger.GetZapLogger().Debugf(err.Error())
}
