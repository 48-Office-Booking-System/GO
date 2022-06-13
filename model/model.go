package model

type User struct {
	ID       int `json:"id" gorm:"primaryKey"`
	RoleID   int
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	Type          Type           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Location      string         `json:"location"`
	ViewCount     int            `json:"view_count"`
	Price         int            `json:"price"`
	KursiMin      int            `json:"kursi_min"`
	KursiMax      int            `json:"kursi_max"`
	Photourl      string         `json:"photourl"`
	Number        string         `json:"number"`
	Facilitations []Facilitation `gorm:"many2many:office_facilitations;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags          []Tag          `gorm:"many2many:office_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
