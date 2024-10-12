package events

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type BookingReservedDomainEvent struct {
	BookingId shared.UUID
}

func NewBookingReservedDomainEvent(bookingId shared.UUID) *BookingReservedDomainEvent {
	return &BookingReservedDomainEvent{bookingId}
}
