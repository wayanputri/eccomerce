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

// updateTransactionStatus implements payment.PaymentService.
func (service *PaymentData) UpdateStatus(accept string, orderID string) (uint, error) {
	data,err:=service.paymentData.UpdateStatus(accept,orderID)
	if err != nil{
		return 0,err
	}
	return data,nil
}

// GetById implements payment.PaymentService.
func (service *PaymentData) GetById(payment_id uint) (features.PaymentEntity, error) {
	data, err := service.paymentData.SelectById(payment_id)
	if err != nil {
		return features.PaymentEntity{}, err
	}
	return data, nil
}

// Add implements payment.PaymentService.
func (service *PaymentData) Add(payment features.PaymentEntity, transactionpaymentId uint) (uint, error) {
	id, err := service.paymentData.Insert(payment, transactionpaymentId)
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
