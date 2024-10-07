package apartments

import (
	"time"

	"github.com/mehedimayall/bookify-go/internal/domain/abstractions"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type Apartment struct {
	abstractions.EntityBase
	Name            shared.Name
	description     Description
	address         Address
	price           shared.Money
	cleaningFee     shared.Money
	amenities       []Amenity
	lastBookedOnUtc time.Time
}
