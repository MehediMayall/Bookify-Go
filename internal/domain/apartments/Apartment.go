package apartments

import (
	"time"

	"github.com/mehedimayall/bookify-go/internal/domain/abstractions"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type Apartment struct {
	abstractions.EntityBase
	name            shared.Name
	description     Description
	address         Address
	price           shared.Money
	cleaningFee     shared.Money
	amenities       []Amenity
	lastBookedOnUtc time.Time
}

func NewApartment(
	name, description, country, city, zipCode, state, street string,
	priceAmount float32, priceCurrency string) (*Apartment, error) {

	return &Apartment{}, nil
}
