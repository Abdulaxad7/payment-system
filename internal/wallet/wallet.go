package wallet

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-system/internal"
	"payment-system/pkg/database"
)

func (w *Wallet) GetCardList(c *gin.Context) {
	var cards []database.Card
	for _, v := range internal.User.Wallet.Card {
		cards = append(cards, v)
	}
	c.JSON(http.StatusOK, cards)
}

func getSessionID(c *gin.Context) uint {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
	return userID.(uint)
}

func GetUserFromSession(c *gin.Context) database.User {
	var user database.User
	id := getSessionID(c)
	if err := database.GORM.Preload("Wallet").First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
	}
	return user
}
