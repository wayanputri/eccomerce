package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/review"
	"errors"

	"gorm.io/gorm"
)

type Review struct {
	db *gorm.DB
}

// Insert implements review.ReviewData.
func (repo Review) Insert(review features.ReviewEntity, product_id uint) (uint, error) {
	
	var product features.Product
	txx:= repo.db.First(&product,product_id)
	if txx.Error != nil{
		return 0,txx.Error
	}

	reviewModel := features.ReviewEntityToModel(review)
	reviewModel.ProductID = product_id
	if reviewModel.Rating > 5 {
		return 0,errors.New("nilai rating maksimal 5")
	}
	tx:=repo.db.Create(&reviewModel)
	if tx.Error != nil{
		return 0,tx.Error
	}
	
	return reviewModel.ID,nil
}

func (repo Review) UpdateRating(Rating int,product_id uint) (int,error){
	var product features.Product
	product.Ratings = Rating
	tx:=repo.db.Model(&product).Where("id=?",product_id).Update("ratings",product.Ratings)
	if tx.Error != nil{
		return 0,tx.Error
	}
	return product.Ratings,nil
}

func New(db *gorm.DB) review.ReviewData {
	return Review{
		db: db,
	}
}
