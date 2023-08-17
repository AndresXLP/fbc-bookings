package repo

import (
	"context"
	"time"

	"fbc-bookings/internal/domain/dto"
	"fbc-bookings/internal/domain/entity"
	"fbc-bookings/internal/infra/adapters/postgres/model"
)

type Repository interface {
	GetBookingsDateByDate(ctx context.Context, date time.Time) ([]dto.BookingDateResponse, error)
	GetBookingsDateWithReservedPlaces(ctx context.Context, date time.Time) (dto.BookingsDatesWithReservedPlacesResponse, error)
	CreateBookingsDates(ctx context.Context, newBookingsDate model.NewBookingsDates) ([]dto.BookingDateResponse, error)
	CreateBooking(ctx context.Context, bookingRequest model.BookingRequest) (dto.BookingResponse, error)
	GetBookingByUserIDAndBookingDateID(ctx context.Context, userID int, bookingDateID int) (entity.Booking, error)
	GetBookingDateByID(ctx context.Context, bookingDateID int) (dto.BookingDateResponse, error)
}
