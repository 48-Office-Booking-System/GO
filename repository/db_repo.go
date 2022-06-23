package repository

import (
	"gorm.io/gorm"
)

type repoUser struct {
	DB *gorm.DB
}
