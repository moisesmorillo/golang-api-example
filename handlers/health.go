package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Message string `json:"message"`
}

// HealthCheck godoc
// @Tags Health
// @Summary Check if service is available
// @Description health service
// @Produce  json
// @Success 200 {object} Health
// @Router /health [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, &Health{
		Message: "Available!",
	})
}
