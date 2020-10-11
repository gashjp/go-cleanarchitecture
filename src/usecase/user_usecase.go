package usecase

import (
	"github.com/gashjp1994/go-ca/model"
	"github.com/gashjp1994/go-ca/usecase/interfaces"
	"github.com/gashjp1994/go-ca/utils"
	"github.com/labstack/echo"
)

type UserUsecase struct {
	UserModel interfaces.UserInterface
	Logger    utils.Logger
}

func (i *UserUsecase) Add(c echo.Context, u model.User) (int, error) {
	i.Logger.Log("fugafugae")
	return i.UserModel.Store(u)
	// return i.UserConnection.Store(u)
}

func (i *UserUsecase) FindAll(c echo.Context) ([]model.User, error) {
	i.Logger.Log("hogehoge", "a")
	return i.UserModel.FindAll()
	//return i.UserConnection.FindAll()
}
