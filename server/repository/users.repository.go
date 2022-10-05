package repository

import (
	Structs "github.com/guiluizmaia/tcc-message-golang/server/structs"
)

var users = []Structs.User{}

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (usersRepository UsersRepository) List() []Structs.User {
	return users
}

func (usersRepository UsersRepository) Create(user Structs.User) Structs.User {
	users = append(users, user)

	return user
}
