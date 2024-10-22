package middlewares

import (
	"net/http"
	"strings"

	"github.com/soulter/tickstats/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/types"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := ""
		if authHeader == "" {
			// 从Cookie中获取
			authHeader = c.GetHeader("Cookie")
			if authHeader == "" {
				c.JSON(http.StatusUnauthorized, types.NotAuthorizedResult)
				c.Abort()
				return
			}
			cookie := strings.Split(authHeader, ";")
			for _, v := range cookie {
				if strings.Contains(v, "token=") {
					tokenString = strings.Split(v, "=")[1]
					break
				}
			}
		} else {
			tokenString = strings.Split(authHeader, "Bearer ")[1]
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.JWTSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, types.NotAuthorizedResult)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, types.NotAuthorizedResult)
			c.Abort()
			return
		}
	}
}
