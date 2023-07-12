package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/transactionpayment"

	"gorm.io/gorm"
)

type TransactionPaymentData struct {
	db *gorm.DB
}

// SelectAll implements transactionpayment.TransactionPaymentData.
func (data TransactionPaymentData) SelectAll() ([]features.TransactionPaymentEntity, error) {
	var paymentTransaction []features.TransactionPayment
	tx:=data.db.Preload("Transactions").Preload("Transactions.Products").Preload("Payments").Find(&paymentTransaction)
	if tx.Error != nil {
		return []features.TransactionPaymentEntity{},tx.Error
	}
	var transactionreads []features.TransactionPaymentEntity
	for _,paymenttransaction := range paymentTransaction{
		transactionreads=append(transactionreads,features.TransactionPaymentModelToEntity(paymenttransaction))
	}
	
	return transactionreads,nil
}

// Delete implements transactionpayment.TransactionPaymentData.
func (data TransactionPaymentData) Delete(transactionpaymentID uint) error {
	var transactionPayment features.TransactionPayment
	txx := data.db.First(&transactionPayment, transactionpaymentID)
	if txx.Error != nil {
		return txx.Error
	}
	tx := data.db.Delete(&transactionPayment, transactionpaymentID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Insert implements transactionpayment.TransactionPaymentData.
func (data TransactionPaymentData) Insert(transactionID []uint) (uint, error) {

	var transaction []features.Transaction
	txx := data.db.Where("id in ?", transactionID).Find(&transaction)
	if txx.Error != nil {
		return 0, txx.Error
	}

	harga := AppendHarga(transaction)
	hargaString:=ConversiHarga(harga)

	var payment features.TransactionPayment
	payment.HargaTotal = hargaString

	txxx := data.db.Create(&payment)
	if txxx.Error != nil {
		return 0, txxx.Error
	}
	tx := data.db.Table("transactions").Where("id IN ?", transactionID).Update("transaction_payment_id", payment.ID)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return payment.ID, nil
}

func New(db *gorm.DB) transactionpayment.TransactionPaymentData {
	return TransactionPaymentData{
		db: db,
	}
}
