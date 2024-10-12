package abstractions

import (
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type EntityBase struct {
	Id           shared.UUID `json:"id"`
	domainEvents []IDomainEvent
}

func NewID() *EntityBase {
	return &EntityBase{
		Id: *shared.NewUUID(""),
	}
}

func (e *EntityBase) AddDomainEvent(event IDomainEvent) {
	e.domainEvents = append(e.domainEvents, event)
}

func (e *EntityBase) ClearDomainEvents() {
	e.domainEvents = []IDomainEvent{}
}

func (e *EntityBase) GetDomainEvents() []IDomainEvent {
	return e.domainEvents
}
