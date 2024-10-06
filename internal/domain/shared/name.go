package shared

import "errors"

type Name struct{ value string }

func NewName(value string) (*Name, error) {
	if value == "" {
		return nil, errors.New("invalid name. please provide a valid name")
	}
	return &Name{
		value: value,
	}, nil
}

func (n *Name) ToString() string {
	return n.value
}
