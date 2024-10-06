package users

import "github.com/mehedimayall/bookify-go/internal/domain/shared"

type User struct {
	shared.BaseEntity
	FirstName shared.Name  `json:"firstname"`
	LastName  shared.Name  `json:"lastname"`
	Email     shared.Email `json:"email"`
}
