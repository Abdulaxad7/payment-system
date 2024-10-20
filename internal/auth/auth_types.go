package auth

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Authentications interface {
	Login(c *gin.Context)
	Signup(c *gin.Context)
	generator(c *gin.Context) error
	errHandler(c *gin.Context, message string)
	newCookie(c *gin.Context, tokenString string)
	VerifyUser(c *gin.Context)
	Logout(c *gin.Context)
}
type Auth struct {
	log slog.Logger
}
