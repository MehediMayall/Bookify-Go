package main

import (
	"fmt"

	"github.com/mehedimayall/bookify-go/internal/domain/apartments"
	"github.com/mehedimayall/bookify-go/internal/domain/shared"
	"github.com/mehedimayall/bookify-go/internal/domain/users"
)

func print[T any](values T) {
	fmt.Println(values)
}

func main() {

	user, err := users.NewUser("Mehedi", "Hasan", "mehedi@gmail.com")

	if err != nil {
		print(err)
		return
	}

	print(user.FirstName.ToString())
	print(user)

	currency := &shared.Currency{}

	print(currency.GetAll())
	print(apartments.Amenities.MountainView)

}
