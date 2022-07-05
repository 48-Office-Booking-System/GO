package controller

import "KOBA/domain"

type UserServiceController struct {
	UserServ domain.UserServiceAdapter
}

type BookingServiceController struct {
	BookingServ domain.BookingServiceAdapter
}

type OfficeServiceController struct {
	OfficeServ domain.OfficeServiceAdapter
}

type ReviewServiceController struct {
	ReviewServ domain.ReviewServiceAdapter
}
