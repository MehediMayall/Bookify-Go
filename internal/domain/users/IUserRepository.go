package users

type IUserRepository interface {
	GetById(Id string) *User
	Add(user *User)
}
