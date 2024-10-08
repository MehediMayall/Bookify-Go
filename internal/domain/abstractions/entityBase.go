package abstractions

import (
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type EntityBase struct {
	Id shared.UUID `json:"id"`
}

func NewID() *EntityBase {
	return &EntityBase{
		Id: *shared.NewUUID(""),
	}
}
