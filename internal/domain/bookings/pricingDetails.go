package bookings

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type PricingDetails struct {
	priceForPeriod    shared.Money
	cleaningFee       shared.Money
	amenitiesUpCharge shared.Money
	totalPrice        shared.Money
}
