package shared

import "errors"

type Currency struct {
	value string
}

func NewCurrency(value string) (*Currency, error) {
	if value == "" {
		return nil, errors.New("invalid currency. please provide a valid currency")
	}
	return &Currency{
		value: value,
	}, nil
}

func (c *Currency) GetAll() []Currency {
	return []Currency{
		*c.USD(),
		*c.EUR(),
	}
}

func (n *Currency) USD() *Currency {
	return &Currency{value: "USD"}
}

func (n *Currency) EUR() *Currency {
	return &Currency{value: "EUR"}
}

func (n *Currency) ToString() string {
	return n.value
}
