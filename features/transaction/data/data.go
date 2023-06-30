package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/transaction"
	"strconv"

	"gorm.io/gorm"
)

type TransactionData struct {
	db *gorm.DB
}

// Delete implements transaction.TransactionData.
func (repo *TransactionData) Delete(transaction_id uint) error {
	var transaction features.Transaction
	tx:=repo.db.Delete(&transaction,transaction_id)
	if tx.Error != nil{
		return tx.Error
	}
	return nil
}

// Update implements transaction.TransactionData.
func (repo *TransactionData) Update(user_id uint, transaction_id uint, transaction features.TransactionEntity) (uint, error) {

	var ModalTransaction features.Transaction
	tx := repo.db.First(&ModalTransaction, transaction_id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	var product features.Product
	txy := repo.db.First(&product, ModalTransaction.ProductID)
	if txy.Error != nil {
		return 0, txy.Error
	}
	harga, errConv := strconv.Atoi(product.Harga)
	if errConv != nil {
		return 0, errConv
	}
	totalHarga := TotalHarga(transaction.JumlahBarang, harga)
	hargacnv := strconv.Itoa(totalHarga)
	transaction.TotalHarga = hargacnv
	errStok := KetersediaanStok(transaction.JumlahBarang, product.Stok)
	if errStok != nil {
		return 0, errStok
	}
	txx := repo.db.Model(&ModalTransaction).Where("id=? AND user_id=?", transaction_id, user_id).Updates(features.TransactionEntityToModel(transaction))
	if txx.Error != nil {
		return 0, tx.Error
	}

	return ModalTransaction.ID, nil
}

// SelectAll implements transaction.TransactionData.
func (repo *TransactionData) SelectAll(user_id uint) ([]features.TransactionEntity, error) {
	var transactions []features.Transaction
	tx := repo.db.Preload("Users").Preload("Products").Find(&transactions, "user_id=?", user_id)
	if tx.Error != nil {
		return []features.TransactionEntity{}, tx.Error
	}
	var transactionEntity []features.TransactionEntity
	for _, transaction := range transactions {
		transactionEntity = append(transactionEntity, features.TransactionModelToEntity(transaction))
	}
	return transactionEntity, nil
}

// SelectById implements transaction.TransactionData.
func (repo *TransactionData) SelectById(transaction_id uint, user_id uint) (features.TransactionEntity, error) {
	var transactions features.Transaction
	tx := repo.db.Preload("Users").Preload("Products").First(&transactions, "id=? AND user_id = ?", transaction_id, user_id)
	if tx.Error != nil {
		return features.TransactionEntity{}, tx.Error
	}
	data := features.TransactionModelToEntity(transactions)
	return data, nil
}

// Insert implements transaction.TransactionData
func (repo *TransactionData) Insert(user_id uint, product_id uint, transaction features.TransactionEntity) (uint, error) {
	var dataProduct features.Product

	tx := repo.db.First(&dataProduct, product_id)
	if tx.Error != nil {
		return 0, tx.Error
	}

	TransactionModel := features.TransactionEntityToModel(transaction)
	harga, errHarga := strconv.Atoi(dataProduct.Harga)
	if errHarga != nil {
		return 0, errHarga
	}

	totalHarga := TotalHarga(transaction.JumlahBarang, harga)
	Harga := strconv.Itoa(totalHarga)

	data := InsertModel(TransactionModel, Harga, user_id, product_id)

	errStok := KetersediaanStok(TransactionModel.JumlahBarang, dataProduct.Stok)
	if errStok != nil {
		return 0, errStok
	}
	txx := repo.db.Create(&data)
	if txx.Error != nil {
		return 0, tx.Error
	}

	return data.ID, nil
}

func New(db *gorm.DB) transaction.TransactionData {
	return &TransactionData{
		db: db,
	}
}
