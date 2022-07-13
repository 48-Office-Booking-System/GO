package domain

import (
	"KOBA/model"
)

type UserRepoAdapter interface {
	CreateUser(user model.User) (id int, err error)
	GetUsers() (users []model.User, err error)
	GetUserByID(id int) (user model.User, err error)
	UpdateUser(user model.User, id int) error
	DeleteUser(id int) error
}

type BookingRepoAdapter interface {
	CreateBooking(booking model.Booking) (id int, err error)
	GetBookings() (bookings []model.Booking, err error)
	GetBookingByID(id int) (booking model.Booking, err error)
	UpdateBooking(booking model.Booking, id int) error
	DeleteBooking(id int) error
	LoginBooking(booking model.Booking) (id int, err error)
}

type OfficeRepoAdapter interface {
	CreateOffice(office model.Office) (id int, err error)
	GetAllOffice() (office []model.Office, err error)
	GetOffice(id int) (office model.Office, err error)
	UpdateOffice(office model.Office, id int) error
	DeleteOffice(id int) error
}

type AdminRepoAdapter interface {
	CreateProduct(admin model.Admin) (id int, err error)
	GetBookings() (admins []model.Admin, err error)
	GetBookingByID(id int) (booking model.Admin, err error)
	UpdateBooking(booking model.Admin, id int) error
	DeleteBooking(id int) error
}
