package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/transaction"

	"github.com/go-playground/validator/v10"
)

type TransactionService struct {
	transactionService transaction.TransactionData
	validate           *validator.Validate
}

// Delete implements transaction.TransactionService.
func (service *TransactionService) Delete(transaction_id uint) error {
	err :=service.transactionService.Delete(transaction_id)
	if err != nil{
		return err
	}
	return nil
}

// Edit implements transaction.TransactionService.
func (service *TransactionService) Edit(user_id uint, transaction_id uint, transaction features.TransactionEntity) (uint, error) {
	Id, err := service.transactionService.Update(user_id, transaction_id, transaction)
	if err != nil {
		return 0, err
	}
	return Id, nil
}

// GetAll implements transaction.TransactionService.
func (service *TransactionService) GetAll(user_id uint) ([]features.TransactionEntity, error) {
	data, err := service.transactionService.SelectAll(user_id)
	if err != nil {
		return []features.TransactionEntity{}, err
	}
	return data, nil
}

// GetById implements transaction.TransactionService.
func (service *TransactionService) GetById(transaction_id uint, user_id uint) (features.TransactionEntity, error) {
	data, err := service.transactionService.SelectById(transaction_id, user_id)
	if err != nil {
		return features.TransactionEntity{}, err
	}
	return data, nil
}

// Add implements transaction.TransactionService
func (service *TransactionService) Add(user_id uint, product_id uint, transaction features.TransactionEntity) (uint, error) {
	errVal := service.validate.Struct(transaction)
	if errVal != nil {
		return 0, errVal
	}
	id, err := service.transactionService.Insert(user_id, product_id, transaction)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(transactions transaction.TransactionData) transaction.TransactionService {
	return &TransactionService{
		transactionService: transactions,
		validate:           validator.New(),
	}
}
