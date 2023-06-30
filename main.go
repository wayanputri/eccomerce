package main

import (
	"belajar/bareng/app/config"
	"belajar/bareng/app/database"
	"belajar/bareng/app/router"
	"belajar/bareng/features/payment/data"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	cfg := config.InitConfig()
	mysql := database.InitMysql(cfg)
	database.InitialMigration(mysql)
	e :=echo.New()
	
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	
	data.GoMitransts()
	router.InitRouter(e,mysql)
	e.Logger.Fatal(e.Start(":8080"))

}