package events

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type UserCreatedDomainEvent struct {
	UserId shared.UUID
}

func NewUserCreatedDomainEvent(userId shared.UUID) *UserCreatedDomainEvent {
	return &UserCreatedDomainEvent{userId}
}
