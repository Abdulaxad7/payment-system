package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"payment-system/pkg/database"
)

func (auth *Auth) Login(c *gin.Context) {
	var model struct {
		Email    string
		Password string
	}
	if c.Bind(&model) != nil {
		auth.errHandler(c, "Invalid body")
		return
	}
	var u database.User
	database.GORM.Where("email = ?", model.Email).First(&u)
	if u.ID == 0 {
		auth.errHandler(c, "Invalid username or password")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(model.Password)) != nil {
		auth.errHandler(c, "Invalid username or password")
		return
	}
	if auth.generator(c) != nil {
		auth.errHandler(c, "Failed to create token")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func newCookie(c *gin.Context, tokenString string) {
	c.SetCookie("_user_token", tokenString, 3600, "/", "", false, true)
}
