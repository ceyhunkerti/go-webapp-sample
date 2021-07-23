package controller

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/labstack/echo/v4"
)

// HealthController is a controller returns the current status of this application.
type HealthController struct {
	context appcontext.Context
}

// NewHealthController is constructor.
func NewHealthController(context appcontext.Context) *HealthController {
	return &HealthController{context: context}
}

// GetHealthCheck returns whether this application is alive or not.
func (controller *HealthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
