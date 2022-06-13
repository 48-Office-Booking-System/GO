package model

type User struct {
	ID       int `json:"id" gorm:"primaryKey"`
	RoleID   int
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TypeID    int
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

type Role struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type Type struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type Tags struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type Facillitations struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}

type Status struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
}
