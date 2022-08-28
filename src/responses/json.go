package responses

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handlerError struct {
	StatusCode  int    `json:"statusCode"`
	TimesTamp   string `json:"timestamp"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func BusinessException(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, handlerError{
		StatusCode:  http.StatusBadRequest,
		TimesTamp:   time.Now().String(),
		Message:     err.Error(),
		Description: c.FullPath(),
	})
}

func Exception(c *gin.Context, status int, err error) {
	c.JSON(status, handlerError{
		StatusCode:  status,
		TimesTamp:   time.Now().String(),
		Message:     err.Error(),
		Description: c.FullPath(),
	})
}
