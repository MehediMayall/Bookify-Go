package events

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type BookingRejectedDomainEvent struct {
	BookingId shared.UUID
}

func NewBookingRejectedDomainEvent(bookingId shared.UUID) *BookingRejectedDomainEvent {
	return &BookingRejectedDomainEvent{bookingId}
}
