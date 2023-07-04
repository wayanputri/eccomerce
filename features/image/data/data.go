package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/image"

	"gorm.io/gorm"
)

type ImageData struct {
	db *gorm.DB
}

// InserImages implements image.ImageData.
func (repo *ImageData) InserImages(image features.ImageEntity) (uint, error) {
	imagesModel := features.ImageEntityToModel(image)
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
