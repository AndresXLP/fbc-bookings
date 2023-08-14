package dto

import (
	"math"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type BookingDate struct {
	ID       int       `json:"id"`
	CheckIn  time.Time `json:"check_in" validate:"required"`
	CheckOut time.Time `json:"check_out" validate:"required"`
	Places   int       `json:"places" validate:"required"`
}

type BookingsDatesRequest struct {
	BookingsDates []BookingDate `json:"bookings_dates" validate:"min=1,dive"`
}

func (b *BookingsDatesRequest) Validate() error {
	return validate.Struct(b)
}

type BookingRequest struct {
	ID            int    `json:"id"`
	BookingDateId int    `json:"booking_date_id" validate:"required"`
	BookedBy      User   `json:"booked_by" validate:"required"`
	BookingUsers  []User `json:"booking_users" validate:"required,unique"`
	Terms         *bool  `json:"terms" validate:"eq=true"`
}

func (n *BookingRequest) Validate() error {
	n.setAge()

	if err := n.validateBookingUsers(); err != nil {
		return err
	}

	return validate.Struct(n)
}

func (n *BookingRequest) validateBookingUsers() error {
	for _, user := range n.BookingUsers {
		if user.Age >= 18 {
			if err := validate.Struct(&user); err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *BookingRequest) setAge() {
	n.BookedBy.Age = n.getIntAge(n.BookedBy.Birthdate)
	for i, user := range n.BookingUsers {
		n.BookingUsers[i].Age = n.getIntAge(user.Birthdate)
	}
}

func (n *BookingRequest) getIntAge(birthday time.Time) int {
	today := time.Now()
	return int(math.Floor(today.Sub(birthday).Hours() / 24 / 365))
}
