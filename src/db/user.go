package db

import (
	"fmt"

	"github.com/gashjp1994/go-ca/model/domain"
	"github.com/jinzhu/gorm"
)

type (
	UserConnection struct {
		Conn *gorm.DB
	}
	User struct {
		gorm.Model
		Name  string `gorm:"size:20;not null"`
		Email string `gorm:"size:100;not null"`
		Age   int    `gorm:"type:smallint"`
	}
)

func (r *UserConnection) Store(u domain.User) (int, error) {
	user := &User{
		Name:  u.Name,
		Email: u.Email,
	}
	if err := r.Conn.Create(user).Error; err != nil {
		return 0, fmt.Errorf("create user error")
	}
	return int(user.ID), nil
}

func (r *UserConnection) FindAll() ([]domain.User, error) {
	var d []domain.User
	users := []User{}
	if err := r.Conn.Find(&users).Error; err != nil {
		return d, err
	}

	for i := 0; i < len(users); i++ {
		d = append(d, domain.User{
			ID:    int(users[i].ID),
			Name:  users[i].Name,
			Email: users[i].Email,
		})
	}
	return d, nil
}
