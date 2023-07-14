package reviewimage

import "belajar/bareng/features"

type ReviewImageData interface {
	Insert(imagesreview features.ReviewImagesEntity, reviewId uint ) (uint, error)
	Delete(imageID uint)error
	SelectAll()([]features.ReviewImagesEntity,error)
}
type ReviewImageService interface {
	Add(imagesreview features.ReviewImagesEntity,reviewId uint) (uint, error)
	Delete(imageID uint)error
	GetAll()([]features.ReviewImagesEntity,error)
}