package handler

import "belajar/bareng/features"

type Response struct {
	Id 			uint			`json:"id"`
	NamaProduct string          `json:"nama_product"`
	Rating      float64         `json:"rating"`
	Deskripsi   string          `json:"deskripsi"`
	Gambar      []ResponseImage `json:"gambar,omitempty"`
}

type ResponseImage struct {
	Link string `json:"link"`
}

func EntityToResponse(data features.ReviewEntity) Response{
	var images []ResponseImage
	for _,image:= range data.ImagesReview{
		images = append(images, GambarResponse(image))
	}
	return Response{
		Id: 		 data.Id,
		NamaProduct: data.Products.Nama,
		Rating: 	 data.Rating,
		Deskripsi: 	 data.Deskripsi,
		Gambar: 	 images,
	}
}
func GambarResponse(data features.ReviewImagesEntity) ResponseImage{
	return ResponseImage{
		Link: data.Link,
	}
}