package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Server is Working (:")
}
