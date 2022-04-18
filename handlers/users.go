package handlers

import (
	"github.com/moisesmorillo/golang-api-example/db/dtos"
	"github.com/moisesmorillo/golang-api-example/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	usersService interfaces.UsersService
}

func NewUsersHandler(usersService interfaces.UsersService) interfaces.UsersHandler {
	return &usersHandler{
		usersService,
	}
}

// @Tags Users
// @Summary Get All Users
// @Description get all users with a limit of 10
// @Accept  json
// @Produce  json
// @Success 200 {object} []dtos.Users
// @Failure 400 {object} map[string]string
// @Router /users/ [get]
func (u usersHandler) Get(c echo.Context) error {
	users, err := u.usersService.Get(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]*[]dtos.Users{
		"data": users,
	})
}

// @Tags Users
// @Summary Create user
// @Description Create user
// @Accept  json
// @Produce  json
// @Param request body dtos.Users true "Request body"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /users/ [post]
func (u usersHandler) Create(c echo.Context) error {
	user := &dtos.Users{}

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := u.usersService.Create(c.Request().Context(), user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user successfully created",
	})
}
