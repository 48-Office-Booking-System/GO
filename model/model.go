package model

import "time"

type User struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	RoleID   int       `json:"role_id,omitempty"`
	Role     Role      `json:"role" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	PhotoUrl string    `json:"photo_url"`
	Number   string    `json:"number"`
	Bookings []Booking `json:"bookings" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	Reviews  []Review  `json:"reviews" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}

type Office struct {
	ID int `json:"id" gorm:"primaryKey"`
	// CreatedBy     int
	// User          User `json:"created_by" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	TypeID        int            `json:"type_id,omitempty"`
	Type          Type           `json:"type" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Latitude      string         `json:"latitude"`
	Longitude     string         `json:"longitude"`
	ViewCount     int            `json:"view_count" gorm:"default:0"`
	Price         int            `json:"price"`
	ChairMin      int            `json:"chair_min"`
	ChairMax      int            `json:"chair_max"`
	PhotoUrls     []PhotoUrl     `json:"photo_urls"`
	Number        string         `json:"number"`
	Facilitations []Facilitation `json:"facilitations" gorm:"many2many:office_facilitations;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Tags          []Tag          `json:"tags" gorm:"many2many:office_tags;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Reviews       []Review       `json:"reviews" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}

type Booking struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	UserID          int       `json:"user_id,omitempty"`
	User            User      `json:"user,omitempty" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	OfficeID        int       `json:"office_id,omitempty"`
	Office          Office    `json:"office,omitempty" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	StatusID        int       `json:"status_id,omitempty"`
	Status          Status    `json:"status" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	StartDateString string    `json:"start_date" gorm:"-"`
	EndDateString   string    `json:"end_date" gorm:"-"`
	StartDate       time.Time `json:"-"`
	EndDate         time.Time `json:"-"`
	StartHourString string    `json:"start_hour" gorm:"-"`
	EndHourString   string    `json:"end_hour" gorm:"-"`
	StartHour       time.Time `json:"-"`
	EndHour         time.Time `json:"-"`
	DurationInHour  int       `json:"duration_in_hour"`
	TotalPrice      int       `json:"total_price"`
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
	OfficeID int    `json:"office_id"`
}

type Review struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	UserID   int    `json:"user_id,omitempty"`
	User     User   `json:"user,omitempty" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	OfficeID int    `json:"office_id,omitempty"`
	Office   Office `json:"office,omitempty" gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	Star     int    `json:"star"`
	Text     string `json:"text"`
	Hidden   int    `json:"hidden" gorm:"default:0"`
}

/*
type CreatedBy struct {
	OfficeID int
	UserID   int
}
*/
