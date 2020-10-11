package modelif

import (
	"github.com/gashjp1994/go-ca/model"
)

type UserDBInterface interface {
	Store(u model.User) (int, error)
	FindAll() ([]model.User, error)
}
