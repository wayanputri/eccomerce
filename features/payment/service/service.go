package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/payment"

	"github.com/go-playground/validator/v10"
)

type PaymentData struct {
	paymentData payment.PaymentData
	validate    *validator.Validate
}

// GetById implements payment.PaymentService.
func (service *PaymentData) GetById(payment_id uint) (features.PaymentEntity, error) {
	data,err:=service.paymentData.SelectById(payment_id)
	if err != nil{
		return features.PaymentEntity{},err
	}
	return data, nil
}

// Add implements payment.PaymentService.
func (service *PaymentData) Add(payment features.PaymentEntity, transactionId uint) (uint, error) {
	id, err := service.paymentData.Insert(payment, transactionId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(service payment.PaymentData) payment.PaymentService {
	return &PaymentData{
		paymentData: service,
		validate:    validator.New(),
	}
}
