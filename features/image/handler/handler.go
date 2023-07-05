package handler

import (
	"belajar/bareng/features"
	"belajar/bareng/features/image"
	"belajar/bareng/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	imageHandler image.ImageService
}

func (handler *ImageHandler) AddImage(c echo.Context) error{
	productId:=c.Param("product_id")
	cnv,errProductId := strconv.Atoi(productId)
	if errProductId != nil{
		return helper.FailedRequest(c,"failed conveert product id",nil)
	}
	var image features.ImageEntity
	errBind:=c.Bind(&image)
	if errBind != nil{
		return helper.FailedRequest(c,"failed bind data", nil)
	}
	link,errUploud:=UploadImage(c)
	if errUploud != nil{
		return helper.FailedRequest(c,"failed uploud data "+errUploud.Error(),nil)
	}
	image.Link = link
	id,err:=handler.imageHandler.AddImages(image,uint(cnv))
	if err != nil{
		return helper.InternalError(c,"failed add image"+err.Error(),nil)
	}
	
	data,err:=handler.imageHandler.GetById(id)
	if err != nil{
		return helper.InternalError(c,"gagal get data "+err.Error(),nil)
	}
	dataResponse:=EntityToResponse(data)
	return helper.Success(c,"succesfuly",dataResponse)
}

func (handler *ImageHandler) GetById(c echo.Context)error{
	id:=c.Param("image_id")
	cnv,errParam := strconv.Atoi(id)
	if errParam != nil{
		return helper.FailedNotFound(c,"faild convert id",nil)
	}

	data,err:=handler.imageHandler.GetById(uint(cnv))
	if err != nil{
		return helper.InternalError(c,"gagal get data "+err.Error(),nil)
	}
	dataResponse:=EntityToResponse(data)
	return helper.Success(c,"succesfuly",dataResponse)
}

func (handler *ImageHandler) GetAll(c echo.Context)error{

	data,err:=handler.imageHandler.GetAll()
	if err != nil{
		return helper.InternalError(c,"gagal get data "+err.Error(),nil)
	}
	var dataImages []Response
	for _,image := range data{
		dataResponse:=EntityToResponse(image)
		dataImages = append(dataImages, dataResponse)
	}
	return helper.Success(c,"succesfuly",dataImages)
}

func New(handler image.ImageService) *ImageHandler{
	return &ImageHandler{
		imageHandler: handler,
	}
}