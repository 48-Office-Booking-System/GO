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

func (rs *userService) GetUsersService(user model.User) (users []model.User, err error) {
	return rs.repo.GetUsers(user)
}

func (rs *userService) GetUserByIDService(id int) (user model.User, err error) {
	return rs.repo.GetUserByID(id)
}

func (rs *userService) UpdateUserService(user model.User, id int) error {
	return rs.repo.UpdateUser(user, id)
}

func (rs *userService) DeleteUserService(id int) error {
	return rs.repo.DeleteUser(id)
}

func (rs *userService) GetRolesService() (roles []model.Role, err error) {
	return rs.repo.GetRoles()
}

func (rs *userService) DeleteRoleService(id int) error {
	return rs.repo.DeleteRole(id)
}

func NewUserService(repo domain.UserRepoAdapter, conf config.Config) domain.UserServiceAdapter {
	return &userService{
		repo: repo,
		conf: conf,
	}
}
