package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/models"
	"github.com/maykealisson/buy-and-hold/src/database"
	"github.com/maykealisson/buy-and-hold/src/responses"
)

func CreateUser(c *gin.Context) {

	var json dtos.UserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := json.Validate()
	if err != nil {
		responses.ValidError(c, err)
		return
	}

	user := models.User{
		Name: json.Name,
		Email: json.Email,
		Password: json.Password,
	}

	userCreated, err := user.SaveUser(database.DB)

	if err != nil {

		//formattedError := formaterror.FormatError(err.Error())

		//responses.ERROR(w, http.StatusInternalServerError, formattedError)
		responses.ValidError(c, err)
		return
	}

	responses.Response(c, http.StatusCreated, userCreated)

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
