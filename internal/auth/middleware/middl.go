package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-system/pkg/database"
	"time"
)

func RequestAuthentication(c *gin.Context) {
	var err error
	tokenString, err := c.Cookie("_user_token")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte("SECRET_KEY"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user database.User
		database.GORM.First(&user, claims["id"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		session := sessions.Default(c)
		if session.Get("userID") == nil {
			session.Set("userID", user.ID)
			if err = session.Save(); err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
