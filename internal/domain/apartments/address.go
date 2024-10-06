package apartments

import (
	"errors"
	"fmt"
)

type Address struct {
	country string `json:"country"`
	city    string `json:"city"`
	zipCode string `json:"zipcode"`
	state   string `json:"state"`
	street  string `json:"street"`
}

func NewAddress(country, city, zipCode, state, street string) (*Address, error) {
	if country == "" {
		return nil, errors.New("invalid country. please provide a valid country")
	}
	return &Address{
		country: country,
		city:    city,
		zipCode: zipCode,
		state:   state,
		street:  street,
	}, nil
}

func (n *Address) ToString() string {
	return fmt.Sprintf("%v", n)
}
