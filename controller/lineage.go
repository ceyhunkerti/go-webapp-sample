package controller

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/service"
	"github.com/labstack/echo/v4"
)

type LineageController struct {
	context appcontext.Context
	service *service.LineageService
}

func NewLineageController(context appcontext.Context) *LineageController {
	return &LineageController{context: context, service: service.NewLineageService(context)}
}

func (controller *LineageController) GetLineageList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllLineages())
}
