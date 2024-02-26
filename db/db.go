package db

import (
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
)

func GetConnection() (*gorm.DB, error){
	
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	endpoint := os.Getenv("endpoint")

	connString := DB_USER+":"+DB_PASS+"@tcp("+endpoint+")/"+DB_NAME+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	log.Println("========= Connecting to AWS-RDS Mysql✨ ========= ")
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("========= Connected to AWS-RDS Mysql✨ ========= ")

	return db,nil
}