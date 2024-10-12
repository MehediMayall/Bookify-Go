package reservebooking

import (
	"time"

	"github.com/mehedimayall/bookify-go/internal/domain/apartments"
	"github.com/mehedimayall/bookify-go/internal/domain/bookings"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
	"github.com/mehedimayall/bookify-go/internal/domain/users"
)

type ReserveBookingCommand struct {
	bookingRepo    bookings.IBookingRepository
	userRepo       users.IUserRepository
	apartmentRepo  apartments.IApartmentRepository
	pricingService bookings.PricingService
}

func NewReserveBookingCommand(
	bookingRepo bookings.IBookingRepository,
	userRepo users.IUserRepository,
	apartmentRepo apartments.IApartmentRepository,
	pricingService bookings.PricingService) *ReserveBookingCommand {

	return &ReserveBookingCommand{
		bookingRepo,
		userRepo,
		apartmentRepo,
		pricingService,
	}
}

func (h *ReserveBookingCommand) Handle(
	userid, apartmentid, bookingid string,
	startDate, endDate time.Time) (string, error) {

	user := h.userRepo.GetById(userid)
	if user == nil {
		return "", users.NotFound(userid)
	}

	apartment := h.apartmentRepo.GetById(apartmentid)
	if apartment == nil {
		return "", apartments.NotFound(apartmentid)
	}

	duration, err := bookings.NewDateRange(startDate, endDate)
	if err == nil {
		return "", err
	}

	if h.bookingRepo.IsOverlaping(*apartment, *duration) {
		return "", bookings.Overlap()
	}

	userID := shared.NewUUID(userid)
	booking, err := bookings.Reserve(apartment, *userID, duration)

	if err == nil {
		return "", err
	}

	return booking.Id.ToString(), nil

}
