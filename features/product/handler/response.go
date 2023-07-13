package handler

import "belajar/bareng/features"

type Response struct {
	Id 		  uint   			`json:"id"`
	Nama      string 			`json:"nama,omitempty"`
	Harga     string 			`json:"harga,omitempty"`
	Deskripsi string 			`json:"deskripsi,omitempty"`
	Stok      int    			`json:"stok,omitempty"`
	Link	  []ImageResponse 	`json:"link_image,omitempty"`
	Ratings   float64			`json:"rating"`
}

type ImageResponse struct{
	Link  	string 
}

type ResponseAll struct {
	Id 		  uint   	   	  `json:"id"`
	Nama      string 		  `json:"nama,omitempty"`
	Harga     string 		  `json:"harga,omitempty"`
	Stok      int    		  `json:"stok,omitempty"`
	Link      ImageResponse   `json:"link,omitempty"`
	Ratings   float64			`json:"rating"`
}

func ImageEntityToResponse(image features.ImageEntity) ImageResponse{
	return ImageResponse{
		Link: image.Link,
	}
}

func EntityToResponse(product features.ProductEntity) Response{
	var images []ImageResponse
	for _,image := range product.Images{
		images = append(images, ImageEntityToResponse(image))
	}
	return Response{
		Id: 		product.Id,
		Nama: 		product.Nama,
		Harga: 		product.Harga,
		Deskripsi: 	product.Deskripsi,
		Stok: 		product.Stok,
		Link:       images,
		Ratings: 	product.Ratings,
	}
}


func EntityToResponseAll(product features.ProductEntity) ResponseAll{
	var imageResponse ImageResponse
	if len(product.Images) > 0 {
		imageResponse = ImageEntityToResponse(product.Images[0])
	}
	return ResponseAll{
		Id: 		product.Id,
		Nama: 		product.Nama,
		Harga: 		product.Harga,
		Stok: 		product.Stok,
		Link: 		imageResponse,
		Ratings: 	product.Ratings,
	}
}