package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/responses"
	"github.com/maykealisson/buy-and-hold/src/services"
)

func CreateUser(c *gin.Context) {

	var dto dtos.UserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate()
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	userAcess, err := services.CreateUser(dto)
	if err != nil {

		//formattedError := formaterror.FormatError(err.Error())

		//responses.ERROR(w, http.StatusInternalServerError, formattedError)
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusCreated, userAcess)

	//c.JSON(http.StatusCreated, json)

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "update usuario" + id})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "delete usuario" + id})

}
