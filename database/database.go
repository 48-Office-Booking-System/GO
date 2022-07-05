package database

import (
	"KOBA/config"
	"KOBA/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString))
	if err != nil {
		fmt.Println("error opening connection : ", err)
	}

	err = DB.Debug().AutoMigrate(
		&model.User{},
		&model.Office{},
		&model.Booking{},
		&model.Tag{},
		&model.Facilitation{},
		&model.Role{},
		&model.Type{},
		&model.Status{},
		&model.PhotoUrl{},
		&model.Review{},
	)
	if err != nil {
		fmt.Print("error migrating table : ", err)
	} else {
		fmt.Println("BERHASILLLL")
	}

	return DB
}
