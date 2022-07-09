package domain

import (
	"KOBA/model"
)

type UserRepoAdapter interface {
	CreateUser(user model.User) (id int, err error)
	GetUsers(user model.User) (users []model.User, err error)
	GetUserByID(id int) (user model.User, err error)
	UpdateUser(user model.User, id int) error
	DeleteUser(id int) error

	GetRoles() (roles []model.Role, err error)
	DeleteRole(id int) error
}

type BookingRepoAdapter interface {
	CreateBooking(booking model.Booking) (id int, err error)
	GetBookings() (bookings []model.Booking, err error)
	GetBookingByID(id int) (booking model.Booking, err error)
	UpdateBooking(booking model.Booking, id int) error
	DeleteBooking(id int) error

	GetStatuses() (statuses []model.Status, err error)
	DeleteStatus(id int) error
}

type OfficeRepoAdapter interface {
	CreateOffice(office model.Office) (id int, err error)
	GetAllOffice() (office []model.Office, err error)
	GetOffice(id int) (office model.Office, err error)
	UpdateOffice(office model.Office, id int) error
	DeleteOffice(id int) error

	GetTypes() (types []model.Type, err error)
	DeleteType(id int) error

	GetFacilitations() (facilitations []model.Facilitation, err error)
	DeleteFacilitation(id int) error

	GetTags() (tags []model.Tag, err error)
	DeleteTag(id int) error
}

type ReviewRepoAdapter interface {
	CreateReview(review model.Review) (id int, err error)
	GetReviews() (reviews []model.Review, err error)
	GetReviewByID(id int) (review model.Review, err error)
	UpdateReview(review model.Review, id int) error
	DeleteReview(id int) error
}
