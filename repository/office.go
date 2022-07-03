package repository

import (
	"KOBA/domain"
	"KOBA/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// tambah
func (r *repoOffice) CreateOffice(office model.Office) (id int, err error) {
	office.ViewCount = 0
	res := r.DB.Debug().Omit("ViewCount").Create(&office)
	if res.RowsAffected < 1 {
		return 0, fmt.Errorf("error creating office")
	}

	return office.ID, nil
}

// lihat semua gedung
func (r *repoOffice) GetAllOffice() (office []model.Office, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		Find(&office).
		Error; err != nil {

		return []model.Office{}, err

	}

	return office, nil
}

// lihat gedung tertentu berdasarkan id
func (r *repoOffice) GetOffice(id int) (office model.Office, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		First(&office, id).
		Error; err != nil {

		return model.Office{}, err

	}

	return office, nil
}

// update gedung berdasarkan id
func (r *repoOffice) UpdateOffice(office model.Office, id int) error {
	res := r.DB.Debug().Where("id = ?", id).Omit(clause.Associations, "ViewCount").UpdateColumns(&office)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error updating office")
	}

	return nil
}

// hapus gedung berdasarkan id
func (r *repoOffice) DeleteOffice(id int) error {
	office := model.Office{}
	office.ID = id

	res := r.DB.Find(&office)

	if res.RowsAffected < 1 {
		return fmt.Errorf("office not found")
	}

	err := r.DB.Debug().Select("PhotoUrls", "Facilitations", "Tags").Delete(&office).Error

	if err != nil {
		return err
	}

	err = r.DB.Debug().Omit("Type").Unscoped().Delete(&office).Error

	if err != nil {
		return err
	}

	return nil
}

func NewOffice(db *gorm.DB) domain.OfficeRepoAdapter {
	return &repoOffice{
		DB: db,
	}
}
