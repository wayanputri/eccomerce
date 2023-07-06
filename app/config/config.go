package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	SECRET_JWT =""
)

type AppConfig struct{
	DB_USERNAME 		string
	DB_PASSWORD 		string
	DB_HOSTNAME 		string
	DB_PORT 			int 
	DB_NAME 			string
	jwtKey 				string
	KEY_SERVER_MIDTRANS string
	KEY_API 			string
	KEY_API_SECRET 		string
	CLOUD_NAME			string
}

func InitConfig() *AppConfig{
	return ReadEnv()
}

func ReadEnv() *AppConfig{
	app := AppConfig{}
	isRead := true

	if val,found := os.LookupEnv("JWT_KEY");found{
		app.jwtKey = val
		isRead = false
	}

	if val,found := os.LookupEnv("DBUSER");found{
		app.DB_USERNAME = val
		isRead = false
	}
	if val,found := os.LookupEnv("DBPASS");found{
		app.DB_PASSWORD = val
		isRead = false
	}
	if val,found := os.LookupEnv("DBHOST");found{
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val,found := os.LookupEnv("DBPORT");found{
		cnv,_:= strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val,found := os.LookupEnv("DBNAME");found{
		app.DB_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("KEY_SERVER_MIDTRANS"); found {
		app.KEY_SERVER_MIDTRANS = val
		isRead = false
	}

	if val, found := os.LookupEnv("KEY_API"); found {
		app.KEY_API = val
		isRead = false
	}
	if val, found := os.LookupEnv("KEY_API_SECRET"); found {
		app.KEY_API_SECRET = val
		isRead = false
	}	
	if val, found := os.LookupEnv("CLOUD_NAME"); found {
		app.CLOUD_NAME = val
		isRead = false
	}	

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil{
			log.Println("error read config: ",err.Error())
			return nil
		}

		app.jwtKey = viper.Get("JWT_KEY").(string)
		app.DB_USERNAME = viper.Get("DBUSER").(string)
		app.DB_PASSWORD = viper.Get("DBPASS").(string)
		app.DB_HOSTNAME = viper.Get("DBHOST").(string)
		app.DB_PORT,_ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DB_NAME = viper.Get("DBNAME").(string)
		app.KEY_SERVER_MIDTRANS = viper.Get("KEY_SERVER_MIDTRANS").(string)
		app.KEY_API = viper.Get("KEY_API").(string)
		app.KEY_API_SECRET = viper.Get("KEY_API_SECRET").(string)
		app.CLOUD_NAME = viper.Get("CLOUD_NAME").(string)
	}
	SECRET_JWT = app.jwtKey
	return &app
}