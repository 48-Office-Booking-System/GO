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
	/* booking.UserID = booking.User.ID
	booking.OfficeID = booking.Office.ID
	booking.StatusID = booking.Status.ID
 */
	res := r.DB.Debug().Omit(clause.Associations, "ID").Save(&booking)
	if res.RowsAffected < 1 {
		return 0, fmt.Errorf("error creating booking")
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
	/* booking.UserID = booking.User.ID
	booking.OfficeID = booking.Office.ID
	booking.StatusID = booking.Status.ID */

	res := r.DB.Debug().Omit(clause.Associations).Save(&booking)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error creating booking")
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

func NewBooking(db *gorm.DB) domain.BookingRepoAdapter {
	return &repoBooking{
		DB: db,
	}
}
