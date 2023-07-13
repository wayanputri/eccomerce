package handler

import (
	"belajar/bareng/features"
	data "belajar/bareng/features/image/handler"
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

	link,errUploud:=data.UploadImage(c)
	if errUploud != nil{
		return helper.FailedRequest(c,"failed uploud data "+errUploud.Error(),nil)
	}
	var image features.ReviewImagesEntity
	image.Link = link

	idAdd,errAdd:=handler.reviewImageHandler.Add(image,uint(cnv))
	if errAdd != nil{
		return helper.InternalError(c,"err add image "+errAdd.Error(),nil)
	}
	return helper.Success(c,"success add image",idAdd)
	
}

func New(handler reviewimage.ReviewImageService) *ReviewImagesHandler{
	return &ReviewImagesHandler{
		reviewImageHandler: handler,
	}
}