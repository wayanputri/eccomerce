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

// UpdateRatingDelete implements review.ReviewData.
func (repo Review) UpdateRatingDelete(reviewID uint,productID uint) error{
	review := []features.Review{}
	FoundProductId := repo.db.Find(&review, "product_id=?", productID)
	if FoundProductId.Error != nil {
		return errors.New(FoundProductId.Error.Error() + ", failed to get review id")
	}

	var review1 features.Review
	txx := repo.db.First(&review1, "product_id = ? AND id = ?",productID,reviewID)
	if txx.Error != nil {
		return txx.Error
	}
	average,err:=AverageRatingsDelete(review,review1.Rating)
	if err != nil{
		return err
	}
	var product features.Product
	tx := repo.db.Model(&product).Where("id=?", productID).Update("ratings", average)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements review.ReviewData.
func (repo Review) Delete(review_Id uint) error {
	
	var review features.Review
	txx := repo.db.First(&review, review_Id)
	if txx.Error != nil {
		return txx.Error
	}
	err:=repo.UpdateRatingDelete(review.ID,review.ProductID)
	if err != nil{
		return err
	}

	tx := repo.db.Delete(&review, review_Id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Insert implements review.ReviewData.
func (repo Review) Insert(review features.ReviewEntity, product_id uint) (uint, error) {

	var product features.Product
	txx := repo.db.First(&product, product_id)
	if txx.Error != nil {
		return 0, txx.Error
	}

	reviewModel := features.ReviewEntityToModel(review)
	reviewModel.ProductID = product_id
	if reviewModel.Rating > 5 {
		return 0, errors.New("nilai rating maksimal 5")
	}
	tx := repo.db.Create(&reviewModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	_, err := repo.UpdateRating(reviewModel.Rating, product_id)
	if err != nil {
		return 0, err
	}

	return reviewModel.ID, nil
}

func (repo Review) UpdateRating(Rating float64, product_id uint) (float64, error) {

	review := []features.Review{}
	FoundProductId := repo.db.Find(&review, "product_id=?", product_id)
	if FoundProductId.Error != nil {
		return 0, errors.New(FoundProductId.Error.Error() + ", failed to get review id")
	}

	var product features.Product
	ratingAvr, err := AverageRatingsInsert(review, Rating)
	if err != nil {
		return 0, err
	}
	tx := repo.db.Model(&product).Where("id=?", product_id).Update("ratings", ratingAvr)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return product.Ratings, nil
}

func New(db *gorm.DB) review.ReviewData {
	return Review{
		db: db,
	}
}
