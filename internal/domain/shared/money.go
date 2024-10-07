package shared

type Money struct {
	amount   float32
	currency Currency
}

func NewMoney(amount float32, currency Currency) (*Money, error) {

	return &Money{
		amount:   amount,
		currency: currency,
	}, nil
}

func ZeroMoney(currency Currency) *Money {
	return &Money{
		amount:   0,
		currency: currency,
	}
}

func (n *Money) GetValue() float32 {
	return n.amount
}

func (n *Money) GetCurrency() Currency {
	return n.currency
}
