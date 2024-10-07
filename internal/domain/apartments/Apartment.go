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
	name shared.Name,
	description Description,
	address Address,
	price shared.Money,
	cleaningFee shared.Money,
	amenities []Amenity) (*Apartment, error) {

	return &Apartment{
		name:        name,
		description: description,
		address:     address,
		price:       price,
		cleaningFee: cleaningFee,
		amenities:   amenities,
	}, nil
}

func (a *Apartment) ApartmentId() shared.UUID {
	return a.Id
}
func (a *Apartment) Price() shared.Money {
	return a.price
}

func (a *Apartment) CleaningFee() shared.Money {
	return a.cleaningFee
}

func (a *Apartment) Amenities() []Amenity {
	return a.amenities
}
