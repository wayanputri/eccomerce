package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/review"
)

type ReviewData struct {
	reviewData review.ReviewData
}

// Delete implements review.ReviewService.
func (service ReviewData) Delete(review_Id uint) error {
	err := service.reviewData.Delete(review_Id)
	if err != nil{
		return err
	}
	return nil
}

// Add implements review.ReviewService.
func (service ReviewData) Add(review features.ReviewEntity, product_id uint) (uint, error) {
	id, err := service.reviewData.Insert(review, product_id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(service review.ReviewData) review.ReviewService {
	return ReviewData{
		reviewData: service,
	}
}
