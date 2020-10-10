package controllers

import (
	"net/http"

	"github.com/gashjp1994/go-ca/db"
	"github.com/gashjp1994/go-ca/model"
	"github.com/gashjp1994/go-ca/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type UserController struct {
	UserUsecase usecase.UserUsecase // TODO: ださいかもしらん
}

func NewUserController(conn *gorm.DB) *UserController {
	return &UserController{
		UserUsecase: usecase.UserUsecase{
			UserConnection: &db.UserConnection{
				Conn: conn,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) error {
	type (
		Request struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		Response struct {
			UserID int `json:"user_id"`
		}
	)
	req := Request{}
	if err := c.Bind(req); err != nil {
		return err
	}

	user := model.User{Name: req.Name, Email: req.Email}
	id, err := controller.UserUsecase.Add(c, user)
	if err != nil {
		return err
	}
	res := Response{UserID: id}
	return c.JSON(http.StatusOK, res)

}
