package handler

import (
	"belajar/bareng/features"
)

type Response struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	ImageName   string `json:"image_name"`
	Link        string `json:"link"`
}

func EntityToResponse(image features.ImageEntity) Response{
	return Response{
		Id: 			image.Id,
		ProductName: 	image.Products.Nama,
		ImageName: 		image.Nama,
		Link: 			image.Link,
		}
}