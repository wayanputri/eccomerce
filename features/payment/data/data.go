package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/payment"
	"belajar/bareng/helper"
	"errors"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

// UpdateStatus implements payment.PaymentData.
func (repo *PaymentRepo) UpdateStatus(accept string, OrderID string) (uint, error) {
	var payment features.Payment
	txx:=repo.db.First(&payment,"order_id=?",OrderID)
	if txx.Error != nil{
		return 0,txx.Error
	}
	tx := repo.db.Model(&payment).Where("order_id=?",OrderID).Update("status",accept )
	if tx.Error != nil{
		return 0,tx.Error
	}
	return payment.ID,nil
}

// SelectById implements payment.PaymentData.
func (repo *PaymentRepo) SelectById(payment_id uint) (features.PaymentEntity, error) {
	var paymentModel features.Payment
	tx := repo.db.Preload("TransactionPayments").First(&paymentModel, payment_id)
	if tx.Error != nil {
		return features.PaymentEntity{}, tx.Error
	}
	data := features.PaymentModelToEntity(paymentModel)
	return data, nil

}

// Insert implements payment.PaymentData.
func (repo *PaymentRepo) Insert(payment features.PaymentEntity, transactionpaymentId uint) (uint, error) {
	var transactionpayment features.TransactionPayment
	tx := repo.db.First(&transactionpayment, transactionpaymentId)
	if tx.Error != nil {
		return 0, errors.New("id transaction payment tidak ditemukan")
	}
	orderID, errOrderId := helper.GenerateUUID()
	if errOrderId != nil {
		return 0, errOrderId
	}
	paymentModel := features.PaymentEntityToModel(payment)
	dataResponse := requestCreditCard(transactionpayment.HargaTotal, orderID, paymentModel.Bank)
	paymentModel = PaymentResponse(paymentModel, transactionpaymentId, orderID, dataResponse)

	txx := repo.db.Where("transaction_payment_id = ?", transactionpaymentId).Create(&paymentModel)
	if txx.Error != nil {
		return 0, txx.Error
	}
	return paymentModel.ID, nil
}

func New(db *gorm.DB) payment.PaymentData {
	return &PaymentRepo{
		db: db,
	}
}
