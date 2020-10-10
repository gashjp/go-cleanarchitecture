package usecase

import (
	"github.com/gashjp1994/go-ca/model"
	"github.com/gashjp1994/go-ca/usecase/interfaces"
)

type UserUsecase struct {
	UserConnection interfaces.UserInterface
}

func (i *UserUsecase) Add(c interfaces.Context, u model.User) (int, error) {
	c.Logger().Debug(u)
	return i.UserConnection.Store(u)
}
