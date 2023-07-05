package image

import "belajar/bareng/features"

type ImageData interface {
	InserImages(image features.ImageEntity,productId uint) (uint, error)
	SelectById(imageId uint) (features.ImageEntity,error)
}

type ImageService interface {
	AddImages(image features.ImageEntity,productId uint) (uint, error)
	GetById(imageId uint) (features.ImageEntity,error)
}

