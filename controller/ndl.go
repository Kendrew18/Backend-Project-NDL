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
	order_ndl := c.FormValue("order_ndl")
	w_s_order := c.FormValue("w_s_order")
	width := c.FormValue("width")
	lenght := c.FormValue("lenght")
	gusset := c.FormValue("gusset")
	prod_size := c.FormValue("prod_size")
	w := c.FormValue("w")
	c_ndl := c.FormValue("c_ndl")
	color := c.FormValue("color")
	nama_layer := c.FormValue("layer")
	layer_detail := c.FormValue("layer_detail")
	width_layer := c.FormValue("width_layer")
	rm := c.FormValue("rm")
	diff := c.FormValue("diff")
	lyr := c.FormValue("lyr")
	ink := c.FormValue("ink")
	adh := c.FormValue("adh")
	total := c.FormValue("total")

	dr, _ := strconv.Atoi(durasi)
	r_ndl, _ := strconv.Atoi(repeat_ndl)
	up_ndl, _ := strconv.Atoi(up)
	t_ndl, _ := strconv.Atoi(toleransi)
	w_s_ndl, _ := strconv.ParseFloat(w_s_order, 64)
	width_ndl, _ := strconv.ParseFloat(width, 64)
	lenght_ndl, _ := strconv.ParseFloat(lenght, 64)
	gusset_ndl, _ := strconv.ParseFloat(gusset, 64)
	p_s_ndl, _ := strconv.ParseFloat(prod_size, 64)
	w_ndl, _ := strconv.ParseFloat(w, 64)
	c_ndl_flt, _ := strconv.ParseFloat(c_ndl, 64)
	color_ndl, _ := strconv.Atoi(color)
	ink_ndl, _ := strconv.ParseFloat(ink, 64)

	result, err := models.Input_NDL(ws_no, tambah_data_tanggal, customer_delivery_date,
		job_done, dr, analyzer_version, order_status, cylinder_status, gol, cust, item_name,
		model, up_ndl, r_ndl, t_ndl, order_ndl, w_s_ndl, width_ndl, lenght_ndl, gusset_ndl, p_s_ndl, w_ndl,
		c_ndl_flt, color_ndl, nama_layer, layer_detail, width_layer, rm, diff, lyr, ink_ndl, adh, total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

/*func ReadStock(c echo.Context) error {
	result, err := models.Read_Stock()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}*/

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
