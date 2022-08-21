package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func BusinessException(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func Exception(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}
