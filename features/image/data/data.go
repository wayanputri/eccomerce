package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/image"

	"gorm.io/gorm"
)

type ImageData struct {
	db *gorm.DB
}

// SelectById implements image.ImageData.
func (repo *ImageData) SelectById(imageId uint) (features.ImageEntity, error) {
	var image features.Image
	tx:=repo.db.Preload("Products").First(&image,imageId)
	if tx.Error != nil{
		return features.ImageEntity{},tx.Error
	}
	data:=features.ImageModelToEntity(image)
	return data,nil
}

// InserImages implements image.ImageData.
func (repo *ImageData) InserImages(image features.ImageEntity, productId uint) (uint, error) {

	var product features.Product
	tx := repo.db.First(&product, productId)
	if tx.Error != nil {
		return 0, tx.Error
	}

	imagesModel := features.ImageEntityToModel(image)
	imagesModel.ProductID = productId
	txx := repo.db.Create(&imagesModel)
	if txx.Error != nil {
		return 0, txx.Error
	}

	return imagesModel.ID, nil
}

func New(db *gorm.DB) image.ImageData {
	return &ImageData{
		db: db,
	}
}
