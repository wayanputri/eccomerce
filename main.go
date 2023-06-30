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

// export JWT_KEY='ecommerce_secret_dong'
// export DBUSER='root'
// export DBPASS='wayan123'
// export DBHOST='127.0.0.1'
// export DBPORT='3306'
// export DBNAME='ecommerce'
// export KEY_SERVER_MIDTRANS='SB-Mid-server-0NrVq555rAslGxfgAKi_NqMn'
// export KEY_CLIENT_MIDTRANS='SB-Mid-client-mFPo7qREcM9p6x0y'