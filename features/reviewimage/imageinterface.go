package reviewimage

import "belajar/bareng/features"

type ReviewImageData interface {
	Insert(imagesreview features.ReviewImagesEntity, reviewId uint ) (uint, error)
}
type ReviewImageService interface {
	Add(imagesreview features.ReviewImagesEntity,reviewId uint) (uint, error)
}