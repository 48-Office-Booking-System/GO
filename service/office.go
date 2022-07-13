package service

import (
	"KOBA/config"
	"KOBA/domain"
	"KOBA/model"
)

type officeService struct {
	conf config.Config
	repo domain.OfficeRepoAdapter
}

func (rs *officeService) CreateOfficeService(office model.Office) (id int, err error) {
	return rs.repo.CreateOffice(office)
}

func (rs *officeService) GetAllOfficeService(id int) (allOffice []model.Office, err error) {
	return rs.repo.GetAllOffice()
}

func (rs *officeService) GetOfficeService(id int) (office model.Office, err error) {
	return rs.repo.GetOffice(id)
}

func (rs *officeService) UpdateOfficeService(office model.Office, id int) error {
	return rs.repo.UpdateOffice(office, id)
}

func (rs *officeService) DeleteOfficeService(id int) error {
	return rs.repo.DeleteOffice(id)
}

func NewOfficeService(repo domain.OfficeRepoAdapter, conf config.Config) domain.OfficeServiceAdapter {
	return &officeService{
		repo: repo,
		conf: conf,
	}
}
