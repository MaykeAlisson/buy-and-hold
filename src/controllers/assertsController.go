package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/responses"
)

func GetAssertBy(c *gin.Context) {

	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{"message": "busca todos assert por " + name})

}

func CreateAssert(c *gin.Context) {

	var dto dtos.AssertDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

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
