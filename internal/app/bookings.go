package app

import (
	"context"
	"net/http"
	"time"

	"fbc-bookings/internal/domain/dto"
	"fbc-bookings/internal/domain/entity"
	"fbc-bookings/internal/domain/ports/postgres/repo"
	"fbc-bookings/internal/infra/adapters/postgres/model"
	"github.com/labstack/echo/v4"
)

type Bookings interface {
	GetBookingsDate(ctx context.Context, date time.Time) ([]dto.BookingDateResponse, error)
	GetBookingsDateWithReservedPlaces(ctx context.Context, date time.Time) (dto.BookingsDatesWithReservedPlacesResponse, error)
	CreateBookingsDates(ctx context.Context, newBookings dto.BookingsDatesRequest) ([]dto.BookingDateResponse, error)
	CreateBooking(ctx context.Context, bookingRequest dto.BookingRequest) (dto.BookingResponse, error)
}

type bookings struct {
	repo repo.Repository
}

func NewBookingApp(repo repo.Repository) Bookings {
	return &bookings{repo}
}

func (app *bookings) GetBookingsDate(ctx context.Context, date time.Time) ([]dto.BookingDateResponse, error) {
	return app.repo.GetBookingsDate(ctx, date)
}

func (app *bookings) GetBookingsDateWithReservedPlaces(ctx context.Context, date time.Time) (dto.BookingsDatesWithReservedPlacesResponse, error) {
	bookingsDate, err := app.repo.GetBookingsDateWithReservedPlaces(ctx, date)
	if err != nil {
		return dto.BookingsDatesWithReservedPlacesResponse{}, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if len(bookingsDate) == 0 {
		return dto.BookingsDatesWithReservedPlacesResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{
			Message: "No Bookings",
		})
	}

	return bookingsDate, nil
}

func (app *bookings) CreateBookingsDates(ctx context.Context, newBookingsDates dto.BookingsDatesRequest) ([]dto.BookingDateResponse, error) {
	bookingsDates, _ := app.GetBookingsDate(ctx, newBookingsDates.BookingsDates[0].CheckIn)
	if len(bookingsDates) > 0 {
		return []dto.BookingDateResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "Ya existen fechas de reservas para el dia seleccionado"})
	}

	var newBookingsDateModel model.NewBookingsDates
	newBookingsDateModel.BuildModelCreateBookingsDate(newBookingsDates)
	bookingsDatesCreated, err := app.repo.CreateBookingsDates(ctx, newBookingsDateModel)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: err.Error()})
	}

	return bookingsDatesCreated, nil
}

func (app *bookings) CreateBooking(ctx context.Context, bookingRequest dto.BookingRequest) (dto.BookingResponse, error) {
	var bookingRequestModel model.BookingRequest

	booking, err := app.repo.GetBookingByUserIDAndBookingDateID(ctx, bookingRequest.BookedBy.ID, bookingRequest.BookingDateId)

	if booking.ID != 0 {
		return dto.BookingResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{
			Message: "the user already has an active reservation for this day",
		})
	}

	bookingRequestModel.BuildModel(bookingRequest)
	newBooking, err := app.repo.CreateBooking(ctx, bookingRequestModel)
	if err != nil {
		return dto.BookingResponse{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: err.Error()})
	}

	return newBooking, nil
}
