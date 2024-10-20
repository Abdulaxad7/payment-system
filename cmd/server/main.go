package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"payment-system/internal/auth"
	"payment-system/internal/auth/middleware"
	"payment-system/internal/payment"
	"payment-system/internal/wallet"
	"payment-system/pkg/database"
)

func init() {
	_ = godotenv.Load()
	database.Initialize()
}

func main() {
	s := &Server{}
	a, w, p := StructInitialize()
	s = s.init().InitializeRedis()

	RegisterRouters(s, a, w, p)

	s.Run(os.Getenv("PORT"))
}

func StructInitialize() (*auth.Auth, *wallet.Wallet, *payment.Payment) {
	return &auth.Auth{}, &wallet.Wallet{}, &payment.Payment{}
}

func RegisterRouters(s *Server, a *auth.Auth, w *wallet.Wallet, p *payment.Payment) {
	s.Router.POST("/signup", a.Signup)
	s.Router.POST("/login/", a.Login)
	s.Router.GET("/verify", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "We have sent verification code to your email",
		})
	})
	s.Router.POST("/pay",
		middleware.RequestAuthentication,
		p.NewPayment,
	)
	s.Router.GET("/pay/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "We have sent verification code to your email",
		})
	})
	s.Router.GET("/wallet",
		middleware.RequestAuthentication,
		w.WalletInfo)
	s.Router.POST("/wallet/create",
		middleware.RequestAuthentication,
		w.CreateWallet)
	s.Router.POST("/wallet/create/card",
		middleware.RequestAuthentication,
		w.AddCardToWallet)
	s.Router.DELETE("/wallet/delete",
		middleware.RequestAuthentication,
		w.DeleteWallet)
	s.Router.DELETE("/wallet/delete/card",
		middleware.RequestAuthentication,
		w.DeleteCardFromWallet)
	s.Router.POST("/verify",
		a.VerifyUser)
	s.Router.GET("/logout",
		a.Logout)
	s.Router.GET("/wallet/cards", w.GetCardList)
}
