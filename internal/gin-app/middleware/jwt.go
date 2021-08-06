package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"quick-go/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 0
		token := c.Query("token")

		if len(token) <= 0 {
			code = -1
		} else {
			_, err := util.ParseToken(token)

			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = -10
				default:
					code = -1
				}
			}
		}

		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "JWT auth failed",
			})

			c.Abort()
			return
		}
	}
}
