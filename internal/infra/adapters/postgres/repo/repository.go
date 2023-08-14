package repo

import (
	"context"
	"time"

	"fbc-bookings/internal/domain/dto"
	"fbc-bookings/internal/domain/entity"
	"fbc-bookings/internal/domain/ports/postgres/repo"
	"fbc-bookings/internal/infra/adapters/postgres/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &repository{db}
}

func (repo repository) GetBookingsDate(ctx context.Context, date time.Time) ([]dto.BookingDateResponse, error) {
	var bookings model.BookingsDates

	repo.db.WithContext(ctx).
		Table("booking_dates").
		Scan(&bookings)

	return bookings.ToDomainDTO(), nil
}

func (repo repository) GetBookingsDateWithReservedPlaces(ctx context.Context, date time.Time) (dto.BookingsDatesWithReservedPlacesResponse, error) {
	var bookings model.BookingsDatesWithReservedPlaces

	err := repo.db.WithContext(ctx).
		Table("booking_dates as bd").
		Select("bd.id as id, bd.check_in as check_in, bd.check_out as check_out, bd.places as places, COALESCE(SUM(b.reserved_places), 0) AS count").
		Joins("LEFT JOIN bookings b ON b.booking_date_id = bd.id AND DATE(bd.check_in) = ?", date.String()[:10]).
		Where("DATE(bd.check_in) = ?", date.String()[:10]).
		Group("bd.id, bd.check_in").
		Order("bd.check_in").
		Scan(&bookings).Error
	if err != nil {
		return dto.BookingsDatesWithReservedPlacesResponse{}, err
	}

	return bookings.ToDomainDTO(), nil
}

func (repo repository) CreateBookingsDates(ctx context.Context, newBookingsDate model.NewBookingsDates) ([]dto.BookingDateResponse, error) {
	var bookingsDatesCreated []dto.BookingDateResponse
	err := repo.db.WithContext(ctx).
		Table("booking_dates").
		Create(&newBookingsDate).
		Scan(&bookingsDatesCreated).
		Error
	if err != nil {
		return []dto.BookingDateResponse{}, err
	}

	return bookingsDatesCreated, nil
}

func (repo repository) CreateBooking(ctx context.Context, bookingRequest model.BookingRequest) (dto.BookingResponse, error) {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := repo.GetAndUpdateOrCreateUserWithTx(ctx, tx, &bookingRequest.BookedBy); err != nil {
			return err
		}

		newBooking := model.Booking{
			BookingDateId:  bookingRequest.BookingDateId,
			BookedById:     bookingRequest.BookedBy.ID,
			ReservedPlaces: bookingRequest.BookingPlaces,
			Status:         "PENDING",
		}

		if err := tx.WithContext(ctx).Create(&newBooking).
			Scan(&newBooking).
			Error; err != nil {
			tx.Rollback()
			return err
		}

		for i, _ := range bookingRequest.BookingUsers {
			if err := repo.GetAndUpdateOrCreateUserWithTx(ctx, tx, &bookingRequest.BookingUsers[i]); err != nil {
				return err
			}
		}

		var bookingUsers model.BookingUsers
		bookingUsers.BuildModel(newBooking.ID, bookingRequest.BookingUsers)

		if err := tx.WithContext(ctx).
			Table("booking_users").Create(&bookingUsers).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	}); err != nil {
		return dto.BookingResponse{}, err
	}

	return bookingRequest.ToDomainDTO(), nil
}

func (repo repository) GetAndUpdateOrCreateUserWithTx(ctx context.Context, tx *gorm.DB, user *model.User) error {
	err := tx.WithContext(ctx).Where(model.User{DocumentID: user.DocumentID}).
		Table("users").
		FirstOrCreate(user).
		Omit("document_id,names,surnames,birthdate").
		Updates(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo repository) GetBookingByUserIDAndBookingDateID(ctx context.Context, userID int, bookingDateID int) (entity.Booking, error) {
	booking := model.Booking{}

	if err := repo.db.WithContext(ctx).
		Where(model.Booking{BookedById: userID, BookingDateId: bookingDateID}).
		Table("bookings").
		Scan(&booking).Error; err != nil {
		return entity.Booking{}, err
	}

	return booking.ToDomainEntity(), nil
}
