package handler

import (
	"net/http"
	"time"

	"fbc-bookings/internal/app"
	"fbc-bookings/internal/domain/dto"
	"fbc-bookings/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type Bookings interface {
	GetBookingsDate(c echo.Context) error
	CreateBookingsDates(c echo.Context) error
	CreateBooking(c echo.Context) error
}

type bookings struct {
	app app.Bookings
}

func NewBookingHandler(app app.Bookings) Bookings {
	return &bookings{app}
}

func (hand *bookings) GetBookingsDate(c echo.Context) error {
	ctx := c.Request().Context()

	date := c.QueryParam("date")
	t, err := time.Parse("2006-01-02", date)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	bookingsDate, err := hand.app.GetBookingsDateWithReservedPlaces(ctx, t)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Bookings loaded successfully",
		Data:    bookingsDate,
	})
}

func (hand *bookings) CreateBookingsDates(c echo.Context) error {
	ctx := c.Request().Context()

	var bookingsDate dto.BookingsDatesRequest
	if err := c.Bind(&bookingsDate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{Message: err.Error()})
	}

	if err := bookingsDate.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	bookingsDatesCreated, err := hand.app.CreateBookingsDates(ctx, bookingsDate)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Bookings Date Created Successfully",
		Data:    bookingsDatesCreated,
	})
}

func (hand *bookings) CreateBooking(c echo.Context) error {
	ctx := c.Request().Context()
	request := dto.BookingRequest{}

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{
			Message: err.Error(),
		})
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{
			Message: err.Error(),
		})
	}

	newBooking, err := hand.app.CreateBooking(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Successful Booking",
		Data:    newBooking,
	})
}
