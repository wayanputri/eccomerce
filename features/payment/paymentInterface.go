package payment

import (
	"belajar/bareng/features"
)

type PaymentData interface{
	Insert(payment features.PaymentEntity,transactionId uint) (uint,error)
	SelectById(payment_id uint) (features.PaymentEntity,error)
}

type PaymentService interface{
	Add(payment features.PaymentEntity,transactionId uint) (uint,error)
	GetById(payment_id uint) (features.PaymentEntity,error)
}