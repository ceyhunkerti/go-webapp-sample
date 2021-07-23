package controller

import (
	"net/http"

	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
	"github.com/Screen17/catalog/model/dto"
	"github.com/Screen17/catalog/service"
	"github.com/Screen17/catalog/session"
	"github.com/labstack/echo/v4"
)

// AccountController is a controller for managing user account.
type AccountController struct {
	context      appcontext.Context
	service      *service.AccountService
	dummyAccount *model.Account
}

// NewAccountController is constructor.
func NewAccountController(context appcontext.Context) *AccountController {
	return &AccountController{
		context:      context,
		service:      service.NewAccountService(context),
		dummyAccount: model.NewAccountWithPlainPassword("test", "test", 1),
	}
}

// GetLoginStatus returns the status of login.
func (controller *AccountController) GetLoginStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

// GetLoginAccount returns the account data of logged in user.
func (controller *AccountController) GetLoginAccount(c echo.Context) error {
	if !controller.context.GetConfig().Extension.SecurityEnabled {
		return c.JSON(http.StatusOK, controller.dummyAccount)
	}
	return c.JSON(http.StatusOK, session.GetAccount(c))
}

// Login is the method to login using username and password by http post.
func (controller *AccountController) Login(c echo.Context) error {
	dto := dto.NewLoginDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}

	account := session.GetAccount(c)
	if account == nil {
		authenticate, a := controller.service.AuthenticateByUsernameAndPassword(dto.UserName, dto.Password)
		if authenticate {
			_ = session.SetAccount(c, a)
			_ = session.Save(c)
			return c.JSON(http.StatusOK, a)
		}
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, account)
}

// Logout is the method to logout by http post.
func (controller *AccountController) Logout(c echo.Context) error {
	_ = session.SetAccount(c, nil)
	_ = session.Delete(c)
	return c.NoContent(http.StatusOK)
}
