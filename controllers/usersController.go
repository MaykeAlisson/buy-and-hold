package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"message": "Cria usuario"})

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "update usuario" + id})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "delete usuario" + id})

}
