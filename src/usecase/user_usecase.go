package usecase

import (
	"github.com/gashjp1994/go-ca/model"
	"github.com/gashjp1994/go-ca/usecase/interfaces"
	"github.com/labstack/echo"
)

type UserUsecase struct {
	UserConnection interfaces.UserInterface
}

func (i *UserUsecase) Add(c echo.Context, u model.User) (int, error) {
	c.Logger().Debug("a", "b")
	return i.UserConnection.Store(u)
}
