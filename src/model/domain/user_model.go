package domain

import (
	"github.com/gashjp1994/go-ca/model"
	"github.com/gashjp1994/go-ca/model/modelif"
	"github.com/gashjp1994/go-ca/utils"
)

type UserModel struct {
	UserDB modelif.UserDBInterface
	Logger utils.Logger
}

func (i *UserModel) Store(u model.User) (int, error) {
	return i.UserDB.Store(u)
}

func (i *UserModel) FindAll() ([]model.User, error) {
	return i.UserDB.FindAll()
}
