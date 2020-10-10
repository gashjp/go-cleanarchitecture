package usecase

import (
	logg "log"

	"github.com/gashjp1994/go-ca/model"
	"github.com/gashjp1994/go-ca/usecase/interfaces"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type UserUsecase struct {
	UserConnection interfaces.UserInterface
}

func (i *UserUsecase) Add(c echo.Context, u model.User) (int, error) {
	log.Debug("fugafugae")
	return i.UserConnection.Store(u)
}

func (i *UserUsecase) FindAll(c echo.Context) ([]model.User, error) {
	log.Debug("hogehoge")
	logg.Printf("hogehoge")
	return i.UserConnection.FindAll()
}
