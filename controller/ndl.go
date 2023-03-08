package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-NDL/models"
	"strconv"
)

func InputStock(c echo.Context) error {
	ws_no := c.FormValue("ws_no")
	tambah_data_tanggal := c.FormValue("tambah_data_tanggal")
	customer_delivery_date := c.FormValue("customer_delivery_date")
	job_done := c.FormValue("job_done")
	durasi := c.FormValue("durasi")
	analyzer_version := c.FormValue("analyzer_version")
	order_status := c.FormValue("order_status")
	cylinder_status := c.FormValue("tambah_data_tanggal")
	gol := c.FormValue("customer_delivery_date")
	cust := c.FormValue("job_done")
	item_name := c.FormValue("durasi")
	model := c.FormValue("analyzer_version")
	up := c.FormValue("up")
	repeat_ndl := c.FormValue("repeat_ndl")
	toleransi := c.FormValue("toleransi")
	order_ndl := c.FormValue("job_done")
	w_s_order := c.FormValue("durasi")
	width := c.FormValue("analyzer_version")
	lenght := c.FormValue("order_status")
	gusset := c.FormValue("tambah_data_tanggal")
	prod_size := c.FormValue("customer_delivery_date")
	w := c.FormValue("job_done")
	c_ndl := c.FormValue("durasi")
	color := c.FormValue("analyzer_version")
	layer := c.FormValue("up")

	result, err := models.Input_NDL(ws_no, tambah_data_tanggal, customer_delivery_date,
		job_done, durasi, analyzer_version, order_status, cylinder_status, gol, cust, item_name,
		model, up, repeat_ndl, toleransi, order_ndl, w_s_order, width, lenght, gusset, prod_size, w,
		c, color, layer)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadStock(c echo.Context) error {
	result, err := models.Read_Stock()

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
