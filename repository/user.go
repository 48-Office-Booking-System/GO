package repository

import (
	"KOBA/domain"
	"KOBA/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repoUser) CreateUser(user model.User) (id int, err error) {
	res := r.DB.Debug().Omit(clause.Associations).Create(&user)
	if res.RowsAffected < 1 {
		return 0, res.Error
	}

	return user.ID, nil
}

func (r *repoUser) GetUsers(user model.User) (users []model.User, err error) {
	err = r.DB.Debug().
		Where(&user).
		Preload(clause.Associations).
		Find(&users).
		Error

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (r *repoUser) GetUserByID(id int) (user model.User, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		First(&user, id).
		Error; err != nil {

		return model.User{}, err

	}

	return user, nil
}

func (r *repoUser) UpdateUser(user model.User, id int) error {
	res := r.DB.Debug().Where("id = ?", id).Omit(clause.Associations).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return res.Error
	}

	return nil
}

func (r *repoUser) DeleteUser(id int) error {
	user := model.User{}
	user.ID = id

	res := r.DB.Find(&user)

	if res.RowsAffected < 1 {
		return fmt.Errorf("user not found")
	}

	err := r.DB.Debug().Omit("Role").Unscoped().Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repoUser) GetRoles() (roles []model.Role, err error) {
	if err = r.DB.Debug().
		Find(&roles).
		Error; err != nil {

		return []model.Role{}, err

	}

	return roles, nil
}

func (r *repoUser) DeleteRole(id int) error {
	role := model.Role{}
	role.ID = id

	res := r.DB.Find(&role)

	if res.RowsAffected < 1 {
		return fmt.Errorf("role not found")
	}

	err := r.DB.Debug().Unscoped().Delete(&role).Error

	if err != nil {
		return err
	}

	return nil
}

func NewUser(db *gorm.DB) domain.UserRepoAdapter {
	return &repoUser{
		DB: db,
	}
}
