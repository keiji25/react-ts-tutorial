package model

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	err = godotenv.Load("../.env")
	if err != nil {
		fmt.Println(".envの読み込みに失敗しました。")
		os.Exit(1)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/db_name?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST)
	count := 5
	for count > 0 {
		if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			time.Sleep(time.Second * 2)
			count--
			if count == 0 {
				fmt.Println("DBに接続できませんでした。")
				os.Exit(1)
			}
			continue
		}
		break
	}
}
