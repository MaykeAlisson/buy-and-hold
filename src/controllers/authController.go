package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/responses"
	"github.com/maykealisson/buy-and-hold/src/services"
)

func Auth(c *gin.Context) {

	var dto dtos.UserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("login")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	acess, err := services.AuthService().Longin(dto)
	if err != nil {
		responses.Exception(c, 401, err)
		return
	}

	responses.Response(c, http.StatusOK, acess)

}
