package handler

import "belajar/bareng/features"

type Response struct {
	Id 			 uint 	`json:"id"`
	NamaUser     string `json:"nama_user"`
	NamaBarang   string `json:"nama_product"`
	Alamat       string `json:"alamat"`
	VA           string `json:"virtual_account"`
	Bank         string `json:"bank"`
	TotalHarga   string `json:"total_harga"`
	JumlahBarang int    `json:"jumlah_barang"`
}

func EntityToResponse(payment features.PaymentEntity) Response{
	return Response{
		Id: 			payment.Id,
		NamaUser: 		payment.Transactions.Users.Nama,
		NamaBarang: 	payment.Transactions.Products.Nama,
		Alamat: 		payment.Transactions.Users.Alamat,
		VA: 			payment.VA,
		Bank: 			payment.Bank,
		TotalHarga: 	payment.Transactions.TotalHarga,
		JumlahBarang: 	payment.Transactions.JumlahBarang,
	}
}

type NotificationPayload struct {
	OrderID    string  `json:"order_id"`
	Amount     float64 `json:"gross_amounts"`
	Status     string  `json:"fraud_status"`
	CustomerID string  `json:"customer_id"`
	Customers  []features.UserEntity
}

func Notifikasi(payment NotificationPayload) features.PaymentEntity{
	return features.PaymentEntity{
		OrderID: payment.OrderID,
		Status:  payment.Status,
	}
}