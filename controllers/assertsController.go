package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAssertBy(c *gin.Context) {

	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{"message": "busca todos assert por " + name})

}

func CreateAssert(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Cria assert"})

}

func UpdateAssert(c *gin.Context) {

	assertId := c.Param("assertId")

	c.JSON(http.StatusOK, gin.H{"message": "update por id " + assertId})

}

func DeleteAssert(c *gin.Context) {

	assertId := c.Param("assertId")

	c.JSON(http.StatusOK, gin.H{"message": "delete assert id " + assertId})

}
