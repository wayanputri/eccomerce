package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/transactionpayment"
)

type TransactionPaymentService struct {
	transactionPaymentService transactionpayment.TransactionPaymentData
}

// GetAll implements transactionpayment.TransactionPaymentService.
func (service TransactionPaymentService) GetAll() ([]features.TransactionPaymentEntity, error) {
	data,err:=service.transactionPaymentService.SelectAll()
	if err != nil{
		return []features.TransactionPaymentEntity{},err
	}
	return data,nil
}

// Delete implements transactionpayment.TransactionPaymentService.
func (service TransactionPaymentService) Delete(transactionpaymentID uint) error {
	err := service.transactionPaymentService.Delete(transactionpaymentID)
	if err != nil {
		return err
	}
	return nil
}

// Add implements transactionpayment.TransactionPaymentService.
func (service TransactionPaymentService) Add(transactionID []uint) (uint, error) {
	id, err := service.transactionPaymentService.Insert(transactionID)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(service transactionpayment.TransactionPaymentData) transactionpayment.TransactionPaymentService {
	return TransactionPaymentService{
		transactionPaymentService: service,
	}
}
