package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello is a basic status check route to make sure that the server is running
func (h *Handlers) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"msg": "hello!"})
}
