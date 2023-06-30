package transaction

import "belajar/bareng/features"

type TransactionData interface {
	Insert(user_id uint, product_id uint, transaction features.TransactionEntity) (uint,error)
	SelectById(transaction_id uint,user_id uint) (features.TransactionEntity,error)
	SelectAll(user_id uint)([]features.TransactionEntity,error)
	Update(user_id uint, transaction_id uint, transaction features.TransactionEntity) (uint,error)	
	Delete(transaction_id uint)error
}
type TransactionService interface{
	Add(user_id uint, product_id uint, transaction features.TransactionEntity)(uint,error)
	GetById(transaction_id uint,user_id uint) (features.TransactionEntity,error)
	GetAll(user_id uint)([]features.TransactionEntity,error)
	Edit(user_id uint, transaction_id uint, transaction features.TransactionEntity) (uint,error)
	Delete(transaction_id uint)error
}