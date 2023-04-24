package controllers

import (
	models "goserverapi/models"
	response "goserverapi/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err.Error(), response.EmptyObj{}))
		return
	}

	auth := models.Auth{}

	auth.Email = input.Email

	auth.Password = input.Password

	if err := models.ValidateAuth(auth); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Invalid Email or Password", err.Error(), response.EmptyObj{}))
		return
	}

	result, err := models.LoginCheck(auth.Email, auth.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to checked login", err.Error(), response.EmptyObj{}))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(true, "Success login", result))
}
