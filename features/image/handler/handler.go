package handler

import (
	"belajar/bareng/features"
	"belajar/bareng/features/image"
	"belajar/bareng/helper"

	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	imageHandler image.ImageService
}

func (handler *ImageHandler) AddImage(c echo.Context) error{
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
	id,err:=handler.imageHandler.AddImages(image)
	if err != nil{
		return helper.InternalError(c,"failed add image"+err.Error(),nil)
	}
	return helper.SuccessCreate(c,"success create image",id)
}

func New(handler image.ImageService) *ImageHandler{
	return &ImageHandler{
		imageHandler: handler,
	}
}