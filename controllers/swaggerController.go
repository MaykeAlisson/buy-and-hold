package controllers

import (
	"github.com/gin-gonic/gin"
)

func Swagger(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Meu swagger"})

}
