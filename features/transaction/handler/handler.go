package handler

import (
	"belajar/bareng/app/middlewares"
	"belajar/bareng/features"
	"belajar/bareng/features/transaction"
	"belajar/bareng/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionHandler transaction.TransactionService
}

func (handler *TransactionHandler) AddTransaction(c echo.Context) error{
	id_user:=middlewares.ExtractTokenUserId(c)
	id_product :=c.Param("product_id")
	cnv,err := strconv.Atoi(id_product)
	if err != nil{
		return helper.FailedRequest(c,"gagal mengambil id product",nil)
	}
	var transactions features.TransactionEntity
	errBind:=c.Bind(&transactions)
	if errBind != nil{
		return helper.FailedRequest(c,"gagal mengambil data",nil)
	}
	transaction_id,errAdd:=handler.transactionHandler.Add(id_user,uint(cnv),transactions)
	if errAdd != nil{
		return helper.FailedRequest(c,"failed data transaction "+errAdd.Error(),nil)
	}
	data,errGet:=handler.transactionHandler.GetById(transaction_id,id_user)
	if errGet != nil{
		return helper.FailedNotFound(c,"gagal read data terbaru",nil)
	}
	output := EntityToResponse(data)

	return helper.SuccessCreate(c,"success create data transaction",output)
}

func (handler *TransactionHandler) GetById(c echo.Context)error{
	id :=c.Param("transaksi_id")
	conv,err:=strconv.Atoi(id)
	if err != nil{
		return helper.FailedRequest(c,"id gagal di convert",nil)
	}

	UserId:=middlewares.ExtractTokenUserId(c)
	data,errGet:= handler.transactionHandler.GetById(uint(conv),UserId)
	if errGet != nil{
		return helper.FailedRequest(c,"gagal read transaction "+errGet.Error(),nil)
	}
	output := EntityToResponse(data)
	return helper.Success(c,"succes read transaction",output)

}

func (handler *TransactionHandler) GetAll(c echo.Context) error{
	userId:=middlewares.ExtractTokenUserId(c)
	data,err:=handler.transactionHandler.GetAll(userId)
	if err != nil{
		return helper.FailedRequest(c,"gagal read data",nil)
	}
	var TransactionOutput []Response
	for _,transactions := range data{
		TransactionOutput = append(TransactionOutput, EntityToResponse(transactions))
	}
	return helper.Success(c,"success read all data",map[string]any{
		"data":TransactionOutput,
	})
}

func (handler *TransactionHandler) Edit(c echo.Context) error{
	idUser := middlewares.ExtractTokenUserId(c)
	id:=c.Param("transaksi_id")
	cnv,err:= strconv.Atoi(id)
	if err != nil{
		return helper.FailedRequest(c,"transaction_id tidak valid",nil)
	}
	var transactions features.TransactionEntity
	errBind:=c.Bind(&transactions)
	if errBind != nil{
		return helper.FailedRequest(c,"data tidak gagal di bind",nil)
	}
	_,errEdit:=handler.transactionHandler.Edit(idUser,uint(cnv),transactions)
	if errEdit != nil{
		return helper.InternalError(c,"gagal edit data "+errEdit.Error(),nil)
	}
	data,errGet:=handler.transactionHandler.GetById(uint(cnv),idUser)
	if errGet != nil{
		return helper.FailedNotFound(c,"gagal read data terbaru",nil)
	}
	output := EntityToResponse(data)
	return helper.Success(c,"succes edit data transaction ",output)
}

func (handler *TransactionHandler) Delete(c echo.Context) error{
	userId:=middlewares.ExtractTokenUserId(c)
	id:=c.Param("transaksi_id")
	cnv,err:= strconv.Atoi(id)
	if err != nil{
		return helper.FailedRequest(c,"transaction_id tidak valid",nil)
	}
	_,errGet:=handler.transactionHandler.GetById(uint(cnv),userId)
	if errGet != nil{
		return helper.FailedNotFound(c,"data tidak ditemukan, fail deleted",nil)
	}
	errDelete:=handler.transactionHandler.Delete(uint(cnv))
	if errDelete != nil{
		return helper.FailedRequest(c,"failed delete transaction"+errDelete.Error(),nil)
	}
	return helper.Success(c,"success delete data transaction",cnv)
}

func New(transactions transaction.TransactionService) *TransactionHandler{
	return &TransactionHandler{
		transactionHandler: transactions,
	}
}