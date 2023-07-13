package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/reviewimage"
	"errors"

	"gorm.io/gorm"
)

type ImagesReviewData struct {
	db *gorm.DB
}

// Insert implements reviewimage.ReviewImageData.
func (data *ImagesReviewData) Insert(imagesreview features.ReviewImagesEntity, reviewId uint) (uint, error) {
	var ImagesModel features.Review
	tx:=data.db.First(&ImagesModel,reviewId)
	if tx.Error != nil{
		return 0,errors.New("id tidak ditemukan")
	}
	ImagesReview:= features.ReviewImageEntityToModel(imagesreview)

	ImagesReview.ReviewID = reviewId
	txx:= data.db.Create(&ImagesReview)
	if txx.Error != nil{
		return 0,txx.Error
	}
	return ImagesReview.ID,nil

}

func New(data *gorm.DB) reviewimage.ReviewImageData {
	return &ImagesReviewData{
		db: data,
	}
}
