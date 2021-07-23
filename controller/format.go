package controller

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/service"
	"github.com/labstack/echo/v4"
)

// FormatController is a controller for managing format data.
type FormatController struct {
	context appcontext.Context
	service *service.FormatService
}

// NewFormatController is constructor.
func NewFormatController(context appcontext.Context) *FormatController {
	return &FormatController{context: context, service: service.NewFormatService(context)}
}

// GetFormatList returns the list of all formats.
func (controller *FormatController) GetFormatList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllFormats())
}
