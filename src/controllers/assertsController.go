package controllers

import (
	"net/http"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/responses"
	"github.com/maykealisson/buy-and-hold/src/services"
)

func GetAssertBy(c *gin.Context) {

	name := c.Query("name")

	// verifica se name não e null  
	// pega o id do usuario no token 
	// passa para o service o id e o name 
	assets, err := services.AssertService().FindByName(uint32(userId), name)

	c.JSON(http.StatusOK, gin.H{"message": "busca todos assert por " + name})

}

func CreateAssert(c *gin.Context) {

	var err error
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

	// extrair id do usuario to token
	assert, err := services.AssertService().CreateAssert(uint32(userId), dto)

	c.JSON(http.StatusOK, gin.H{"message": "Cria assert"})

}

func UpdateAssert(c *gin.Context) {

	var err error
	assertId, errorFormt := strconv.ParseUint(c.Param("assertId"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "id error format"})
		return
	}
	
	var dto dtos.AssertDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("update")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	// extrair id do usuario to token

	assert, err := services.AssertService().CreateAssert(uint32(assertId), userId, dto)
	c.JSON(http.StatusOK, gin.H{"message": "update por id " + assertId})

}

func DeleteAssert(c *gin.Context) {
	
	var err error
	assertId, errorFormt := strconv.ParseUint(c.Param("assertId"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "id error format"})
		return
	}
	

	// extrair id do usuario to token

	err := services.AssertService().Delete(uint32(assertId), userId, dto)

	c.JSON(http.StatusOK, gin.H{"message": "delete assert id " + assertId})

}
