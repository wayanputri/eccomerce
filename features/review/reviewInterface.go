package review

import (
	"belajar/bareng/features"
)

type ReviewData interface {
	Insert(review features.ReviewEntity,product_id uint) (uint,error)
	UpdateRating(Rating float64,product_id uint) (float64,error)
	Delete(review_Id uint) error
	UpdateRatingDelete(reviewID uint, productID uint) error
	SelectAll()([]features.ReviewEntity,error)
}

type ReviewService interface{
	Add(review features.ReviewEntity,product_id uint) (uint,error)
	Delete(review_Id uint) error
	SelectAll()([]features.ReviewEntity,error)
}