package apartments

import "errors"

type Description struct {
	value string
}

func NewDescription(value string) (*Description, error) {
	if value == "" {
		return nil, errors.New("invalid description. please provide a valid description")
	}
	return &Description{
		value: value,
	}, nil
}

func (n *Description) ToString() string {
	return n.value
}
