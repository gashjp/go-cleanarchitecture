package modelif

import "github.com/gashjp1994/go-ca/model/domain"

type UserDBInterface interface {
	Store(u domain.User) (int, error)
	FindAll() ([]domain.User, error)
}
