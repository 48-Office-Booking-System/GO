package repository

import (
	"KOBA/domain"
	"KOBA/model"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const dateFormat = "2006-01-02"
const hourFormat = "15:04:05 WIB"

//create booking
func (r *repoBooking) CreateBooking(booking model.Booking) (id int, err error) {
	booking.StartDate, err = time.Parse(dateFormat, booking.StartDateString)
	if err != nil {
		return 0, err
	}

	booking.EndDate, err = time.Parse(hourFormat, booking.EndDateString)
	if err != nil {
		return 0, err
	}

	booking.StartHour, err = time.Parse(hourFormat, booking.StartHourString)
	if err != nil {
		return 0, err
	}

	booking.EndHour, err = time.Parse(hourFormat, booking.EndHourString)
	if err != nil {
		return 0, err
	}

	office := model.Office{
		ID: booking.OfficeID,
	}
	r.DB.First(&office, booking.OfficeID)

	booking.DurationInHour = int(math.Round(booking.EndHour.Sub(booking.StartHour).Hours()))
	booking.TotalPrice = office.Price * booking.DurationInHour

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
	var err error

	temp := model.Booking{
		ID: id,
	}

	r.DB.Preload("Office").First(&temp, booking.ID)

	if booking.StartDateString != "" {
		booking.StartDate, err = time.Parse(dateFormat, booking.StartDateString)
		if err != nil {
			return err
		}
		temp.StartDate = booking.StartDate
	}

	if booking.EndDateString != "" {
		booking.EndDate, err = time.Parse(hourFormat, booking.EndDateString)
		if err != nil {
			return err
		}
		temp.EndDate = booking.EndDate
	}

	if booking.StartHourString != "" {
		booking.StartHour, err = time.Parse(hourFormat, booking.StartHourString)
		if err != nil {
			return err
		}
		temp.StartHour = booking.StartHour
	}

	if booking.EndHourString != "" {
		booking.EndHour, err = time.Parse(hourFormat, booking.EndHourString)
		if err != nil {
			return err
		}
		temp.EndHour = booking.EndHour
	}

	booking.DurationInHour = int(math.Round(booking.EndHour.Sub(booking.StartHour).Hours()))
	booking.TotalPrice = temp.Office.Price * booking.DurationInHour

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
