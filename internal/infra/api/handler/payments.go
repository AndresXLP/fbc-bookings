package handler

import (
	"net/http"

	"fbc-bookings/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

type Payments interface {
	ProcessPayment(c echo.Context) error
}

type payments struct {
}

func NewPaymentsHandler() Payments {
	return &payments{}
}

func (hand *payments) ProcessPayment(c echo.Context) error {
	request := dto.Payment{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	request.GetCheckSum256()
	return c.JSON(http.StatusOK, request)
}
