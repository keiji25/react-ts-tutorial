package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	err = godotenv.Load("../.env")
	if err != nil {
		log.Println(".envの読み込みに失敗しました。")
		os.Exit(1)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_NAME)
	count := 5
	for count > 0 {
		if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			count--
			if count == 0 {
				log.Println("DBに接続できませんでした。")
				os.Exit(1)
			}
			time.Sleep(time.Second * 2)
			continue
		}
		log.Println("DB接続完了")
		break
	}

	DB.AutoMigrate(&Board{})
}
