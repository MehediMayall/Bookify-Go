package events

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type BookingCompletedDomainEvent struct {
	BookingId shared.UUID
}

func NewBookingCompletedDomainEvent(bookingId shared.UUID) *BookingCompletedDomainEvent {
	return &BookingCompletedDomainEvent{bookingId}
}
