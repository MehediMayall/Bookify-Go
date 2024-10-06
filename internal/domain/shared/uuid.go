package shared

import "errors"

type UUID struct {
	value string
}

func NewUUID(value string) (*UUID, error) {
	if value == "" {
		return nil, errors.New("invalid uuid. please provide a valid uuid")
	}
	return &UUID{
		value: value,
	}, nil
}

func (n *UUID) ToString() string {
	return n.value
}
