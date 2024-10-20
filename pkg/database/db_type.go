package database

import (
	"gorm.io/gorm"
	"time"
)

type USER User
type DB struct {
	*gorm.DB
	DBPublicHost string
	DBUser       string
	DBAddr       string
	DBPassword   string
	DBName       string
}
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Wallet    Wallet    `json:"wallet" gorm:"foreignKey:ID"`
}
type Wallet struct {
	ID   uint   `json:"ID" gorm:"primaryKey"`
	Card []Card `json:"cards" gorm:"foreignKey:ID"`
}

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	Receiver  string    `gorm:"not null"`
	Amount    float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"autoCreateTime"`
}
type Card struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"ID"`
	WalletID     uint          `json:"walletID" gorm:"not null"` // Foreign key for Wallet
	CardType     string        `json:"cardType" gorm:"not null"`
	CardHolder   string        `json:"cardHolder"`
	CardNumber   string        `json:"cardNumber" gorm:"unique"`
	CardThruDate string        `json:"cardThruDate" gorm:"not null"`
	CardCVV      uint          `json:"cardCVV" gorm:"not null"`
	CardBalance  float64       `json:"cardBalance" gorm:"not null"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:ID"`
}
