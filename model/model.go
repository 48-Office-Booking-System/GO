package model

type Name struct {
	ID     int `json:"id" gorm:"primaryKey"`
	RoleID int
	//Role Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Photourl string `json:"photourl"`
	number   string `json:"number"`
	//Booking
}

type Office struct {
	ID        int `json:"id" gorm:"primaryKey"`
	CreatedBy int
	//CreatedBy User
	TypeID int
	//Role Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	ViewCount   int    `json:"view_count"`
	Price       int    `json:"price"`
	KursiMin    int    `json:"kursi_min"`
	KursiMax    int    `json:"kursi_max"`
	Photourl    string `json:"photourl"`
	Number      string `json:"number"`
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