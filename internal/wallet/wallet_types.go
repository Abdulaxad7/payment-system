package wallet

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Wallets interface {
	WalletInfo(*gin.Context)
	WalletBalance(*gin.Context)
	CreateWallet(*gin.Context)
	GetWallet(*gin.Context)
	AddCardToWallet(*gin.Context)
	DeleteCardFromWallet(*gin.Context)
	DeleteWallet(*gin.Context)
	GetCardList(*gin.Context)
}
type Wallet struct {
	log slog.Logger
}
