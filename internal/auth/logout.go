package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (auth *Auth) Logout(c *gin.Context) {
	if _, err := c.Cookie("_user_token"); err != nil {
		d := gin.Error{Err: err}
		fmt.Println(d)
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	c.SetCookie("_user_token", "", -1, "/l", "", false, true)

}
