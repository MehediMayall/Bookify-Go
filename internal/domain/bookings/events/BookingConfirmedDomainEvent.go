package events

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type BookingConfirmedDomainEvent struct {
	BookingId shared.UUID
}

func NewBookingConfirmedDomainEvent(bookingId shared.UUID) *BookingConfirmedDomainEvent {
	return &BookingConfirmedDomainEvent{bookingId}
}
