package handler

import (
	"net/http"
	"os"

	"github.com/gashjp1994/go-ca/controllers"

	"github.com/gashjp1994/go-ca/db/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Setup() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = customHTTPErrorHandler

	conn, err := mysql.Connect()
	if err != nil {
		e.Logger.Fatal(err.Error())
		return
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/user", func(c echo.Context) error {
		controller := controllers.NewUserController(conn)
		return controller.Create(c)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	type (
		Response struct {
			Result int    `json:"result"`
			Msg    string `json:"msg"`
		}
	)
	// https://echo.labstack.com/guide/error-handling
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	res := Response{Result: code, Msg: err.Error()}
	c.JSON(http.StatusOK, res)
}
