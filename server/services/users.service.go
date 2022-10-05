package services

import (
	"encoding/json"

	Repository "github.com/guiluizmaia/tcc-message-golang/server/repository"
	Structs "github.com/guiluizmaia/tcc-message-golang/server/structs"
)

type UserService struct {
	repository Repository.UsersRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) CreateUser(userByte []byte) (Structs.User, error) {
	var user Structs.User
	json.Unmarshal(userByte, &user)

	return userService.repository.Create(user), nil
}
