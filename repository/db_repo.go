package repository

import (
	"gorm.io/gorm"
)

type repoUser struct {
	DB *gorm.DB
}

type repoBooking struct {
	DB *gorm.DB
}
