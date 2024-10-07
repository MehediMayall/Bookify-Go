package shared

import (
	"github.com/google/uuid"
)

type UUID struct {
	value string
}

func NewUUID(value string) *UUID {
	if value == "" {
		value = uuid.NewString()
	}
	return &UUID{value: value}
}

func (n *UUID) ToString() string {
	return n.value
}
