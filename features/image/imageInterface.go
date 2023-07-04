package image

import "belajar/bareng/features"

type ImageData interface {
	InserImages(image features.ImageEntity) (uint, error)
}

type ImageService interface {
	AddImages(image features.ImageEntity) (uint, error)
}

