package model

import "time"

type User struct {
	ID       int `json:"id" gorm:"primaryKey"`
	RoleID   int `json:"role_id,omitempty"`
	Role     Role      `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	PhotoUrl string    `json:"photo_url"`
	Number   string    `json:"number"`
	Bookings []Booking `json:"bookings"`
	//Reviews  []Review  `json:"reviews"`
}

type Office struct {
	ID int `json:"id" gorm:"primaryKey"`
	// CreatedBy     int
	// User          User `json:"created_by" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TypeID        int `json:"type_id,omitempty"`
	Type          Type           `json:"type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Location      string         `json:"location"`
	ViewCount     int            `json:"view_count" gorm:"default:0"`
	Price         int            `json:"price"`
	ChairMin      int            `json:"chair_min"`
	ChairMax      int            `json:"chair_max"`
	PhotoUrls     []PhotoUrl     `json:"photo_urls"`
	Number        string         `json:"number"`
	Facilitations []Facilitation `json:"facilitations" gorm:"many2many:office_facilitations;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags          []Tag          `json:"tags" gorm:"many2many:office_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	//	Reviews       []Review       `json:"reviews"`
}

type Booking struct {
	ID              int `json:"id" gorm:"primaryKey"`
	UserID          int `json:"user_id,omitempty"`
	User            User `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OfficeID        int `json:"office_id,omitempty"`
	Office          Office `json:"office,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StatusID        int `json:"status_id,omitempty"`
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

type PhotoUrl struct {
	Url      string `json:"url"`
	OfficeID int `json:"office_id"`
}

/*
type Review struct {
	Text string `json:"text"`
	Star int    `json:"star"`
}
*/

type CreatedBy struct {
	OfficeID int
	UserID   int
}
