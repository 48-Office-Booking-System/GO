package domain

import (
	"KOBA/model"
)

type UserServiceAdapter interface {
	CreateUserService(user model.User) (id int, err error)
	UpdateUserService(user model.User, id int) error
	DeleteUserService(id int) error
}

type BookingServiceAdapter interface {
	CreateBookingService(booking model.Booking) (id int, err error)
	GetBookingsService() (bookings []model.Booking, err error)
	GetBookingByIDService(id int) (booking model.Booking, err error)
	UpdateBookingService(booking model.Booking, id int) error
	DeleteBookingService(id int) error
}
