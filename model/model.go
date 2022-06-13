package model

type name struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	Photo    string `json:"photo" form:"photo"`
}

type role struct {
	Id int `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type type struct {
	Id int `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type tags struct {
	Id int `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type facillitations struct {
	Id int `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type status struct {
	Id int `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}