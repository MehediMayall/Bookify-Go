package bookings

import (
	"time"

	"github.com/mehedimayall/bookify-go/internal/domain/abstractions"
	"github.com/mehedimayall/bookify-go/internal/domain/apartments"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type Booking struct {
	abstractions.EntityBase
	aparmentId shared.UUID
	userId     shared.UUID

	duration DateRange

	priceForPeriod    shared.Money
	cleaningFee       shared.Money
	amenitiesUpCharge shared.Money
	totalPrice        shared.Money

	status int

	createdOnUtc   time.Time
	confirmedOnUtc time.Time
	rejectedOnUtc  time.Time
	completedOnUtc time.Time
	cancelledOnUtc time.Time
}

func Reserve(apartment *apartments.Apartment, userId shared.UUID,
	duration *DateRange) (*Booking, error) {
	pricingService := PricingService{}
	pricingDetails, err := pricingService.CalculatePrice(*apartment, *duration)

	if err != nil {
		return nil, err
	}

	return &Booking{
		EntityBase: *abstractions.NewID(),
		aparmentId: apartment.Id,
		userId:     userId,
		duration:   *duration,

		priceForPeriod:    pricingDetails.priceForPeriod,
		cleaningFee:       pricingDetails.cleaningFee,
		amenitiesUpCharge: pricingDetails.amenitiesUpCharge,
		totalPrice:        pricingDetails.totalPrice,

		status:       BookingStatuses.reserved,
		createdOnUtc: time.Now().UTC(),
	}, nil
}

func (b *Booking) Confirm() error {
	if b.status != BookingStatuses.reserved {
		return NotReserved()
	}

	b.status = BookingStatuses.confirmed
	b.confirmedOnUtc = time.Now()

	return nil
}

func (b *Booking) Reject() error {
	if b.status != BookingStatuses.reserved {
		return NotReserved()
	}

	b.status = BookingStatuses.rejected
	b.rejectedOnUtc = time.Now()

	return nil
}

func (b *Booking) Cancel() error {
	if b.status != BookingStatuses.reserved {
		return NotReserved()
	}

	b.status = BookingStatuses.cancelled
	b.cancelledOnUtc = time.Now()

	return nil
}

func (b *Booking) Complete() error {
	if b.status != BookingStatuses.reserved {
		return NotReserved()
	}

	b.status = BookingStatuses.completed
	b.completedOnUtc = time.Now()

	return nil
}
