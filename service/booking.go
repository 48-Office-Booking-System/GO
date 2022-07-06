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

func (rs *bookingService) UpdateBookingService(booking model.Booking, id int) error {
	return rs.repo.UpdateBooking(booking, id)
}

func (rs *bookingService) DeleteBookingService(id int) error {
	return rs.repo.DeleteBooking(id)
}

func (rs *bookingService) GetStatusesService() (statuses []model.Status, err error) {
	return rs.repo.GetStatuses()
}

func (rs *bookingService) DeleteStatusService(id int) error {
	return rs.repo.DeleteStatus(id)
}

func NewBookingService(repo domain.BookingRepoAdapter, conf config.Config) domain.BookingServiceAdapter {
	return &bookingService{
		repo: repo,
		conf: conf,
	}
}
