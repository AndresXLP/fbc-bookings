package handler

import (
	"net/http"

	"fbc-bookings/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

func NotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, entity.Response{
		Message: "Sorry, page not found",
	})
}
