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

func NewRecipe(db *gorm.DB) domain.UserRepoAdapter {
	return &repoUser{
		DB: db,
	}
}
