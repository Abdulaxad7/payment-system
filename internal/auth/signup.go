package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"payment-system/internal/mails"
	"payment-system/pkg/database"
	"time"
)

var user database.User
var m mails.Mail

func (auth *Auth) Signup(c *gin.Context) {

	var err error
	var model struct {
		Name     string
		Email    string
		Password string
	}
	if c.Bind(&model) != nil {
		auth.errHandler(c, "Invalid body")
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(model.Password), 12)
	if err != nil {
		auth.errHandler(c, "Failed to create user")
		auth.log.Error(err.Error())
		return
	}
	user = database.User{
		Name:     model.Name,
		Email:    model.Email,
		Password: string(pass),
	}

	m.Code = m.GenerateCode()
	if err = m.SendEmail(model.Email, m.AuthMail(m.Code)); err != nil {
		auth.errHandler(c, "Failed to send email")
		auth.log.Error(err.Error())
		return
	}
	c.Redirect(http.StatusFound, "/verify")
}

func (auth *Auth) generator(c *gin.Context) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  1,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		auth.errHandler(c, "Failed to create token")
		auth.log.Error(err.Error())
	}
	newCookie(c, tokenString)
	return nil
}

func (auth *Auth) errHandler(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
	//auth.log.Error(err)
}

func (auth *Auth) VerifyUser(c *gin.Context) {
	var model struct {
		Password string
	}
	if c.Bind(&model) != nil {
		auth.errHandler(c, "Invalid body")
		return
	}
	if m.VerifyEmail(model.Password) {
		if auth.generator(c) != nil {
			auth.errHandler(c, "Failed to create token")
			return
		}
		if database.GORM.Create(&user).Error != nil {
			auth.errHandler(c, "Failed to create user")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Signup successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
	}
}
