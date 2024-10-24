package middlewares

import (
	"net/http"
	"strings"

	"github.com/soulter/tickstats/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/types"
)

func JWTAuthMiddleware(directlyAbort bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		abortHandler := func() {
			if directlyAbort {
				c.JSON(http.StatusUnauthorized, types.NotAuthorizedResult)
				c.Abort()
			} else {
				c.Set("isAuthorized", false)
			}
		}

		authHeader := c.GetHeader("Authorization")
		tokenString := ""
		if authHeader == "" {
			// 从Cookie中获取
			authHeader = c.GetHeader("Cookie")
			if authHeader == "" {
				abortHandler()
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
			abortHandler()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
			c.Set("isAuthorized", true)
			c.Next()
		} else {
			abortHandler()
			return
		}
	}
}
