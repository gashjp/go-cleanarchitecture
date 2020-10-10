package interfaces

import "github.com/labstack/echo"

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	Logger() echo.Logger // TODO: とりま
}
