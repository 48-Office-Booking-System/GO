package service

import (
	"KOBA/config"
	"KOBA/domain"
	"KOBA/model"
)

type reviewService struct {
	conf config.Config
	repo domain.ReviewRepoAdapter
}

func (rs *reviewService) CreateReviewService(review model.Review) (id int, err error) {
	return rs.repo.CreateReview(review)
}

func (rs *reviewService) GetReviewsService(review model.Review) (reviews []model.Review, err error) {
	return rs.repo.GetReviews(review)
}

func (rs *reviewService) GetReviewByIDService(id int) (review model.Review, err error) {
	return rs.repo.GetReviewByID(id)
}

func (rs *reviewService) UpdateReviewService(review model.Review, id int) error {
	return rs.repo.UpdateReview(review, id)
}

func (rs *reviewService) DeleteReviewService(id int) error {
	return rs.repo.DeleteReview(id)
}

func NewReviewService(repo domain.ReviewRepoAdapter, conf config.Config) domain.ReviewServiceAdapter {
	return &reviewService{
		repo: repo,
		conf: conf,
	}
}
