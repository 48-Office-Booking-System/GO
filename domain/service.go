package domain

import (
	"KOBA/model"
)

type UserServiceAdapter interface {
	CreateUserService(user model.User) (id int, err error)
}

type BookingServiceAdapter interface {
	CreateBookingService(booking model.Booking) (id int, err error)
	GetBookingsService() (bookings []model.Booking, err error)
	GetBookingByIDService(id int) (booking model.Booking, err error)
}
