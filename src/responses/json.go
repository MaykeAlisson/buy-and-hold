package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func ValidError(c *gin.Context, err error) {
	
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

}