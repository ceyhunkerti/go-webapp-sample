package service

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
	"golang.org/x/crypto/bcrypt"
)

// AccountService is a service for managing user account.
type AccountService struct {
	context appcontext.Context
}

// NewAccountService is constructor.
func NewAccountService(context appcontext.Context) *AccountService {
	return &AccountService{context: context}
}

// AuthenticateByUsernameAndPassword authenticates by using username and plain text password.
func (a *AccountService) AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.Account) {
	rep := a.context.GetRepository()
	logger := a.context.GetLogger()
	account := model.Account{}
	result, err := account.FindByName(rep, username)
	if err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	return true, result
}
