package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/providers"
)

func SetMiddlewareAuthentication(next gin.Context) c *gin.Context {
	return func(w http.ResponseWriter, r *http.Request) {
		err := providers.TokenValid(r)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized "})
			return
		}
		next(c *gin.Context)
	}
}