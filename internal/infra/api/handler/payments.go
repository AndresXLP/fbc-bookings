package handler

import (
	"net/http"

	"fbc-bookings/internal/domain/dto"
	"fbc-bookings/internal/domain/entity"
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

	if request.ValidateChecksum() {
		// generar los cambios respectivos en la base de datos

		return c.JSON(http.StatusOK, entity.Response{
			Message: "Transaction processed successfully",
		})
	}

	return c.JSON(http.StatusUnprocessableEntity, entity.Response{
		Message: "invalid checksum",
	})
}
