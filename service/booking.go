package service

import (
	"KOBA/config"
	"KOBA/domain"
	"KOBA/model"
)

type bookingService struct {
	conf config.Config
	repo domain.BookingRepoAdapter
}

func (rs *bookingService) CreateBookingService(booking model.Booking) (id int, err error) {
	return rs.repo.CreateBooking(booking)
}

func (rs *bookingService) GetBookingsService() (bookings []model.Booking, err error) {
	return rs.repo.GetBookings()
}

func (rs *bookingService) GetBookingByIDService(id int) (booking model.Booking, err error) {
	return rs.repo.GetBookingByID(id)
}

func NewBookingService(repo domain.BookingRepoAdapter, conf config.Config) domain.BookingServiceAdapter {
	return &bookingService{
		repo: repo,
		conf: conf,
	}
}
