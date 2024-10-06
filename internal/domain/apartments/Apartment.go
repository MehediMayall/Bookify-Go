package apartments

import (
	"github.com/mehedimayall/bookify-go/internal/domain/abstractions"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type Apartment struct {
	abstractions.EntityBase
	name shared.Name
}
