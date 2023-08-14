package groups

import (
	"fbc-bookings/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type Bookings interface {
	Resource(c *echo.Group)
}

type bookings struct {
	bookingsHandler handler.Bookings
}

func NewBookingGroup(bookingHand handler.Bookings) Bookings {
	return &bookings{bookingHand}
}

func (group bookings) Resource(c *echo.Group) {
	groupPath := c.Group("/bookings")
	groupPath.GET("", group.bookingsHandler.GetBookingsDate)
	groupPath.POST("", group.bookingsHandler.CreateBookingsDates)
	groupPath.POST("/create", group.bookingsHandler.CreateBooking)
}
