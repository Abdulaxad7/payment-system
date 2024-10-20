package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var GORM *gorm.DB

func Initialize() {
	d, err := Connect()
	if err != nil {
		panic(err)
	}
	GORM = d
}
func Connect() (*gorm.DB, error) {
	db, err := getConnection()
	if err != nil ||
		UserInit(db) != nil ||
		TransactionInit(db) != nil ||
		CardInit(db) != nil ||
		WalletInit(db) != nil {
		return db, err
	}
	return db, nil
}

func getConnection() (*gorm.DB, error) {
	d, err := gorm.Open(mysql.Open(query()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return d, nil
}

func query() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_NAME"),
	)
}

func UserInit(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
func TransactionInit(db *gorm.DB) error {
	return db.AutoMigrate(&Transaction{})
}

func CardInit(db *gorm.DB) error {
	return db.AutoMigrate(&Card{})
}
func WalletInit(db *gorm.DB) error {
	return db.AutoMigrate(&Wallet{})
}
