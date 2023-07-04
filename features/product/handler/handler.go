package handler

import (
	"belajar/bareng/app/middlewares"
	"belajar/bareng/features"
	"belajar/bareng/features/product"
	"belajar/bareng/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productHandler product.ProductServise
}

func (handler *ProductHandler) AddProduct(c echo.Context)error{
	userId:=middlewares.ExtractTokenUserId(c)
	var InputProduct features.ProductEntity

	errBind:=c.Bind(&InputProduct)
	if errBind != nil{
		return helper.FailedRequest(c,"data tidak ditemukan",nil)
	}
	InputProduct.UserId = userId
	id,err:=handler.productHandler.Add(InputProduct)
	if err != nil{
		return helper.InternalError(c, "create data failed karena nama,harga,stok tidak boleh kosong",nil)
	}
	data, errGet:=handler.productHandler.GetById(id)
	if errGet != nil{
		return helper.FailedRequest(c,"get user by id failed",nil)
	}
	dataOutput := EntityToResponse(data)
	return helper.Success(c,"berhasil read data by id",dataOutput)
}
func (handler *ProductHandler) GetAllProduct(c echo.Context) error{
	data,err:= handler.productHandler.GetAll()
	if err != nil{
		return helper.FailedNotFound(c,"data tidak ditemukan",nil)
	}
	var dataOutput []ResponseAll
	for _,output:=range data{
		Output:=EntityToResponseAll(output)
		dataOutput = append(dataOutput,Output)
	}
	return helper.Success(c,"success read data product",dataOutput)
	
}

func (handler *ProductHandler) GetById(c echo.Context) error{
	id:= c.Param("id")
	cnv,err:=strconv.Atoi(id)
	if err != nil{
		return helper.FailedNotFound(c,"id tidak ditemukan",nil)
	}
	data,errGet:=handler.productHandler.GetById(uint(cnv))
	if errGet != nil{
		return helper.FailedRequest(c,"get product by id failed",nil)
	}
	dataOutput := EntityToResponse(data)
	return helper.Success(c,"berhasil read data by id",dataOutput)
}

func (handler *ProductHandler) Edit(c echo.Context)error{
	id:=c.Param("id")
	cvn,err:= strconv.Atoi(id)
	if err != nil{
		return helper.FailedNotFound(c,"id tidak ditemukan",nil)
	}
	var product features.ProductEntity
	errbind:=c.Bind(&product)
	if errbind != nil{
		return helper.FailedNotFound(c,"data tidak ditemukan",nil)
	}
	UserId:=middlewares.ExtractTokenUserId(c)
	errUser:=handler.productHandler.SelectByUserId(UserId)
	if errUser != nil{
		return helper.FailedRequest(c,"tidak bisa mengedit product user lain",nil)
	}
	idEdit,errEdit:=handler.productHandler.Edit(uint(cvn),product)
	if errEdit != nil{
		return helper.InternalError(c,"gagal mengedit data",nil)
	}
	data,errGet:=handler.productHandler.GetById(idEdit)
	if errGet != nil{
		return helper.FailedRequest(c,"get product by id failed",nil)
	}
	dataoutput := EntityToResponse(data)
	return helper.Success(c,"success edit data product",dataoutput)
}
func (handler *ProductHandler) DeleteProduct(c echo.Context) error{
	id:=c.Param("id")
	cvn,err:= strconv.Atoi(id)
	if err != nil{
		return helper.FailedNotFound(c,"id tidak ditemukan",nil)
	}

	UserId:=middlewares.ExtractTokenUserId(c)
	errUser:=handler.productHandler.SelectByUserId(UserId)
	if errUser != nil{
		return helper.FailedRequest(c,"tidak bisa menghapus product user lain",nil)
	}

	errDelete:=handler.productHandler.Delete(uint(cvn))
	if errDelete != nil{
		return helper.InternalError(c,"failed delete data",nil)
	}
	return helper.SuccessWithOutData(c,"success delete data")
}


func New(products product.ProductServise) *ProductHandler{
	return &ProductHandler{
		productHandler: products,
	}
}