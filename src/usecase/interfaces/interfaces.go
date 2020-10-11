package interfaces

import "github.com/gashjp1994/go-ca/model/domain"

type UserInterface interface {
	Store(domain.User) (int, error)
	FindAll() ([]domain.User, error)
}
