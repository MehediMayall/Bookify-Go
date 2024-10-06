package users

import "errors"

type Email struct{ value string }

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, errors.New("invalid email. please provide a valid email")
	}
	return &Email{
		value: value,
	}, nil
}

func (n *Email) ToString() string {
	return n.value
}
