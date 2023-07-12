package handler

import "belajar/bareng/features"

type Response struct {
	Id            uint                  `json:"id"`
	HargaTotal    string                `json:"harga_total"`
	StatusPayment []ResponsePayment     `json:"status_payment,omitempty"`
	Transaction   []ResponseTransaction `json:"transactions,omitempty"`
}

type ResponsePayment struct{
	Status string `json:"status"`
	OrderID string `json:"order_id"`
}

type ResponseTransaction struct {
	Harga        string          `json:"harga"`
	JumlahBarang int             `json:"jumlah_barang"`
	Product      ResponseProduct `json:"product"`
}

type ResponseProduct struct {
	NamaProduct  string `json:"nama_product"`
	HargaProduct string `json:"harga_product"`
}

func EntityToResponse(data features.TransactionPaymentEntity) Response{
	var transactionss []ResponseTransaction
	for _,transaction:=range data.Transactions{
		transactionss = append(transactionss, TransactionEntityToResponse(transaction))
	}
	var paymentss []ResponsePayment
	for _,payment:=range data.Payments{
		paymentss = append(paymentss, PaymentEntityToResponse(payment))
	}
	return Response{
		Id:            data.Id,
		HargaTotal:    data.HargaTotal,
		StatusPayment: paymentss,
		Transaction:   transactionss,
	}
}

func PaymentEntityToResponse(data features.PaymentEntity) ResponsePayment{
	return ResponsePayment{
		Status: data.Status,
		OrderID: data.OrderID,
	}
}

func TransactionEntityToResponse(data features.TransactionEntity) ResponseTransaction{
	return ResponseTransaction{
		Harga: data.TotalHarga,
		JumlahBarang: data.JumlahBarang,
		Product: ProductEntityToResponse(data.Products),
		
	}
}

func ProductEntityToResponse(data features.ProductEntity) ResponseProduct{
	return ResponseProduct{
		NamaProduct: data.Nama,
		HargaProduct: data.Harga,
	}
}