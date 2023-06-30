package handler

import "belajar/bareng/features"

type Response struct {
	Id 		  uint   `json:"id"`
	Nama      string `json:"nama,omitempty"`
	Harga     string `json:"harga,omitempty"`
	Deskripsi string `json:"deskripsi,omitempty"`
	Stok      int    `json:"stok,omitempty"`
}

func EntityToResponse(product features.ProductEntity) Response{
	return Response{
		Id: 		product.Id,
		Nama: 		product.Nama,
		Harga: 		product.Harga,
		Deskripsi: 	product.Deskripsi,
		Stok: 		product.Stok,
	}
}

type ResponseAll struct {
	Id 		  uint   `json:"id"`
	Nama      string `json:"nama,omitempty"`
	Harga     string `json:"harga,omitempty"`
	Stok      int    `json:"stok,omitempty"`
}

func EntityToResponseAll(product features.ProductEntity) ResponseAll{
	return ResponseAll{
		Id: 		product.Id,
		Nama: 		product.Nama,
		Harga: 		product.Harga,
		Stok: 		product.Stok,
	}
}