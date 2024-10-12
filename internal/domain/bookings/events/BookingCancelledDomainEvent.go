package events

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type BookingCancelledDomainEvent struct {
	BookingId shared.UUID
}

func NewBookingCancelledDomainEvent(bookingId shared.UUID) *BookingCancelledDomainEvent {
	return &BookingCancelledDomainEvent{bookingId}
}
