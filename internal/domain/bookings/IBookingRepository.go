package bookings

import (
	"github.com/mehedimayall/bookify-go/internal/domain/apartments"
)

type IBookingRepository interface {
	Add(booking *Booking) string
	GetById(id string) *Booking
	IsOverlaping(apartment apartments.Apartment, duration DateRange) bool
}
