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

func NewServiceRepo(repo domain.UserRepoAdapter, conf config.Config) domain.ServiceAdapter {
	return &userService{
		repo: repo,
		conf: conf,
	}
}
