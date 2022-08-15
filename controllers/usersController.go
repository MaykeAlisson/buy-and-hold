package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/dtos"
)

func CreateUser(c *gin.Context) {

	var json dtos.UserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := json.Validate()
	if err != nil {
		// responses.ERROR(w, http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusCreated, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, json)

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "update usuario" + id})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "delete usuario" + id})

}
