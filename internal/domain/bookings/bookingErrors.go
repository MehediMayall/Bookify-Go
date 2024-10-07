package bookings

import (
	"errors"
)

func NotFound() error {
	return errors.New("the booking with the specified identified was not found")
}

func Overlap() error {
	return errors.New("the current booking is overlapping with an existing one")
}

func NotReserved() error {
	return errors.New("the booking is not pending")
}
