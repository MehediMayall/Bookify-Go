package reservebooking

import (
	"github.com/mehedimayall/bookify-go/internal/domain/bookings"
	"github.com/mehedimayall/bookify-go/internal/domain/bookings/events"
	"github.com/mehedimayall/bookify-go/internal/domain/users"
	"github.com/mehedimayall/bookify-go/internal/usecases/abstractions"
)

type ReserveBookingDomainEventHandler struct {
	bookingRepo    bookings.IBookingRepository
	userRepo       users.IUserRepository
	pricingService bookings.PricingService
	mailService    abstractions.IEmailService
}

func NewReserveBookingDomainEventHandler(
	bookingRepo bookings.IBookingRepository,
	userRepo users.IUserRepository,
	pricingService bookings.PricingService,
	mailService abstractions.IEmailService) *ReserveBookingDomainEventHandler {

	return &ReserveBookingDomainEventHandler{
		bookingRepo,
		userRepo,
		pricingService,
		mailService,
	}
}

func (h *ReserveBookingDomainEventHandler) Handle(event events.BookingReservedDomainEvent) {
	booking := h.bookingRepo.GetById(event.BookingId.ToString())

	if booking == nil {
		return
	}

	userId := booking.GetUserId()

	user := h.userRepo.GetById(userId.ToString())
	if user == nil {
		return
	}

	h.mailService.Send(user.Email.ToString(), "Booking Reserved", "Your booking has been reserved")
}
