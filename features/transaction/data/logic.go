package data

import (
	"belajar/bareng/features"
	"errors"
)

func InsertModel(transaction features.Transaction, harga string,user_id uint,product_id uint) features.Transaction{

	transaction.TotalHarga = harga
	transaction.UserID = user_id
	transaction.ProductID = product_id
	transaction.Status = "pending"
	return transaction
}

func TotalHarga(JumlahBarang int,harga int)int{
	TotalHarga := JumlahBarang *harga
	return TotalHarga
}

func KetersediaanStok(jumlahBarang int,stok int)error{
	if jumlahBarang > stok {
		return errors.New("product terbatas, stok tidak cukup")
	} 
	return nil
}