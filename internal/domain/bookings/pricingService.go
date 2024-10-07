package bookings

import (
	"github.com/mehedimayall/bookify-go/internal/domain/apartments"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type PricingService struct {
}

func (p *PricingService) CalculatePrice(apartment apartments.Apartment, duration DateRange) (*PricingDetails, error) {
	// Get the currency from the apartment price
	price := apartment.Price()
	currency := price.GetCurrency()
	clearningFee := apartment.CleaningFee()

	// Calculate PriceForPeriod
	priceForPeriod, err := shared.NewMoney(float32(duration.daysInNumber)*price.GetValue(), currency)
	if err != nil {
		return nil, err
	}

	// Calculate amenities charge percent
	amenitiesChargePercent := 0.0
	for amenity := range apartment.Amenities() {
		switch amenity {
		case apartments.Amenities.GardenView, apartments.Amenities.MountainView:
			amenitiesChargePercent += 0.05
		case apartments.Amenities.AirConditioning:
			amenitiesChargePercent += 0.01
		case apartments.Amenities.Parking:
			amenitiesChargePercent += 0.01
		default:
			amenitiesChargePercent += 0.0

		}
	}

	// Calculate amenitiesUpCharge
	amenitiesUpCharge := shared.ZeroMoney(currency)

	if amenitiesChargePercent > 0 {
		amenitiesUpCharge, err = shared.NewMoney(priceForPeriod.GetValue()*float32(amenitiesChargePercent), currency)
	}

	totalPrice, err := shared.NewMoney(priceForPeriod.GetValue(), currency)
	totalPrice.Add(amenitiesUpCharge.GetValue())
	totalPrice.Add(clearningFee.GetValue())

	return &PricingDetails{
		priceForPeriod:    *priceForPeriod,
		cleaningFee:       clearningFee,
		amenitiesUpCharge: *amenitiesUpCharge,
		totalPrice:        *totalPrice,
	}, nil

}
