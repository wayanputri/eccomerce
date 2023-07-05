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
func (repo *PaymentRepo) UpdateStatus(payload features.PaymentEntity, OrderID string) (uint, error) {
	var payment features.Payment
	txx:=repo.db.First(&payment,"order_id=?",OrderID)
	if txx.Error != nil{
		return 0,txx.Error
	}
	tx := repo.db.Model(&payment).Where("order_id=?",OrderID).Updates(features.PaymentEntityToModel(payload))
	if tx.Error != nil{
		return 0,tx.Error
	}
	return payment.ID,nil
}

// SelectById implements payment.PaymentData.
func (repo *PaymentRepo) SelectById(payment_id uint) (features.PaymentEntity, error) {
	var paymentModel features.Payment
	tx := repo.db.Preload("Transactions").Preload("Transactions.Users").Preload("Transactions.Products").First(&paymentModel, payment_id)
	if tx.Error != nil {
		return features.PaymentEntity{}, tx.Error
	}
	data := features.PaymentModelToEntity(paymentModel)
	return data, nil

}

// Insert implements payment.PaymentData.
func (repo *PaymentRepo) Insert(payment features.PaymentEntity, transactionId uint) (uint, error) {
	var transaction features.Transaction
	tx := repo.db.First(&transaction, transactionId)
	if tx.Error != nil {
		return 0, errors.New("id transaction tidak ditemukan")
	}
	orderID, errOrderId := helper.GenerateUUID()
	if errOrderId != nil {
		return 0, errOrderId
	}
	paymentModel := features.PaymentEntityToModel(payment)
	dataResponse := requestCreditCard(transaction.TotalHarga, orderID)
	paymentModel = PaymentResponse(paymentModel, transactionId, orderID, dataResponse)

	txx := repo.db.Where("transaction_id = ?", transactionId).Create(&paymentModel)
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
