package handler

import "belajar/bareng/features"

type Response struct {
	Id 			 uint 	`json:"id"`
	VA           string `json:"virtual_account"`
	Bank         string `json:"bank"`
	TotalHarga   string `json:"total_harga"`
	OrderID		 string `json:"order_id"`
}

func EntityToResponse(payment features.PaymentEntity) Response{
	return Response{
		Id: 			payment.Id,
		OrderID: 		payment.OrderID,
		VA: 			payment.VA,
		Bank: 			payment.Bank,
		TotalHarga: 	payment.TransactionPayments.HargaTotal,

	}
}
