package model

import "time"

type User struct {
	ID       int `json:"id" gorm:"primaryKey"`
	RoleID   int
	Role     Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Photourl string `json:"photourl"`
	Number   string `json:"number"`
	//Booking
}

type Office struct {
	ID            int `json:"id" gorm:"primaryKey"`
	CreatedBy     int
	User          User `json:"created_by" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TypeID        int
	Type          Type           `json:"type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Location      string         `json:"location"`
	ViewCount     int            `json:"view_count"`
	Price         int            `json:"price"`
	KursiMin      int            `json:"kursi_min"`
	KursiMax      int            `json:"kursi_max"`
	Photourl      string         `json:"photourl"`
	Number        string         `json:"number"`
	Facilitations []Facilitation `json:"facilitations" gorm:"many2many:office_facilitations;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags          []Tag          `json:"tags" gorm:"many2many:office_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Booking struct {
	ID              int `json:"id" gorm:"primaryKey"`
	UserID          int
	User            User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OfficeID        int
	Office          Office `json:"office" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StatusID        int
	Status          Status    `json:"status" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	BuktiPembayaran string    `json:"bukti_pembayaran"`
}

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Type struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Tag struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Facilitation struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Status struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
