package controllers

import (
	"Backend-Project-NDL/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func ReadEXCEL(c echo.Context) error {
	result, err := models.Read_EXCEL(c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func InputNDL(c echo.Context) error {
	status := c.FormValue("status")

	result, err := models.Input_NDL(status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadNDL(c echo.Context) error {
	page := c.FormValue("page")

	pg, _ := strconv.Atoi(page)

	result, err := models.Read_NDL(pg)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStock(c echo.Context) error {
	kode_stock := c.FormValue("kode_stock")
	nama_barang := c.FormValue("nama_barang")
	jumlah_barang := c.FormValue("jumlah_barang")
	harga_barang := c.FormValue("harga_barang")
	satuan_barang := c.FormValue("satuan_barang")

	jb, _ := strconv.ParseFloat(jumlah_barang, 64)

	hb, _ := strconv.Atoi(harga_barang)

	result, err := models.Update_Stock(kode_stock, nama_barang, jb, hb, satuan_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
