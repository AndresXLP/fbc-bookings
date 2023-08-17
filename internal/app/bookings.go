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
	GetBookingsDateByDate(ctx context.Context, date time.Time) ([]dto.BookingDateResponse, error)
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

func (app *bookings) GetBookingsDateByDate(ctx context.Context, date time.Time) ([]dto.BookingDateResponse, error) {
	return app.repo.GetBookingsDateByDate(ctx, date)
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
	bookingsDates, _ := app.GetBookingsDateByDate(ctx, newBookingsDates.BookingsDates[0].CheckIn)
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

	bookingDate, err := app.repo.GetBookingDateByID(ctx, bookingRequest.BookingDateId)
	if err != nil {
		return dto.BookingResponse{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: err.Error()})
	}

	if bookingDate.ID == 0 {
		return dto.BookingResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "booking date not found"})
	}

	bookingByUser, err := app.repo.GetBookingByUserIDAndBookingDateID(ctx, bookingRequest.BookedBy.ID, bookingRequest.BookingDateId)
	if err != nil {
		return dto.BookingResponse{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: err.Error()})
	}

	if bookingByUser.ID != 0 {
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
