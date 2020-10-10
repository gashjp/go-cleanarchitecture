package interfaces

import (
	"github.com/gashjp1994/go-ca/model"
)

type UserInterface interface {
	Store(model.User) (int, error)
	FindAll() ([]model.User, error)
}
