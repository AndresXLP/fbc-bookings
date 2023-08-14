package model

import (
	"time"

	"fbc-bookings/internal/domain/dto"
	"fbc-bookings/internal/domain/entity"
	"github.com/thoas/go-funk"
)

type BookingDate struct {
	ID        int
	CheckIn   time.Time
	CheckOut  time.Time
	Places    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookingDateWithReservedPlace struct {
	BookingDate
	Count int
}

func (b *BookingDate) ToDomainDTO() dto.BookingDateResponse {
	return dto.BookingDateResponse{
		ID:       b.ID,
		CheckIn:  b.CheckIn,
		CheckOut: b.CheckOut,
		Places:   b.Places,
	}
}

type BookingsDates []BookingDate

func (b *BookingsDates) ToDomainDTO() []dto.BookingDateResponse {
	var bookingDates []dto.BookingDateResponse

	for _, booking := range *b {
		bookingDates = append(bookingDates, booking.ToDomainDTO())
	}

	return bookingDates
}

func (b *BookingDateWithReservedPlace) toDomainDTO() dto.BookingsDatesWithReservedPlaces {
	return dto.BookingsDatesWithReservedPlaces{
		BookingDateResponse: dto.BookingDateResponse{
			ID:       b.ID,
			CheckIn:  b.CheckIn,
			CheckOut: b.CheckOut,
			Places:   b.Places,
		},
		ReservedPlaces: b.Count,
	}
}

type BookingsDatesWithReservedPlaces []BookingDateWithReservedPlace

func (b *BookingsDatesWithReservedPlaces) ToDomainDTO() dto.BookingsDatesWithReservedPlacesResponse {
	var bookingsDates dto.BookingsDatesWithReservedPlacesResponse
	for _, date := range *b {
		bookingsDates = append(bookingsDates, date.toDomainDTO())
	}

	return bookingsDates
}

type NewBookingsDates []BookingDate

func (b *NewBookingsDates) BuildModelCreateBookingsDate(newBookingsDate dto.BookingsDatesRequest) {
	for _, bookingDate := range newBookingsDate.BookingsDates {
		newBooking := BookingDate{
			CheckIn:  bookingDate.CheckIn,
			CheckOut: bookingDate.CheckOut,
			Places:   bookingDate.Places,
		}

		*b = append(*b, newBooking)
	}
}

type BookingRequest struct {
	BookingDateId int
	BookedBy      User
	BookingUsers  []User
	BookingPlaces int
}

func (b *BookingRequest) BuildModel(request dto.BookingRequest) {
	places := 1
	b.BookingDateId = request.BookingDateId
	b.BookedBy = User{
		ID:           request.BookedBy.ID,
		DocumentType: request.BookedBy.DocumentType,
		DocumentID:   request.BookedBy.DocumentID,
		Names:        request.BookedBy.Names,
		Surnames:     request.BookedBy.Surnames,
		Birthdate:    request.BookedBy.Birthdate,
		Age:          request.BookedBy.Age,
		Gender:       request.BookedBy.Gender,
		Email:        request.BookedBy.Email,
		Phone:        request.BookedBy.Phone,
	}
	b.BookingUsers = funk.Map(request.BookingUsers, func(v dto.User) User {
		if v.Age >= 18 {
			places++
		}
		return User{
			ID:           v.ID,
			DocumentType: v.DocumentType,
			DocumentID:   v.DocumentID,
			Names:        v.Names,
			Surnames:     v.Surnames,
			Birthdate:    v.Birthdate,
			Age:          v.Age,
			Gender:       v.Gender,
			Email:        v.Email,
			Phone:        v.Phone,
		}
	}).([]User)
	b.BookingPlaces = places
}

func (b *BookingRequest) ToDomainDTO() dto.BookingResponse {
	booking := dto.BookingResponse{
		BookingDateID: b.BookingDateId,
		Places:        b.BookingPlaces,
	}
	booking.BookedBy = b.BookedBy.ToDomainDTO()
	booking.BookingUsers = funk.Map(b.BookingUsers, func(v User) dto.User {
		return v.ToDomainDTO()
	}).([]dto.User)

	return booking
}

type Booking struct {
	ID             int
	BookingDateId  int
	BookedById     int
	ReservedPlaces int
	Status         string
}

func (b *Booking) ToDomainEntity() entity.Booking {
	return entity.Booking{
		ID:             b.ID,
		BookingDateID:  b.BookingDateId,
		BookedByID:     b.BookedById,
		ReservedPlaces: b.ReservedPlaces,
		Status:         b.Status,
	}
}

type bookingUser struct {
	BookingId int
	UserId    int
}

type BookingUsers []bookingUser

func (u *BookingUsers) BuildModel(bookingId int, users []User) {
	newBookingUser := bookingUser{
		BookingId: bookingId,
	}
	for _, user := range users {
		newBookingUser.UserId = user.ID
		*u = append(*u, newBookingUser)
	}
}
