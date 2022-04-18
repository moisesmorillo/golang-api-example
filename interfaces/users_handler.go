package interfaces

import "github.com/labstack/echo/v4"

type UsersHandler interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
}
