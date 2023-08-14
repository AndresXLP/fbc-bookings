package entity

type Booking struct {
	ID             int
	BookingDateID  int
	BookedByID     int
	ReservedPlaces int
	Status         string
}
