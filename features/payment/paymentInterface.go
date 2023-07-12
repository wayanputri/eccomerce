package payment

import (
	"belajar/bareng/features"
)

type PaymentData interface{
	Insert(payment features.PaymentEntity,transactionpaymentId uint) (uint,error)
	SelectById(payment_id uint) (features.PaymentEntity,error)
	UpdateStatus(accept string, OrderID string) (uint,error)
}

type PaymentService interface{
	Add(payment features.PaymentEntity,transactionpaymentId uint) (uint,error)
	GetById(payment_id uint) (features.PaymentEntity,error)
	UpdateStatus(accept string, orderID string) (uint,error)
}