package handler

import (
	"belajar/bareng/features"
	"belajar/bareng/features/payment"
	"belajar/bareng/helper"
	"strconv"

	"github.com/labstack/echo/v4"
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

func New(service payment.PaymentService) *PaymentHandler{
	return &PaymentHandler{
		paymentHandler: service,
	}
}