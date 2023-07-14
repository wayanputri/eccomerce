package handler

import (
	"belajar/bareng/features"
	image "belajar/bareng/features/image/handler"
	"belajar/bareng/features/reviewimage"
	"belajar/bareng/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewImagesHandler struct {
	reviewImageHandler reviewimage.ReviewImageService
}

func (handler *ReviewImagesHandler) AddImage(c echo.Context) error{
	id:=c.Param("id")
	cnv,err:= strconv.Atoi(id)
	if err != nil{
		return helper.FailedRequest(c,"failed convert id", nil)
	}

	link,errUploud:=image.UploadImage(c)
	if errUploud != nil{
		return helper.FailedRequest(c,"failed uploud data "+errUploud.Error(),nil)
	}else{
		var image features.ReviewImagesEntity
		image.Link = link
	
		idAdd,errAdd:=handler.reviewImageHandler.Add(image,uint(cnv))
		if errAdd != nil{
			return helper.InternalError(c,"err add image "+errAdd.Error(),nil)
		}
		return helper.Success(c,"success add image",idAdd)
	}

}

func (handler *ReviewImagesHandler) GetAll(c echo.Context) error{
	data,err:=handler.reviewImageHandler.GetAll()
	if err != nil{
		return helper.InternalError(c,"failed get data "+err.Error(),nil)
	}
	var dataResponse []Response
	for _,response:= range data{
		dataResponse = append(dataResponse, EntityResponse(response))
	}
	return helper.Success(c,"success get image",dataResponse)
}

func (handler *ReviewImagesHandler) Delete(c echo.Context) error{
	id:=c.Param("id")
	snv,err:=strconv.Atoi(id)
	if err != nil{
		return helper.FailedRequest(c,"failed conversi id",nil)
	}
	errDel:=handler.reviewImageHandler.Delete(uint(snv))
	if errDel != nil{
		return helper.InternalError(c,"failed delete image "+errDel.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"success deleted")
}

func New(handler reviewimage.ReviewImageService) *ReviewImagesHandler{
	return &ReviewImagesHandler{
		reviewImageHandler: handler,
	}
}