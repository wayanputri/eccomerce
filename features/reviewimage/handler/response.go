package handler

import "belajar/bareng/features"

type Response struct {
	Id          uint   `json:"id"`
	NamaProduct string `json:"nama_product"`
	Rating 	float64 `json:"rating"`
	Link        string `json:"link"`
}

func EntityResponse(data features.ReviewImagesEntity) Response{
	return Response{
		Id: 		 data.Id,
		NamaProduct: data.Reviews.Products.Nama,
		Rating: 	 data.Reviews.Rating,
		Link: 		 data.Link,
	}
}
