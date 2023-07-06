package handler

import (
	"belajar/bareng/app/config"
	"belajar/bareng/features"
	"belajar/bareng/features/payment"
	"belajar/bareng/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentHandler struct {
	paymentHandler payment.PaymentService
}
func (handler *PaymentHandler) Add(c echo.Context) error{
	id:=c.Param("transaksi_id")
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
	// 1. Initiate Gateway
	cfg := config.InitConfig()
	var client = coreapi.Client{}
	client.New(cfg.KEY_SERVER_MIDTRANS, midtrans.Sandbox)

	// 2. Set Payment Override or Append via gateway options for specific request
	client.Options.SetPaymentAppendNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")

	// 1. Initialize empty map
	var notificationPayload map[string]interface{}

	// 2. Parse JSON request body and use it to set json to payload
	err := c.Bind(&notificationPayload)
	if err != nil {
		// do something on error when decoding
		return err
	}

	// 3. Get order-id from payload
	orderID, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return echo.NewHTTPError(http.StatusBadRequest, "order_id not found")
	}

	// 4. Check transaction to Midtrans with param orderID
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