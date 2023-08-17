package handler

import (
	"fmt"
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

	checsum := request.GetCheckSum256()
	fmt.Println(checsum == request.Signature.Checksum)
	return c.JSON(http.StatusOK, request)
}
