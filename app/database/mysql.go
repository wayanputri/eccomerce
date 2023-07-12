package database

import (
	"belajar/bareng/app/config"
	"belajar/bareng/features"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME,cfg.DB_PASSWORD,cfg.DB_HOSTNAME,cfg.DB_PORT,cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	return DB
}

func InitialMigration(db *gorm.DB){
	db.AutoMigrate(&features.User{},&features.Product{},&features.Image{},&features.Review{},&features.ReviewImages{},&features.TransactionPayment{},&features.Transaction{},&features.TransactionPayment{},&features.Payment{})	
}