package controllers

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/responses"
	"github.com/maykealisson/buy-and-hold/src/services"
)

func GetByMonth(c *gin.Context) {
	var err error
	month, errorFormt := strconv.ParseInt(c.Param("month"))
	if errorFormt != nil || month <= 0 || month > 12 {
		c.JSON(400, gin.H{"message": "month error format"})
		return
	}

    // pega id do usuario do token
	launchs, err := services.LaunchService().FindByMonth(userId, month)
	responses.Response(c, http.StatusOk, launchs)

}

func GetByAssert(c *gin.Context) {

	var err error
	assertId, errorFormt := strconv.ParseUint(c.Param("assertId"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "assertId error format"})
		return
	}

	// pega id do usuario do token
	launchs, err := services.LaunchService().FindByAssert(userId, assertId)
	responses.Response(c, http.StatusOk, launchs)

}

func CreateLaunch(c *gin.Context) {

	var err error
	assertId, errorFormt := strconv.ParseUint(c.Param("assertId"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "assertId error format"})
		return
	}
	var dto dtos.LaunchDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	// pega id do usuario do token
	launch, err := services.LaunchService().Create(userId, assertId, dto)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusCreated, launch)

}

func UpdateLaunch(c *gin.Context) {

	var err error
	assertId, errorFormt := strconv.ParseUint(c.Param("assertId"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "assertId error format"})
		return
	}

	launchId, err := strconv.ParseUint(c.Param("launchId"), 2, 32)
	if err != nil {
		c.JSON(400, gin.H{"message": "launchId error format"})
		return
	}

	var dto dtos.LaunchDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("update")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	// pega id do usuario do token
	launch, err := services.LaunchService().Update(userId, assertId, launchId, dto)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusOk, launch)

}

func DeleteLaunch(c *gin.Context) {

	var err error
	assertId, errorFormt := strconv.ParseUint(c.Param("assertId"), 2, 32)
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": "assertId error format"})
		return
	}

	launchId, err := strconv.ParseUint(c.Param("launchId"), 2, 32)
	if err != nil {
		c.JSON(400, gin.H{"message": "launchId error format"})
		return
	}

	var dto dtos.LaunchDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("update")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	// pega id do usuario do token
    err := services.LaunchService().Delete(userId, assertId, launchId)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusOk)

}
