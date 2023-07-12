package handler

import (
	"belajar/bareng/app/config"
	"belajar/bareng/features"
	"belajar/bareng/features/payment"
	"belajar/bareng/helper"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentHandler struct {
	paymentHandler payment.PaymentService
}
func (handler *PaymentHandler) Add(c echo.Context) error{
	id:=c.Param("transaksi_payment_id")
	cnv,err:=strconv.Atoi(id)
	if err != nil{
		return helper.FailedNotFound(c, "id tidak terbaca", nil)
	}
	var payment features.PaymentEntity
	errBind:=c.Bind(&payment)
	if errBind != nil{
		return helper.FailedRequest(c,"data tidak terbaca",nil)
	}
	PaymentId,errAdd:= handler.paymentHandler.Add(payment,uint(cnv))
	if errAdd != nil{
		return helper.FailedRequest(c,"failet create data payment"+errAdd.Error(),nil)
	}

	data,errGet:=handler.paymentHandler.GetById(PaymentId)
	if errGet != nil{
		return helper.FailedRequest(c,"failed get data payment",nil)
	}
	dataResponse:=EntityToResponse(data)

	return helper.Success(c,"success add data payment",dataResponse)
}

func (handler *PaymentHandler) Notification(c echo.Context) error {
	cfg := config.InitConfig()
	var client = coreapi.Client{}
	client.New(cfg.KEY_SERVER_MIDTRANS, midtrans.Sandbox)

	client.Options.SetPaymentAppendNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")

	var notificationPayload map[string]interface{}

	err := c.Bind(&notificationPayload)
	if err != nil {

		return helper.FailedRequest(c,"gagal bind data",nil)
	}

	orderID, exists := notificationPayload["order_id"].(string)
	if !exists {
	
		return helper.FailedRequest(c,"failed get orderId",nil)
	}

	transactionStatusResp, errTrans := client.CheckTransaction(orderID) 
	if errTrans != nil {
		return helper.InternalError(c,"internal server error "+errTrans.Error(),nil)
	}
	
	id,errUpdate:=handler.paymentHandler.UpdateStatus(transactionStatusResp.TransactionStatus,orderID)
	if errUpdate != nil{
		return helper.InternalError(c,"failed update data "+errUpdate.Error(),nil)
	}
	return helper.Success(c,"successfuly",id)
}


func New(service payment.PaymentService) *PaymentHandler{
	return &PaymentHandler{
		paymentHandler: service,
	}
}