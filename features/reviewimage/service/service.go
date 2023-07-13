package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/reviewimage"
)

type ImagesReviewService struct {
	imagesReviewService reviewimage.ReviewImageData
}

// Delete implements reviewimage.ReviewImageService.
func (service ImagesReviewService) Delete(imageID uint) error {
	err:=service.imagesReviewService.Delete(imageID)
	if err != nil{
		return err
	}
	return nil
}

// Add implements reviewimage.ReviewImageService.
func (service ImagesReviewService) Add(imagesreview features.ReviewImagesEntity, reviewId uint) (uint, error) {
	id, err := service.imagesReviewService.Insert(imagesreview, reviewId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(service reviewimage.ReviewImageData) reviewimage.ReviewImageService {
	return ImagesReviewService{
		imagesReviewService: service,
	}
}
