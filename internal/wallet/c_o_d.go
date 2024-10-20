package wallet

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-system/internal"
	"payment-system/pkg/database"
)

func (w *Wallet) CreateWallet(c *gin.Context) {
	var err error
	internal.User = GetUserFromSession(c)
	if internal.User.Wallet.ID != 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "You already have a wallet"})
		return
	}
	if err = database.GORM.Create(&internal.User.Wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "Failed to create wallet"})
		w.log.Error(err.Error())
		return
	} else {
		c.JSON(http.StatusOK,
			gin.H{"message": "Wallet created"})
	}
}

func (w *Wallet) AddCardToWallet(c *gin.Context) {
	var err error
	var i = 1
	internal.User = GetUserFromSession(c)
	if internal.User.Wallet.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You don't have a wallet",
		})
		return
	}
	card := database.Card{}
	if err = c.Bind(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind card",
		})
		w.log.Error(err.Error())
	}

	internal.User.Wallet.Card = append(internal.User.Wallet.Card, card)
	if err = database.GORM.Create(&internal.User.Wallet.Card).Error; err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("You already attached card [%v]", card.CardNumber)})
		//w.log.Error(err.Error())
		return
	} else {
		i++
		c.JSON(http.StatusOK, gin.H{"message": "Card created"})
	}
}

func (w *Wallet) DeleteWallet(c *gin.Context) {
	internal.User = GetUserFromSession(c)
	if internal.User.Wallet.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You don't have a wallet",
		})
		return
	} else if len(internal.User.Wallet.Card) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You have cards in your wallet",
		})
		return
	} else {
		if err := database.GORM.Delete(&internal.User.Wallet).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to delete wallet: ",
			})
			w.log.Error(err.Error())
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Wallet deleted",
			})
		}
	}
}

func (w *Wallet) DeleteCardFromWallet(c *gin.Context) {
	if internal.User.Wallet.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You don't have a wallet",
		})
		return
	} else {
		var card struct {
			CardNumber string `json:"card_number" binding:"required"`
		}
		if err := c.Bind(&card); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to bind card",
			})
			return
		}
		if err := database.GORM.Exec(
			fmt.Sprintf(
				"DELETE FROM cards WHERE card_number = '%s';",
				card.CardNumber,
			)).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to delete card",
			})
			w.log.Error(err.Error())
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Card deleted: " + card.CardNumber,
			})
		}
	}
}
