package repository

import (
	"KOBA/domain"
	"KOBA/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//create booking
func (r *repoReview) CreateReview(review model.Review) (id int, err error) {
	res := r.DB.Debug().Omit(clause.Associations).Create(&review)
	if res.RowsAffected < 1 {
		return 0, res.Error
	}

	return review.ID, nil
}

//Gett all booking
func (r *repoReview) GetReviews() (reviews []model.Review, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		Find(&reviews).
		Error; err != nil {

		return []model.Review{}, err

	}

	return reviews, nil
}

//Get booking by id
func (r *repoReview) GetReviewByID(id int) (review model.Review, err error) {
	if err = r.DB.Debug().
		Preload(clause.Associations).
		First(&review, id).
		Error; err != nil {

		return model.Review{}, err

	}

	return review, nil
}

//update booking
func (r *repoReview) UpdateReview(review model.Review, id int) error {
	review.ID = id

	res := r.DB.Debug().Where("id = ?", id).Omit(clause.Associations).UpdateColumns(&review)
	if res.RowsAffected < 1 {
		return res.Error
	}

	return nil
}

//delete booking
func (r *repoReview) DeleteReview(id int) error {
	review := model.Review{}
	review.ID = id

	res := r.DB.Find(&review)

	if res.RowsAffected < 1 {
		return res.Error
	}

	err := r.DB.Debug().Omit(clause.Associations).Unscoped().Delete(&review).Error

	if err != nil {
		return err
	}

	return nil
}

func NewReview(db *gorm.DB) domain.ReviewRepoAdapter {
	return &repoReview{
		DB: db,
	}
}
