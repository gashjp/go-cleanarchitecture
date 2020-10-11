package usecase

import (
	"fmt"

	"github.com/gashjp1994/go-ca/model/domain"
	"github.com/gashjp1994/go-ca/usecase/interfaces"
	"github.com/gashjp1994/go-ca/utils"
	"github.com/labstack/echo"
)

type UserUsecase struct {
	UserModel interfaces.UserInterface
	Logger    utils.Logger
}

// func (i *UserUsecase) Add(c echo.Context, u domain.User) (int, error) {
func (i *UserUsecase) Add(c echo.Context, name, email string) (int, error) {
	// validation
	if name == "" {
		return 0, fmt.Errorf("name is empty")
	}
	if email == "" {
		return 0, fmt.Errorf("email is empty")
	}

	i.Logger.Log("fugafugae")
	return i.UserModel.Store(domain.User{
		Name:  name,
		Email: email,
	})
	// return i.UserConnection.Store(u)
}

func (i *UserUsecase) FindAll(c echo.Context) ([]domain.User, error) {
	i.Logger.Log("hogehoge", "a")
	return i.UserModel.FindAll()
	//return i.UserConnection.FindAll()
}
