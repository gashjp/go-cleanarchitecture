package db

import (
	"fmt"

	"github.com/gashjp1994/go-ca/model"
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

func (r *UserConnection) Store(u model.User) (int, error) {
	user := &User{
		Name:  u.Name,
		Email: u.Email,
	}
	if err := r.Conn.Create(user).Error; err != nil {
		return 0, fmt.Errorf("create user error")
	}
	return int(user.ID), nil
}
