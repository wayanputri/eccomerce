package data

import (
	"belajar/bareng/features"

	"github.com/midtrans/midtrans-go/coreapi"
)

func PaymentResponse(paymentModel features.Payment,transactionpaymentId uint, orderID string, dataResponse *coreapi.ChargeResponse) features.Payment{
	paymentModel.TransactionPaymentID = transactionpaymentId
	paymentModel.OrderID = orderID
	paymentModel.Bank = dataResponse.VaNumbers[0].Bank
	paymentModel.OrderID = dataResponse.OrderID
	paymentModel.VA = dataResponse.VaNumbers[0].VANumber
	paymentModel.Status = dataResponse.TransactionStatus	
	return paymentModel
}



