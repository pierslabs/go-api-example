package domain

type UserRepository interface {
	FindByID(id string) (*User, error)
	Create(user *User) (*User, error)
}
