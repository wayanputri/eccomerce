package review

import "belajar/bareng/features"

type ReviewData interface {
	Insert(review features.ReviewEntity,product_id uint) (uint,error)
	UpdateRating(Rating int,product_id uint) (int,error)
}

type ReviewService interface{
	Add(review features.ReviewEntity,product_id uint) (uint,error)
}