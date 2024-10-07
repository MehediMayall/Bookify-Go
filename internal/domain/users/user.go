package users

import (
	"github.com/google/uuid"
	"github.com/mehedimayall/bookify-go/internal/domain/abstractions"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
)

type User struct {
	abstractions.EntityBase
	FirstName shared.Name
	LastName  shared.Name
	Email     Email
}

func NewUser(firstname, lastname, email string) (*User, error) {
	userid, err := shared.NewUUID(uuid.NewString())
	if err != nil {
		return nil, err
	}

	Firstname, err := shared.NewName(firstname)
	if err != nil {
		return nil, err
	}

	Lastname, err := shared.NewName(lastname)
	if err != nil {
		return nil, err
	}

	Email, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		EntityBase: abstractions.EntityBase{
			Id: *userid,
		},
		FirstName: *Firstname,
		LastName:  *Lastname,
		Email:     *Email,
	}, nil
}
