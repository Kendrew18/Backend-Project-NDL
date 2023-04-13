package controllers

import (
	"Backend-Project-NDL/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateTemplate(c echo.Context) error {
	internal_instruction_number := c.FormValue("internal_instruction_number")

	result, err := models.Update_Template(internal_instruction_number)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
