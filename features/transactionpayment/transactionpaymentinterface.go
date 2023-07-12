package transactionpayment

type TransactionPaymentData interface {
	Insert(transactionID []uint) (uint, error)
	//Delete()
}

type TransactionPaymentService interface {
	Add(transactionID []uint) (uint, error)
}