package wallet

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *Wallet) WalletInfo(c *gin.Context) {
	user := GetUserFromSession(c)
	if user.Wallet.ID == 0 {
		c.JSON(http.StatusOK, "You dont have a wallet yet. Create one with /wallet/create")
	} else if len(user.Wallet.Card) == 0 {
		c.JSON(http.StatusOK, "No cards in wallet yet")
	} else {
		c.JSON(http.StatusOK, user.Wallet.Card)
	}
}

func (w *Wallet) WalletBalance(c *gin.Context) {
	user := GetUserFromSession(c)
	if len(user.Wallet.Card) == 0 {
		c.JSON(http.StatusOK, "No cards in wallet yet")
	} else {
		totalBalance := 0.0
		for i := range user.Wallet.Card {
			totalBalance += user.Wallet.Card[i].CardBalance
		}
		c.JSON(http.StatusOK, fmt.Sprintf("Total balance: %f", totalBalance))
	}
}
