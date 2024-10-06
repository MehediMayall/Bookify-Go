package apartments

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type Apartment struct {
	shared.BaseEntity
	name shared.Name
}
