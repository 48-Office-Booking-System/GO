package domain

import (
	"KOBA/model"
)

type UserRepoAdapter interface {
	CreateUser(user model.User) (id int, err error)
}

type BookingRepoAdapter interface {
	CreateBooking(booking model.Booking) (id int, err error)
	GetBookings() (bookings []model.Booking, err error)
	GetBookingByID(id int) (booking model.Booking, err error)
	UpdateBooking(booking model.Booking, id int) error
	DeleteBooking(id int) error
}

type OfficeRepoAdapter interface {
	CreateOffice(office model.Office) (id int, err error)
	GetAllOffice() (office []model.Office, err error)
	GetOfficeByID(id int) (office model.Office, err error)
	UpdateOffice(office model.Office, id int) error
	DeleteOffice(id int) error
}