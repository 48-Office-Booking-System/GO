package domain

import (
	"KOBA/model"
)

type ServiceAdapter interface {
	CreateUserService(user model.User) (id int, err error)
}
