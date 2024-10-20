package payment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"payment-system/internal"
	"payment-system/internal/mails"
	"payment-system/internal/wallet"
	"payment-system/pkg/database"
)

func (p *Payment) NewPayment(c *gin.Context) {
	internal.User = wallet.GetUserFromSession(c)
	fmt.Println(internal.User)
	var card struct {
		From     string  `json:"from"`
		Receiver string  `json:"receiver"`
		Amount   float64 `json:"amount"`
	}
	if err := c.Bind(&card); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to bind")
		p.log.Error(err.Error())
		return
	}
	var ca database.Card
	for _, v := range internal.User.Wallet.Card {
		if card.From == v.CardNumber {
			if card.Amount <= v.CardBalance {
				ca = v
			}
		}
	}
	if ca.CardNumber == "" {
		c.JSON(http.StatusNotFound, "Card not found: "+card.From)
		return
	} else {
		if p.Verify(c, card.From, card.Receiver, card.Amount) {
			p.Trans(c, ca, card.Amount)
		}
	}
}

func (p *Payment) Trans(c *gin.Context,
	card database.Card,
	amount float64) {
	if err := database.GORM.Model(&card).
		Where("card_number = ?", card.CardNumber).
		Update("card_balance", card.CardBalance-amount).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, "Failed to finish transaction")
		p.log.Error(err.Error())
		return
	} else {
		c.JSON(http.StatusOK, "Payment successful")
	}
}

func (p *Payment) Verify(c *gin.Context, from, receiver string, amount float64) bool {
	m := mails.Mail{}
	gen := p.GenerateCode()
	if err := m.SendEmail(internal.User.Email,
		m.TransactionMail(from, receiver, gen, amount)); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to sending mail")
		return false
	}
	var code struct {
		pin int
	}
	_ = c.Bind(&code)
	if code.pin == gen {
		return true
	}
	return false
}

func (p *Payment) GenerateCode() int {
	return rand.Intn(9999)
}
