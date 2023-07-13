package handler

import (
	"belajar/bareng/features"
	"belajar/bareng/features/review"
	"belajar/bareng/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	reviewHandler review.ReviewService
}

func (handler ReviewHandler) Add(c echo.Context) error{
	id:=c.Param("product_id")
	idConv,err := strconv.Atoi(id)
	if err != nil{
		return helper.FailedNotFound(c,"gagal conv product id", nil)
	}
	var review features.ReviewEntity
	errBind:=c.Bind(&review)
	if errBind != nil{
		return helper.FailedRequest(c,"gagal bind data",nil)
	}
	idReview,errAdd:=handler.reviewHandler.Add(review,uint(idConv))
	if errAdd != nil{
		return helper.InternalError(c,"failed add review "+errAdd.Error(),nil)
	}
	return helper.SuccessCreate(c,"success create review ",idReview)
}

func (handler ReviewHandler) Delete(c echo.Context) error{
	id:=c.Param("id")
	convId, err:= strconv.Atoi(id)
	if err != nil{
		return helper.FailedNotFound(c,"failed conversi id param",nil)
	}
	errDelete:=handler.reviewHandler.Delete(uint(convId))
	if errDelete != nil{
		return helper.FailedRequest(c,"failed delete review",nil)
	}
	return helper.SuccessWithOutData(c,"deleted success")
}

func New(handler review.ReviewService) *ReviewHandler{
	return &ReviewHandler{
		reviewHandler: handler,
	}
}