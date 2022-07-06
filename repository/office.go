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
	res := r.DB.Debug().Create(&office)
	if res.RowsAffected < 1 {
		return 0, res.Error
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
		return res.Error
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

func (r *repoOffice) GetTypes() (types []model.Type, err error) {
	if err = r.DB.Debug().
		Find(&types).
		Error; err != nil {

		return []model.Type{}, err

	}

	return types, nil
}

func (r *repoOffice) DeleteType(id int) error {
	tipe := model.Type{}
	tipe.ID = id

	res := r.DB.Find(&tipe)

	if res.RowsAffected < 1 {
		return fmt.Errorf("type not found")
	}

	err := r.DB.Debug().Unscoped().Delete(&tipe).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repoOffice) GetFacilitations() (facilitations []model.Facilitation, err error) {
	if err = r.DB.Debug().
		Find(&facilitations).
		Error; err != nil {

		return []model.Facilitation{}, err

	}

	return facilitations, nil
}

func (r *repoOffice) DeleteFacilitation(id int) error {
	facilitation := model.Facilitation{}
	facilitation.ID = id

	res := r.DB.Find(&facilitation)

	if res.RowsAffected < 1 {
		return fmt.Errorf("facilitation not found")
	}

	err := r.DB.Debug().Unscoped().Delete(&facilitation).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repoOffice) GetTags() (tags []model.Tag, err error) {
	if err = r.DB.Debug().
		Find(&tags).
		Error; err != nil {

		return []model.Tag{}, err

	}

	return tags, nil
}

func (r *repoOffice) DeleteTag(id int) error {
	tag := model.Tag{}
	tag.ID = id

	res := r.DB.Find(&tag)

	if res.RowsAffected < 1 {
		return fmt.Errorf("tag not found")
	}

	err := r.DB.Debug().Unscoped().Delete(&tag).Error

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
