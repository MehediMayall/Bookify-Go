package apartments

import "errors"

func NotFound(id string) error {
	return errors.New("the booking is not found with booking id: " + id)
}
