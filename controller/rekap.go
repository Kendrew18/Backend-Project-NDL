package controllers

import (
	"Backend-Project-NDL/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func ReadRekap(c echo.Context) error {
	page := c.FormValue("page")

	pg, _ := strconv.Atoi(page)

	result, err := models.Read_Rekap(pg)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
