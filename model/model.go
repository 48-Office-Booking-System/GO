package model

type name struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	Photo    string `json:"photo" form:"photo"`
}
