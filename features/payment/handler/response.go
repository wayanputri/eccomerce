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

func Notifikasi(payment NotificationPayload) features.PaymentEntity{

	return features.PaymentEntity{
		OrderID: payment.OrderID,
		Status:  payment.TransactionStatus,
		Bank: 	 payment.VaNumbers[0].Bank,
		VA: 	 payment.VaNumbers[0].VaNumber,
	}
}

type NotificationPayload struct {
	TransactionTime    string                 `json:"transaction_time"`
	TransactionStatus  string                 `json:"transaction_status"`
	OrderID            string                 `json:"order_id"`
	GrossAmount        string                 `json:"gross_amount"`
	Bank               string                 `json:"bank"`
	VaNumbers          []VirtualAccountNumber `json:"va_numbers"`
}

type VirtualAccountNumber struct {
	Bank       string `json:"bank"`
	VaNumber   string `json:"va_number"`
	Expiration string `json:"expiration_date"`
}
