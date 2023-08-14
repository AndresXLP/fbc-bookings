package dto

import "time"

type BookingDateResponse struct {
	ID       int       `json:"id"`
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	Places   int       `json:"places"`
}

type BookingsDatesWithReservedPlacesResponse []BookingsDatesWithReservedPlaces

type BookingsDatesWithReservedPlaces struct {
	BookingDateResponse
	ReservedPlaces int `json:"reserved_places"`
}

type BookingResponse struct {
	BookingDateID int    `json:"booking_date_id"`
	BookedBy      User   `json:"booked_by"`
	BookingUsers  []User `json:"booking_users"`
	Places        int    `json:"places"`
}
