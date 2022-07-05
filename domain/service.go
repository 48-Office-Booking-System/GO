package domain

import (
	"KOBA/model"
)

type UserServiceAdapter interface {
	CreateUserService(user model.User) (id int, err error)
	GetUsersService() (users []model.User, err error)
	GetUserByIDService(id int) (user model.User, err error)
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

type OfficeServiceAdapter interface {
	CreateOfficeService(office model.Office) (id int, err error)
	GetAllOfficeService() (office []model.Office, err error)
	GetOfficeService(id int) (office model.Office, err error)
	UpdateOfficeService(office model.Office, id int) error
	DeleteOfficeService(id int) error
}

type ReviewServiceAdapter interface {
	CreateReviewService(review model.Review) (id int, err error)
	GetReviewsService() (reviews []model.Review, err error)
	GetReviewByIDService(id int) (review model.Review, err error)
	UpdateReviewService(review model.Review, id int) error
	DeleteReviewService(id int) error
}
