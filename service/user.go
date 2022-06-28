package service

import (
	"KOBA/config"
	"KOBA/domain"
	"KOBA/model"
)

type userService struct {
	conf config.Config
	repo domain.UserRepoAdapter
}

func (rs *userService) CreateUserService(user model.User) (id int, err error) {
	return rs.repo.CreateUser(user)
}

func (rs *userService) UpdateUserService(user model.User, id int) error {
	return rs.repo.UpdateUser(user, id)
}

func (rs *userService) DeleteUserService(id int) error {
	return rs.repo.DeleteUser(id)
}

func NewUserService(repo domain.UserRepoAdapter, conf config.Config) domain.UserServiceAdapter {
	return &userService{
		repo: repo,
		conf: conf,
	}
}
