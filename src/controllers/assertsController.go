package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/providers"
	"github.com/maykealisson/buy-and-hold/src/responses"
	"github.com/maykealisson/buy-and-hold/src/services"
)

func GetAssertBy(c *gin.Context) {

	name := c.Query("name")

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}

	if name == "" {
		assets, err := services.AssertService().FindAllByUser(userId)
		if err != nil {
			responses.BusinessException(c, err)
			return
		}

		responses.Response(c, http.StatusOK, assets)
		return
	}

	assets, err := services.AssertService().FindByName(userId, strings.ToUpper(name))
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusOK, assets)

}

func CreateAssert(c *gin.Context) {

	var dto dtos.AssertDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}
	assert, err := services.AssertService().CreateAssert(userId, dto)

	responses.Response(c, http.StatusOK, assert)

}

func UpdateAssert(c *gin.Context) {

	assertId, errorFormt := strconv.Atoi(c.Param("id"))
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": errorFormt.Error()})
		return
	}

	var dto dtos.AssertDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err := dto.Validate("update")
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}

	assert, err := services.AssertService().Update(uint32(assertId), userId, dto)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusCreated, assert)

}

func DeleteAssert(c *gin.Context) {

	assertId, errorFormt := strconv.Atoi(c.Param("id"))
	if errorFormt != nil {
		c.JSON(400, gin.H{"message": errorFormt.Error()})
		return
	}

	userId, errUserId := providers.JwtProvider().GetUserId(c)
	if errUserId != nil || userId == 0 {
		responses.BusinessException(c, errUserId)
		return
	}

	err := services.AssertService().Delete(uint32(assertId), userId)
	if err != nil {
		responses.BusinessException(c, err)
		return
	}

	responses.Response(c, http.StatusOK, nil)

}
