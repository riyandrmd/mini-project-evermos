package config

import (
	"fmt"
	"os"
	"toko-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database:" + err.Error())
	}

	DB = db

	err = db.AutoMigrate(
		&model.User{},
		&model.Toko{},
		&model.Alamat{},
		&model.Category{},
		&model.Product{},
		&model.ProductImage{},
		&model.Trx{},
		&model.TrxDetail{},
		&model.LogProduk{},
	)
	if err != nil {
		panic("Failed to migrate table: " + err.Error())
	}
}
