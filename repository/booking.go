package repository

import (
	"KOBA/domain"
	"KOBA/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type name struct {
}

//create booking
func (r *repoBooking) CreateBooking(booking model.Booking) (id int, err error) {
	res := r.DB.Debug().Save(&booking)
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

//Gett booking by id
func (r *repoBooking) GetBookingByID(id int) (booking model.Booking, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		First(&booking, id).
		Error; err != nil {

		return model.Booking{}, err

	}

	return booking, nil
}

func NewBooking(db *gorm.DB) domain.BookingRepoAdapter {
	return &repoBooking{
		DB: db,
	}
}
