package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/transactionpayment"
	"strconv"

	"gorm.io/gorm"
)

type TransactionPaymentData struct {
	db *gorm.DB
}

// Insert implements transactionpayment.TransactionPaymentData.
func (data TransactionPaymentData) Insert(transactionID []uint) (uint, error) {

	var transaction []features.Transaction
	txx:=data.db.Where("id in ?",transactionID).Find(&transaction)
	if txx.Error != nil{
		return 0,txx.Error
	}
	
	var harga []string
	for _,frice := range transaction{
		harga = append(harga, frice.TotalHarga)
	}

	var hargatotal int
	for _, hargatot := range harga{
		hargaInt,err:=strconv.Atoi(hargatot)
		if err != nil{
			return 0,err
		}
		hargatotal += hargaInt
	}

	hargaString := strconv.Itoa(hargatotal)

	var payment features.TransactionPayment
	payment.HargaTotal = hargaString

	txxx:=data.db.Create(&payment)
	if txxx.Error != nil{
		return 0,txxx.Error
	}
	tx:=data.db.Table("transactions").Where("id IN ?", transactionID).Update("transaction_payment_id", payment.ID)
	if tx.Error != nil {
		return 0,tx.Error
	}
	return payment.ID,nil
	}

func New(db *gorm.DB) transactionpayment.TransactionPaymentData {
	return TransactionPaymentData{
		db: db,
	}
}
