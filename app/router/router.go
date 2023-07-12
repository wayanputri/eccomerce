package router

import (
	"belajar/bareng/app/middlewares"
	DataUser "belajar/bareng/features/user/data"
	HandlerUser "belajar/bareng/features/user/handler"
	ServiseUser "belajar/bareng/features/user/service"

	DataProduct "belajar/bareng/features/product/data"
	HandlerProduct "belajar/bareng/features/product/handler"
	ServiseProduct "belajar/bareng/features/product/service"

	DataTransaction "belajar/bareng/features/transaction/data"
	HandlerTransaction "belajar/bareng/features/transaction/handler"
	ServiseTransaction "belajar/bareng/features/transaction/service"

	DataPayment "belajar/bareng/features/payment/data"
	HandlerPayment "belajar/bareng/features/payment/handler"
	ServisePayment "belajar/bareng/features/payment/service"

	DataImage "belajar/bareng/features/image/data"
	HandlerImage "belajar/bareng/features/image/handler"
	ServiseImage "belajar/bareng/features/image/service"

	DataTransactionPayment "belajar/bareng/features/transactionpayment/data"
	HandlerTransactionPayment "belajar/bareng/features/transactionpayment/handler"
	ServiseTransactionPayment "belajar/bareng/features/transactionpayment/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(c *echo.Echo, db *gorm.DB){
	dataUser := DataUser.New(db)
	serviceUser := ServiseUser.New(dataUser)
	handlerUser := HandlerUser.New(serviceUser)

	c.POST("/users",handlerUser.AddUser)
	c.POST("/login",handlerUser.LoginUser)
	c.GET("/users",handlerUser.GetByIdUser,middlewares.JWTMiddleware())
	c.PUT("/users",handlerUser.EditUser,middlewares.JWTMiddleware())
	c.DELETE("/users",handlerUser.DeleteUser,middlewares.JWTMiddleware())

	dataProduct := DataProduct.New(db)
	serviceProduct := ServiseProduct.New(dataProduct)
	handlerProduct := HandlerProduct.New(serviceProduct)

	c.POST("/products",handlerProduct.AddProduct,middlewares.JWTMiddleware())
	c.GET("/products",handlerProduct.GetAllProduct)
	c.GET("/products/:id",handlerProduct.GetById)
	c.PUT("/products/:id",handlerProduct.Edit,middlewares.JWTMiddleware())
	c.DELETE("/products/:id",handlerProduct.DeleteProduct,middlewares.JWTMiddleware())

	dataTransaction := DataTransaction.New(db)
	serviceTransaction := ServiseTransaction.New(dataTransaction)
	handlerTransaction := HandlerTransaction.New(serviceTransaction)

	c.POST("/products/:product_id/transactions",handlerTransaction.AddTransaction,middlewares.JWTMiddleware())
	c.GET("/transactions/:transaksi_id",handlerTransaction.GetById,middlewares.JWTMiddleware())
	c.GET("/transactions",handlerTransaction.GetAll,middlewares.JWTMiddleware())
	c.PUT("/transactions/:transaksi_id",handlerTransaction.Edit,middlewares.JWTMiddleware())
	c.DELETE("/transactions/:transaksi_id",handlerTransaction.Delete,middlewares.JWTMiddleware())
	

	dataPayment := DataPayment.New(db)
	servicePayment := ServisePayment.New(dataPayment)
	handlerPayment := HandlerPayment.New(servicePayment)

	c.POST("/transactions/:transaksi_id/payments",handlerPayment.Add,middlewares.JWTMiddleware())
	c.POST("/notification/handling",handlerPayment.Notification)

	dataImage := DataImage.New(db)
	serviceImage := ServiseImage.New(dataImage)
	handlerImage := HandlerImage.New(serviceImage)

	c.POST("/products/:product_id/uplouds",handlerImage.AddImage)
	c.GET("/images/:image_id",handlerImage.GetById)
	c.GET("/images",handlerImage.GetAll)

	dataTansactionPayment := DataTransactionPayment.New(db)
	serviceTansactionPayment := ServiseTransactionPayment.New(dataTansactionPayment)
	handlerTansactionPayment := HandlerTransactionPayment.New(serviceTansactionPayment)

	c.POST("/transaction/payment",handlerTansactionPayment.Add)
}