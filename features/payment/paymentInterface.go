package payment

import (
	"belajar/bareng/features"
)

type PaymentData interface{
	Insert(payment features.PaymentEntity,transactionId uint) (uint,error)
}

type PaymentService interface{
	Add(payment features.PaymentEntity,transactionId uint) (uint,error)
}