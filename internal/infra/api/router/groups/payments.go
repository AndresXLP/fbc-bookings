package groups

import (
	"fbc-bookings/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type Payments interface {
	Resource(e *echo.Group)
}

type payments struct {
	paymentsHandler handler.Payments
}

func NewPaymentsGroup(paymentsHand handler.Payments) Payments {
	return &payments{paymentsHand}
}

func (group payments) Resource(e *echo.Group) {
	groupPath := e.Group("/payments")
	groupPath.POST("/wompi/transactions", group.paymentsHandler.ProcessPayment)
}
