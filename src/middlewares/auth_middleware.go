package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/providers"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer_schema = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(401, gin.H{"message": "required token"})
			ctx.Abort()
			return
		}

		token := header[len(Bearer_schema):]

		if !providers.JwtProvider().TokenValid(token) {
			ctx.JSON(401, gin.H{"message": "invalid token"})
			ctx.Abort()
			return
		}
	}
}
