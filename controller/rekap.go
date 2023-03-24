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

func UpdateStatusRekap(c echo.Context) error {
	ws_no := c.FormValue("ws_no")
	status_rekap := c.FormValue("status_rekap")

	sr, _ := strconv.Atoi(status_rekap)

	result, err := models.Update_Status_Rekap(ws_no, sr)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
