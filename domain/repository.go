package domain

import (
	"KOBA/model"
)

type UserRepoAdapter interface {
	CreateUser(user model.User) (id int, err error)
	UpdateUser(user model.User, id int) error
	DeleteUser(id int) error
}

type BookingRepoAdapter interface {
	CreateBookingService(booking model.Booking) (id int, err error)
	GetBookings() (bookings []model.Booking, err error)
	GetBookingByID(id int) (booking model.Booking, err error)
	UpdateBooking(booking model.Booking, id int) error
	DeleteBooking(id int) error
}
