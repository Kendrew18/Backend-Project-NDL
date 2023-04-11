package controllers

import (
	"Backend-Project-NDL/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InputPOsupplier(c echo.Context) error {
	ws_no := c.FormValue("ws_no")
	layer := c.FormValue("layer")
	nama_po := c.FormValue("nama_po")
	tanggal := c.FormValue("tanggal")
	meter := c.FormValue("meter")
	kg := c.FormValue("kg")
	diff_pc := c.FormValue("diff_pc")

	result, err := models.Input_PO(ws_no, layer, nama_po, tanggal, meter, kg, diff_pc)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func ReadPO(c echo.Context) error {
	ws_no := c.FormValue("ws_no")
	layer := c.FormValue("layer")
	all_layer := c.FormValue("all_layer")

	lyr, _ := strconv.Atoi(layer)

	result, err := models.Read_PO(ws_no, lyr, all_layer)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func LyrPO(c echo.Context) error {
	ws_no := c.FormValue("ws_no")

	result, err := models.Lyr_PO(ws_no)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
