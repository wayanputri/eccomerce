package handler

import (
	"belajar/bareng/features"
	"time"
)

type Response struct {
	Id			 uint	    `json:"id"`
	NamaUser     string 	`json:"nama_user"`
	NamaBarang   string 	`json:"nama_barang"`
	Status       string 	`json:"status"`
	JumlahBarang int 		`json:"jumlah_barang"`
	TotalHarga   string 	`json:"total_harga"`
	CreatedAt	 time.Time	`json:"created_at"`
}

func EntityToResponse(transaction features.TransactionEntity) Response{
	return Response{
		Id: 			transaction.Id,
		NamaUser: 		transaction.Users.Nama,
		NamaBarang: 	transaction.Products.Nama,
		Status: 		transaction.Status,
		JumlahBarang: 	transaction.JumlahBarang,
		TotalHarga: 	transaction.TotalHarga,
		CreatedAt: 		transaction.CreatedAt,	
	}
}
