package controllers

import (
	models "goserverapi/models"
	response "goserverapi/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetListUnit(c *gin.Context) {
	result, err := models.GetListUnit()

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to get list unit", err.Error(), result))
	} else {
		c.JSON(http.StatusOK, response.NewResponse(true, "Unit get list data successfully", result))
	}
}

func GetRowUnit(c *gin.Context) {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request param id", err.Error(), response.EmptyObj{}))
	} else {
		unit_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		result, err := models.GetRowUnit(unit_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to get row unit", err.Error(), response.EmptyObj{}))
		} else {
			c.JSON(http.StatusOK, response.NewResponse(true, "Unit get row successfully", result))
		}
	}
}

func InsertUnit(c *gin.Context) {
	var unit models.Unit

	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err.Error(), response.EmptyObj{}))
		return
	}
	if err := models.ValidateUnit(unit); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to validate unit", err.Error(), response.EmptyObj{}))
		return
	}
	checkExist, err := models.CheckExistUnit(0, unit.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to check unit existence", err.Error(), response.EmptyObj{}))
		return
	}
	if checkExist {
		c.JSON(http.StatusConflict, response.NewErrorResponse("Unit code already exists", "", response.EmptyObj{}))
		return
	}
	result, err := models.InsertUnit(unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to insert unit", err.Error(), response.EmptyObj{}))
		return
	}
	c.JSON(http.StatusCreated, response.NewResponse(true, "Unit inserted successfully", result))
}

func UpdateUnit(c *gin.Context) {
	var unit models.Unit

	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request param id", err.Error(), response.EmptyObj{}))
		return
	}
	unit_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err.Error(), response.EmptyObj{}))
		return
	}
	if err := models.ValidateUnit(unit); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to validate unit", err.Error(), response.EmptyObj{}))
		return
	}
	checkExist, err := models.CheckExistUnit(unit_id, unit.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to check unit existence", err.Error(), response.EmptyObj{}))
		return
	}
	if checkExist {
		c.JSON(http.StatusConflict, response.NewErrorResponse("Unit code already exists", "", response.EmptyObj{}))
		return
	}
	result, err := models.UpdateUnit(unit, unit_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to update unit", err.Error(), response.EmptyObj{}))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(true, "Unit updated successfully", result))
}

func DeleteUnit(c *gin.Context) {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request param id", err.Error(), response.EmptyObj{}))
		return
	}
	unit_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	result, err := models.DeleteUnit(unit_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to delete unit", err.Error(), response.EmptyObj{}))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(true, "Unit deleted successfully", result))
}
