package controllers

import (
	"net/http"
	"strconv"

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

	userAcess, err := services.UserService().CreateUser(dto)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusCreated, userAcess)

}

func UpdateUser(c *gin.Context) {

	id, errorFormt := strconv.ParseUint(c.Param("id"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "id error format"})
		return
	}

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

	erroUpdate := services.UserService().UpdateUser(uint32(id), dto)
	if erroUpdate != nil {
		responses.BusinessException(c, erroUpdate)
		return
	}
	c.JSON(http.StatusOK, nil)

}

func DeleteUser(c *gin.Context) {

	id, errorFormt := strconv.ParseUint(c.Param("id"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "id error format"})
		return
	}

	// verifica se id nao e null e confert em int
	// verifica se e o mesmo id que esta no token

	err := services.UserService().DeleteUser(uint32(id))
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)

}
