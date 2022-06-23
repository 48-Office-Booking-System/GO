package domain

import (
	"KOBA/model"
)

type UserRepoAdapter interface {
	CreateUser(user model.User) (id int, err error)
}
