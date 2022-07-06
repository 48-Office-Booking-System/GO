package repository

import (
	"KOBA/domain"
	"KOBA/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//create booking
func (r *repoBooking) CreateBooking(booking model.Booking) (id int, err error) {
	res := r.DB.Debug().Omit("User", "Office").Create(&booking)
	if res.RowsAffected < 1 {
		return 0, res.Error
	}

	return booking.ID, nil
}

//Gett all booking
func (r *repoBooking) GetBookings() (bookings []model.Booking, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		Find(&bookings).
		Error; err != nil {

		return []model.Booking{}, err

	}

	return bookings, nil
}

//Get booking by id
func (r *repoBooking) GetBookingByID(id int) (booking model.Booking, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		First(&booking, id).
		Error; err != nil {

		return model.Booking{}, err

	}

	return booking, nil
}

//update booking
func (r *repoBooking) UpdateBooking(booking model.Booking, id int) error {
	booking.ID = id

	res := r.DB.Debug().Where("id = ?", id).Omit(clause.Associations).UpdateColumns(&booking)
	if res.RowsAffected < 1 {
		return res.Error
	}

	return nil
}

//delete booking
func (r *repoBooking) DeleteBooking(id int) error {
	booking := model.Booking{}
	booking.ID = id

	res := r.DB.Find(&booking)

	if res.RowsAffected < 1 {
		return fmt.Errorf("booking not found")
	}

	err := r.DB.Debug().Omit(clause.Associations).Unscoped().Delete(&booking).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repoBooking) GetStatuses() (statuses []model.Status, err error) {
	if err = r.DB.Debug().
		Find(&statuses).
		Error; err != nil {

		return []model.Status{}, err

	}

	return statuses, nil
}

func (r *repoBooking) DeleteStatus(id int) error {
	status := model.Status{}
	status.ID = id

	res := r.DB.Find(&status)

	if res.RowsAffected < 1 {
		return fmt.Errorf("status not found")
	}

	err := r.DB.Debug().Unscoped().Delete(&status).Error

	if err != nil {
		return err
	}

	return nil
}

func NewBooking(db *gorm.DB) domain.BookingRepoAdapter {
	return &repoBooking{
		DB: db,
	}
}
