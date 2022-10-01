package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/providers"
	"github.com/maykealisson/buy-and-hold/src/responses"
	"github.com/maykealisson/buy-and-hold/src/services"
)

func GetByMonth(c *gin.Context) {

	month, errorFormt := strconv.ParseInt(c.Param("month"), 10, 0)
	if errorFormt != nil || month <= 0 || month > 12 {
		c.JSON(400, gin.H{"message": "month error format"})
		return
	}

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}
	launchs, err := services.LaunchService().FindByMonth(userId, int(month))
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusOK, launchs)

}

func GetByAssert(c *gin.Context) {

	assertId, errorFormt := strconv.ParseInt(c.Param("assert"), 10, 0)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "assertId error format"})
		return
	}

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}
	launchs, err := services.LaunchService().FindByAssert(userId, uint32(assertId))
	if err != nil {
		responses.BusinessException(c, err)
		return
	}
	responses.Response(c, http.StatusOK, launchs)

}

func CreateLaunch(c *gin.Context) {

	var err error
	assertId, errorFormt := strconv.Atoi(c.Param("assert"))
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "assert error format"})
		return
	}
	var dto dtos.LauncheDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = dto.Validate("")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}
	launch, err := services.LaunchService().Create(userId, uint32(assertId), dto)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusCreated, launch)

}

func DeleteLaunch(c *gin.Context) {

	var err error

	launchId, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(400, gin.H{"message": "launchId error format"})
		return
	}

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}
	err = services.LaunchService().Delete(userId, uint32(launchId))
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusOK, nil)

}
