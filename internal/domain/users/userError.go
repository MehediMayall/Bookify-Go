package users

import "errors"

func NotFound(id string) error {
	return errors.New("the user not found using userid: " + id)
}
