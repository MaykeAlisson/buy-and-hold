package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetByMonth(c *gin.Context) {

	month := c.Param("month")

	c.JSON(http.StatusCreated, gin.H{"message": "buscar por mes " + month})

}

func GetByAssert(c *gin.Context) {

	assertId := c.Param("assertId")

	c.JSON(http.StatusOK, gin.H{"message": "busca lancamento por assert " + assertId})

}

func CreateLaunch(c *gin.Context) {

	assertId := c.Param("assertId")

	c.JSON(http.StatusOK, gin.H{"message": "Cria launch para assertId " + assertId})

}

func UpdateLaunch(c *gin.Context) {

	assertId := c.Param("assertId")
	launchId := c.Param("launchId")

	c.JSON(http.StatusOK, gin.H{"message": "update launch " + assertId + launchId})

}

func DeleteLaunch(c *gin.Context) {

	assertId := c.Param("assertId")
	launchId := c.Param("launchId")

	c.JSON(http.StatusOK, gin.H{"message": "delete launch " + assertId + launchId})

}
