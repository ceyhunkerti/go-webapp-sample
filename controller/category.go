package controller

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/service"
	"github.com/labstack/echo/v4"
)

// CategoryController is a controller for managing category data.
type CategoryController struct {
	context appcontext.Context
	service *service.CategoryService
}

// NewCategoryController is constructor.
func NewCategoryController(context appcontext.Context) *CategoryController {
	return &CategoryController{context: context, service: service.NewCategoryService(context)}
}

// GetCategoryList returns the list of all categories.
func (controller *CategoryController) GetCategoryList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllCategories())
}
