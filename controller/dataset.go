package controller

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/service"
	"github.com/labstack/echo/v4"
)

type DatasetController struct {
	context appcontext.Context
	service *service.DatasetService
}

func NewDatasetController(context appcontext.Context) *DatasetController {
	return &DatasetController{context: context, service: service.NewDatasetService(context)}
}

func (controller *DatasetController) GetDatasetList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllDatasets())
}
