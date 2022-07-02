package repository

import (
	"KOBA/domain"
	"KOBA/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repoUser) CreateUser(user model.User) (id int, err error) {
	res := r.DB.Debug().Omit("Booking").Create(&user)
	if res.RowsAffected < 1 {
		return 0, fmt.Errorf("error creating user")
	}

	return user.ID, nil
}

func (r *repoUser) GetUsers() (users []model.User, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		Find(&users).
		Error; err != nil {

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
	res := r.DB.Debug().Where("id = ?", id).Omit("Bookings").UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error updating user")
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

	err := r.DB.Debug().Omit(clause.Associations).Unscoped().Delete(&user).Error

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
