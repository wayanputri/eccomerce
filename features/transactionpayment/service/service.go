package service

import "belajar/bareng/features/transactionpayment"

type TransactionPaymentService struct {
	transactionPaymentService transactionpayment.TransactionPaymentData
}

// Add implements transactionpayment.TransactionPaymentService.
func (service TransactionPaymentService) Add(transactionID []uint) (uint, error) {
	id,err:=service.transactionPaymentService.Insert(transactionID)
	if err != nil{
		return 0,err
	}
	return id , nil
}

func New(service transactionpayment.TransactionPaymentData) transactionpayment.TransactionPaymentService {
	return TransactionPaymentService{
		transactionPaymentService: service,
	}
}
