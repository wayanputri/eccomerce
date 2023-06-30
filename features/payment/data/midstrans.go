package data

import (
	"belajar/bareng/app/config"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func requestCreditCard(harga string,orderId string) *coreapi.ChargeResponse{

	cfg := config.InitConfig()
	midtrans.ServerKey = cfg.KEY_SERVER_MIDTRANS
	authString := encodeAuthString(midtrans.ServerKey, "")
	fmt.Println("AUTH_STRING:", authString)
	midtrans.Environment = midtrans.Sandbox

	totalHarga,_ := strconv.Atoi(harga)

	bankTransferReq := &coreapi.ChargeReq{
		PaymentType:        coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{OrderID: orderId, GrossAmt: int64(totalHarga)}, 
		BankTransfer:       &coreapi.BankTransferDetails{Bank: "bca"},
		Metadata:           nil,
	}
	coreApiRes, errCore := coreapi.ChargeTransaction(bankTransferReq)
	if errCore != nil {
		log.Fatal("Failed to charge transaction:", errCore)
	}
	return coreApiRes
}
func encodeAuthString(username, password string) string {
	auth := username + ":" + password
	authBytes := []byte(auth)
	encodedAuth := base64.StdEncoding.EncodeToString(authBytes)
	return encodedAuth
}