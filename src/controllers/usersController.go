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

	err := dto.Validate("")
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

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var dto dtos.UserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("update")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	// verifica se id nao e null e confert em int 
	// verifica se e o mesmo id que esta no token
	
	erroUpdate := services.UpdateUser(id, dto)
	if erroUpdate != nil {
		responses.BusinessException(c, erroUpdate)
		return
	}
	c.JSON(http.StatusOK, nil)

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// verifica se id nao e null e confert em int 
	// verifica se e o mesmo id que esta no token

	err := services.DeleteUser(id)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)

}
