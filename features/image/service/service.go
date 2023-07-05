package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/image"

	"github.com/go-playground/validator/v10"
)

type ImageService struct {
	imageService image.ImageData
	validate     *validator.Validate
}

// GetById implements image.ImageService.
func (service *ImageService) GetById(imageId uint) (features.ImageEntity, error) {
	data,err:=service.imageService.SelectById(imageId)
	if err != nil{
		return features.ImageEntity{},err
	}
	return data, nil
}

// AddImages implements image.ImageService.
func (service *ImageService) AddImages(image features.ImageEntity, productId uint) (uint, error) {
	id, err := service.imageService.InserImages(image, productId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(service image.ImageData) image.ImageService {
	return &ImageService{
		imageService: service,
		validate:     validator.New(),
	}
}
