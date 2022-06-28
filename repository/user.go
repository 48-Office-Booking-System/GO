package repository

import (
	"KOBA/domain"
	"KOBA/model"
	"fmt"

	"gorm.io/gorm"
)

func (r *repoUser) CreateUser(user model.User) (id int, err error) {
	res := r.DB.Debug().Save(&user)
	if res.RowsAffected < 1 {
		return 0, fmt.Errorf("error creating user")
	}

	return user.ID, nil
}

func (r *repoUser) UpdateUser(user model.User, id int) error {
	user.ID = id

	res := r.DB.Debug().Save(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error creating user")
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

	err := r.DB.Debug().Unscoped().Delete(&user).Error

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
