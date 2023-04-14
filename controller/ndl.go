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

func UpdateNDL(c echo.Context) error {
	ws_no := c.FormValue("ws_no")
	tambah_data_tanggal := c.FormValue("tambah_data_tanggal")
	customer_delivery_date := c.FormValue("customer_delivery_date")
	job_done := c.FormValue("job_done")
	order_status := c.FormValue("order_status")
	cylinder_status := c.FormValue("cylinder_status")
	gol := c.FormValue("gol")
	cust := c.FormValue("cust")
	item_name := c.FormValue("item_name")
	model := c.FormValue("model")
	up := c.FormValue("up")
	repeat_ndl := c.FormValue("repeat_ndl")
	toleransi := c.FormValue("toleransi")
	order_ndl := c.FormValue("order_ndl")
	layer := c.FormValue("layer")
	detail_layer := c.FormValue("detail_layer")

	up_i, _ := strconv.Atoi(up)
	repeat_i, _ := strconv.Atoi(repeat_ndl)
	toleransi_i, _ := strconv.Atoi(toleransi)
	order_i, _ := strconv.Atoi(order_ndl)

	result, err := models.Update_NDL(ws_no, tambah_data_tanggal, customer_delivery_date, job_done,
		order_status, cylinder_status, gol, cust, item_name, model, up_i,
		repeat_i, toleransi_i, order_i, layer, detail_layer)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PageNo(c echo.Context) error {
	result, err := models.Page()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadNDLwsno(c echo.Context) error {
	ws_no := c.FormValue("ws_no")

	result, err := models.Read_NDL_WSNO(ws_no)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
