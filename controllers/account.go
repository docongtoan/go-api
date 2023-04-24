package controllers

import (
	models "goserverapi/models"
	response "goserverapi/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertAccount(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err.Error(), response.EmptyObj{}))
		return
	}
	if err := models.ValidateAccount(account); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to validate account", err.Error(), response.EmptyObj{}))
		return
	}
	checkExist, err := models.CheckExistAccount(0, account.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to check account existence", err.Error(), response.EmptyObj{}))
		return
	}
	if checkExist {
		c.JSON(http.StatusConflict, response.NewErrorResponse("Account code already exists", "", response.EmptyObj{}))
		return
	}
	result, err := models.InsertAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to insert account", err.Error(), response.EmptyObj{}))
		return
	}
	c.JSON(http.StatusCreated, response.NewResponse(true, "Account inserted successfully", result))
}
