package handler

import (
	"belajar/bareng/features/transactionpayment"
	"belajar/bareng/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TransactionPaymentHandler struct {
	transactionPaymentHandler transactionpayment.TransactionPaymentService
}

func (handler TransactionPaymentHandler) Add(c echo.Context) error{
	var request Request

	errBind:=c.Bind(&request)
	if errBind != nil{
		return helper.FailedNotFound(c,"error bind data",nil)
	}
	idStrings := strings.Split(request.Id, ",")
	
	var idsInt []uint
	for _, idString := range idStrings {
		id, err := strconv.Atoi(idString)
		if err != nil {
			return helper.FailedRequest(c, "failed to convert ID parameter", nil)
		}
		idsInt = append(idsInt, uint(id))
	}

	id,err:=handler.transactionPaymentHandler.Add(idsInt)
	if err != nil{
		return helper.InternalError(c,"add transaction error "+err.Error(),nil)
	}
	return helper.SuccessCreate(c,"success create ",id)
}

func (handler TransactionPaymentHandler) Delete(c echo.Context) error{
	id:=c.Param("id")
	idConv,err := strconv.Atoi(id)
	if err != nil{
		return helper.FailedRequest(c,"failed conversi id",nil)
	}
	errDelete:=handler.transactionPaymentHandler.Delete(uint(idConv))
	if errDelete != nil {
		return helper.FailedRequest(c,"failed delete "+errDelete.Error(),nil)
	}
	return helper.Success(c,"success deleted",nil)
}

func (handler TransactionPaymentHandler) GetAll(c echo.Context) error{
	data,err:=handler.transactionPaymentHandler.GetAll()
	if err != nil {
		return helper.FailedRequest(c,"failed get all "+err.Error(),nil)
	}
	var dataResponse []Response
	for _,response:= range data{
		dataResponse = append(dataResponse, EntityToResponse(response))
	}
	
	return helper.Success(c,"success get all",dataResponse)
}

func New(service transactionpayment.TransactionPaymentService) *TransactionPaymentHandler{
	return &TransactionPaymentHandler{
		transactionPaymentHandler: service,
	}
}