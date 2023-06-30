package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/payment"
	"errors"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

// Insert implements payment.PaymentData.
func (repo *PaymentRepo) Insert(payment features.PaymentEntity, transactionId uint) (uint, error) {
	var transaction features.Transaction
	tx := repo.db.First(&transaction,transactionId)
	if tx.Error != nil{
		return 0, errors.New("id transaction tidak ditemukan")
	}
	paymentModel := features.PaymentEntityToModel(payment)
	paymentModel.TransactionID = transactionId
	txx:=repo.db.Where("transaction_id = ?",transactionId).Create(&paymentModel)
	if txx.Error != nil{
		return 0,txx.Error
	}
	return paymentModel.ID, nil
}

func New(db *gorm.DB) payment.PaymentData {
	return &PaymentRepo{
		db: db,
	}
}
