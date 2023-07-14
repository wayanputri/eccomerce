package handler

import (
	"belajar/bareng/app/middlewares"
	"belajar/bareng/features"
	image "belajar/bareng/features/image/handler"
	"belajar/bareng/features/user"
	"belajar/bareng/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userHandler user.UserService
}

func (handler *UserHandler) AddUser(c echo.Context)error{

	user := features.UserEntity{}
	errBind :=c.Bind(&user)
	if errBind != nil{
		return helper.FailedNotFound(c,"Data tidak ditemukan"+errBind.Error(),nil)
	}

	id, errAdd := handler.userHandler.Add(user)
	if errAdd != nil{
		return helper.FailedRequest(c,"gagal menambah data"+errAdd.Error(),nil)
	}

	data,errGet :=handler.userHandler.GetById(uint(id))
	if errGet != nil{
		return helper.FailedNotFound(c,"gagal read data user",nil)
	}
	Output := EntityToResponse(data)
	return helper.Success(c,"berhasil menambah data user",map[string]any{
		"user":Output,
	})
}

func (handler *UserHandler) LoginUser(c echo.Context) error{
	user := features.LoginUser{}
	errBind := c.Bind(&user)
	if errBind != nil{
		return helper.FailedNotFound(c,"data tidak ditemukan",nil)
	}
	id, errLogin := handler.userHandler.Login(user.Email,user.Password)
	if errLogin != nil{
		if strings.Contains(errLogin.Error(),"validation"){	
			return helper.UnAutorization(c,"email tidak sesuai",nil)
		}else{
			return helper.FailedRequest(c,"password tidak sesuai ",nil)
		}
	}
	data,errGet :=handler.userHandler.GetById(uint(id))
	if errGet != nil{
		return helper.FailedNotFound(c,"data tidak ditemukan",nil)
	}
	Output := EntityToResponse(data)

	accessToken,err:= middlewares.CreateToken(int(id))
	if err != nil{
		return helper.InternalError(c,"gagal membuat token",nil)
	}
	return helper.Success(c,"berhasil login",map[string]any{
		"Token":accessToken,
		"user":Output,
	})
}

func (handler *UserHandler) GetByIdUser(c echo.Context)error{
	
	id := middlewares.ExtractTokenUserId(c)
	data,errGet :=handler.userHandler.GetById(uint(id))
	if errGet != nil{
		return helper.FailedNotFound(c,"data tidak ditemukan",nil)
	}
	Output := EntityToResponse(data)
	return helper.Success(c,"succes read data user",map[string]any{
		"user":Output,
	})

}

func (handler *UserHandler) EditUser(c echo.Context)error{
	id := middlewares.ExtractTokenUserId(c)
	var user features.UserEntity
	errBind:=c.Bind(&user)
	if errBind != nil{
		return helper.FailedNotFound(c,"data tidak ditemukan",nil)
	}
	idUser,err:=handler.userHandler.Edit(id,user)
	if err !=nil{
		return helper.InternalError(c,"gagal update profil user",nil)
	}
	data,errGet:=handler.userHandler.GetById(idUser)
	if errGet != nil{
		return helper.InternalError(c,"gagal menampilkan data",nil)
	}
	Output := EntityToResponse(data)

	return helper.Success(c,"success edit data",map[string]any{
		"data":Output,
	})
}

func (handler *UserHandler) DeleteUser(c echo.Context) error{
	id := middlewares.ExtractTokenUserId(c)
	err := handler.userHandler.Delete(id)
	if err != nil{
		return helper.FailedRequest(c,"delete failed "+err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"succes delete data user")
}

func (handler *UserHandler) Upgrade(c echo.Context) error{
	idUser:=middlewares.ExtractTokenUserId(c)
	var user features.UserEntity
	errBind:=c.Bind(&user)
	if errBind != nil{
		return helper.FailedNotFound(c,"error bind data",nil)
	}
	link,errUploud:=image.UploadImage(c)
	if errUploud != nil{
		return helper.FailedRequest(c,"failed uploud data "+errUploud.Error(),nil)
	}
	user.File = link
	user.Role = "pedagang"
	id,err:=handler.userHandler.Upgrade(idUser,user)
	if err != nil{
		return helper.InternalError(c,"failed upgrade "+err.Error(),nil)
	}
	return helper.Success(c,"success upgrade ",id)
}

func New(users user.UserService) *UserHandler{
	return &UserHandler{
		userHandler: users,
	}
}