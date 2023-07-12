package transactionpayment

import "belajar/bareng/features"

type TransactionPaymentData interface {
	Insert(transactionID []uint) (uint, error)
	Delete(transactionpaymentID uint) error
	SelectAll() ([]features.TransactionPaymentEntity,error)
}

type TransactionPaymentService interface {
	Add(transactionID []uint) (uint, error)
	Delete(transactionpaymentID uint) error
	GetAll() ([]features.TransactionPaymentEntity,error)
}