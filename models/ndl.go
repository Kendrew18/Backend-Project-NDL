package models

import (
	"Backend-Project-NDL/db"
	str "Backend-Project-NDL/struct-all-ndl"
	"Backend-Project-NDL/tools"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Read_EXCEL(writer http.ResponseWriter, request *http.Request) (Response, error) {
	var res Response

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name : ", handler.Filename)
	fmt.Println("File Size : ", handler.Size)
	fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

	var tempFile *os.File
	path := ""
	if strings.Contains(handler.Filename, "xlsx") {
		path = "uploads/Read" + ".xlsx"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.xlsx")
	}

	if err != nil {
		return res, err
	}

	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return res, err2
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return res, err
	}

	fmt.Println("Success!!")
	fmt.Println(tempFile.Name())
	tempFile.Close()

	err = os.Rename(tempFile.Name(), path)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fmt.Println(tempFile.Name())

	xlsx, err := excelize.OpenFile("./uploads/Read.xlsx")
	if err != nil {
		return res, err
	}

	sheet1Name := "NDL"

	i := 2
	code := 0

	var Array_R_NDL []str.Read_NDL

	var arr_fl []string

	for code != 1 {
		var R_NDL str.Read_NDL
		WS_NO := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", i))
		fl := ""
		tmp := ""
		if WS_NO != "" {
			fl = fl + "{"
			vc := "A" + strconv.Itoa(i)
			hc := "C" + strconv.Itoa(i)

			style, _ := xlsx.NewStyle(`{"number_format":15}`)
			xlsx.SetCellStyle(sheet1Name, vc, hc, style)

			R_NDL.Tambah_data_tanggal = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i))
			tmp = "|" + R_NDL.Tambah_data_tanggal + "|"
			fl += tmp

			R_NDL.Customer_delivery_date = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i))
			tmp = "|" + R_NDL.Customer_delivery_date + "|"
			fl += tmp

			R_NDL.Job_done = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i))
			tmp = "|" + R_NDL.Job_done + "|"
			fl += tmp

			R_NDL.Durasi, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)))
			tmp = "|" + strconv.Itoa(R_NDL.Durasi) + "|"
			fl += tmp

			R_NDL.Analyzer_version = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i))
			tmp = "|" + R_NDL.Analyzer_version + "|"
			fl += tmp

			R_NDL.Order_status = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", i))
			tmp = "|" + R_NDL.Order_status + "|"
			fl += tmp

			R_NDL.Cylider_status = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i))
			tmp = "|" + R_NDL.Cylider_status + "|"
			fl += tmp

			R_NDL.Gol = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", i))
			tmp = "|" + R_NDL.Gol + "|"
			fl += tmp

			R_NDL.Ws_no = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", i))
			tmp = "|" + R_NDL.Ws_no + "|"
			fl += tmp

			R_NDL.Cust = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", i))
			tmp = "|" + R_NDL.Cust + "|"
			fl += tmp

			R_NDL.Item_name = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", i))
			tmp = "|" + R_NDL.Item_name + "|"
			fl += tmp

			R_NDL.Model = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", i))
			tmp = "|" + R_NDL.Model + "|"
			fl += tmp

			R_NDL.Up, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("M%d", i)))
			tmp = "|" + strconv.Itoa(R_NDL.Up) + "|"
			fl += tmp

			R_NDL.Repeat_ndl, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("N%d", i)))
			tmp = "|" + strconv.Itoa(R_NDL.Repeat_ndl) + "|"
			fl += tmp

			R_NDL.Toleransi, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("O%d", i)))
			tmp = "|" + strconv.Itoa(R_NDL.Toleransi) + "|"
			fl += tmp

			ord := "|["
			ORDER1, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("P%d", i)))
			tmp = "?" + strconv.Itoa(ORDER1) + "?"
			ord += tmp

			ORDER2, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Q%d", i)))
			tmp = "?" + strconv.Itoa(ORDER2) + "?"
			ord += tmp

			ord += "]|"
			fl += ord

			R_NDL.Order_ndl = append(R_NDL.Order_ndl, ORDER1)
			R_NDL.Order_ndl = append(R_NDL.Order_ndl, ORDER2)

			R_NDL.W_s_order, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("R%d", i)), 64)
			R_NDL.W_s_order = math.Round(R_NDL.W_s_order*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.W_s_order) + "|"
			fl += tmp

			R_NDL.Width, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", i)), 64)
			R_NDL.Width = math.Round(R_NDL.Width*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Width) + "|"
			fl += tmp

			R_NDL.Lenght, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("T%d", i)), 64)
			R_NDL.Lenght = math.Round(R_NDL.Lenght*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Lenght) + "|"
			fl += tmp

			R_NDL.Gusset, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("U%d", i)), 64)
			R_NDL.Gusset = math.Round(R_NDL.Gusset*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Gusset) + "|"
			fl += tmp

			R_NDL.Prod_size, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("V%d", i)), 64)
			R_NDL.Prod_size = math.Round(R_NDL.Prod_size*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Prod_size) + "|"
			fl += tmp

			R_NDL.W, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("W%d", i)), 64)
			R_NDL.W = math.Round(R_NDL.W*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.W) + "|"
			fl += tmp

			R_NDL.C, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("X%d", i)), 64)
			R_NDL.C = math.Round(R_NDL.C*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.C) + "|"
			fl += tmp

			R_NDL.Color, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Y%d", i)))
			tmp = "|" + strconv.Itoa(R_NDL.Color) + "|"
			fl += tmp

			R_NDL.Total_layer, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("CB%d", i)), 64)
			R_NDL.Total_layer = math.Round(R_NDL.Total_layer*100) / 100
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Total_layer) + "|"
			fl += tmp

			fmt.Println(tmp)

			//Layer1

			LD1 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Z%d", i))
			if LD1 != "" {
				lyr_fl := "|["

				NL := "1st Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)
				tmp = "?" + NL + "?"
				lyr_fl += tmp

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AA%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AB%d", i))
				LD4_dbl, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AC%d", i)), 64)
				LD4_dbl = math.Round(LD4_dbl*100) / 100
				LD4 := fmt.Sprintf("%.2f", LD4_dbl)

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				tmp = "?" + LD1 + "?"
				lyr_fl += tmp

				tmp = "?" + LD2 + "?"
				lyr_fl += tmp

				tmp = "?" + LD3 + "?"
				lyr_fl += tmp

				tmp = "?" + LD4 + "?"
				lyr_fl += tmp

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AD%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)
				tmp = "?" + fmt.Sprintf("%f", WD) + "?"
				lyr_fl += tmp

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AE%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)
				tmp = "?" + strconv.Itoa(RM) + "?"
				lyr_fl += tmp

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AF%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)
				tmp = "?" + fmt.Sprintf("%f", DIFF) + "?"
				lyr_fl += tmp

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BP%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)
				tmp = "?" + fmt.Sprintf("%f", LYR) + "?"
				lyr_fl += tmp

				R_NDL.Ink_layer, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BQ%d", i)), 64)
				R_NDL.Ink_layer = math.Round(R_NDL.Ink_layer*100) / 100
				tmp = "?" + fmt.Sprintf("%f", R_NDL.Ink_layer) + "?"
				lyr_fl += tmp

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BR%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)
				tmp = "?" + fmt.Sprintf("%f", ADH) + "?"
				lyr_fl += tmp

				lyr_fl += "]|"

				fl += lyr_fl
			} else {

				NL := "1st Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AA%d", i))
				LD3 := ""
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AC%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AD%d", i)), 64)
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AE%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AF%d", i)), 64)
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BP%d", i)), 64)
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				R_NDL.Ink_layer, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BQ%d", i)), 64)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BR%d", i)), 64)
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)
			}

			//Layer2
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AG%d", i))
			if LD1 != "" {
				lyr_fl := "|["

				NL := "2nd Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AH%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AI%d", i))
				LD4_dbl, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AJ%d", i)), 64)
				LD4_dbl = math.Round(LD4_dbl*100) / 100
				LD4 := fmt.Sprintf("%.2f", LD4_dbl)

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AK%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AL%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AM%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BS%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BT%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)

				tmp = "?" + NL + "?"
				lyr_fl += tmp

				tmp = "?" + LD1 + "?"
				lyr_fl += tmp

				tmp = "?" + LD2 + "?"
				lyr_fl += tmp

				tmp = "?" + LD3 + "?"
				lyr_fl += tmp

				tmp = "?" + LD4 + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", WD) + "?"
				lyr_fl += tmp

				tmp = "?" + strconv.Itoa(RM) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", DIFF) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", LYR) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", ADH) + "?"
				lyr_fl += tmp

				lyr_fl += "]|"

				fl += lyr_fl
			} else {

				NL := "2nd Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AH%d", i))
				LD3 := ""
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AJ%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AK%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AL%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AM%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BS%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BT%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)
			}

			//Layer3
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AN%d", i))
			if LD1 != "" {
				lyr_fl := "|["
				NL := "3rd Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AO%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AP%d", i))
				LD4_dbl, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AQ%d", i)), 64)
				LD4_dbl = math.Round(LD4_dbl*100) / 100
				LD4 := fmt.Sprintf("%.2f", LD4_dbl)

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AR%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AS%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AT%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BU%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BV%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)

				tmp = "?" + NL + "?"
				lyr_fl += tmp

				tmp = "?" + LD1 + "?"
				lyr_fl += tmp

				tmp = "?" + LD2 + "?"
				lyr_fl += tmp

				tmp = "?" + LD3 + "?"
				lyr_fl += tmp

				tmp = "?" + LD4 + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", WD) + "?"
				lyr_fl += tmp

				tmp = "?" + strconv.Itoa(RM) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", DIFF) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", LYR) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", ADH) + "?"
				lyr_fl += tmp

				lyr_fl += "]|"

				fl += lyr_fl

			} else {
				NL := "3rd Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AO%d", i))
				LD3 := ""
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AQ%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AR%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AS%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AT%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BU%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BV%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)
			}

			//Layer4
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AU%d", i))
			if LD1 != "" {
				lyr_fl := "|["
				NL := "4th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AV%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AW%d", i))
				LD4_dbl, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AX%d", i)), 64)
				LD4_dbl = math.Round(LD4_dbl*100) / 100
				LD4 := fmt.Sprintf("%.2f", LD4_dbl)

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AY%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AZ%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BA%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BW%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BX%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)

				tmp = "?" + NL + "?"
				lyr_fl += tmp

				tmp = "?" + LD1 + "?"
				lyr_fl += tmp

				tmp = "?" + LD2 + "?"
				lyr_fl += tmp

				tmp = "?" + LD3 + "?"
				lyr_fl += tmp

				tmp = "?" + LD4 + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", WD) + "?"
				lyr_fl += tmp

				tmp = "?" + strconv.Itoa(RM) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", DIFF) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", LYR) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", ADH) + "?"
				lyr_fl += tmp

				lyr_fl += "]|"

				fl += lyr_fl
			} else {
				NL := "4th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AV%d", i))
				LD3 := ""
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AX%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AY%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AZ%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BA%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BW%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BX%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)
			}

			//Layer5
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BB%d", i))
			if LD1 != "" {
				lyr_fl := "|["

				NL := "5th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BC%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BD%d", i))
				LD4_dbl, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BE%d", i)), 64)
				LD4_dbl = math.Round(LD4_dbl*100) / 100
				LD4 := fmt.Sprintf("%.2f", LD4_dbl)

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BF%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BG%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BH%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BY%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BZ%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)

				tmp = "?" + NL + "?"
				lyr_fl += tmp

				tmp = "?" + LD1 + "?"
				lyr_fl += tmp

				tmp = "?" + LD2 + "?"
				lyr_fl += tmp

				tmp = "?" + LD3 + "?"
				lyr_fl += tmp

				tmp = "?" + LD4 + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", WD) + "?"
				lyr_fl += tmp

				tmp = "?" + strconv.Itoa(RM) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", DIFF) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", LYR) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", ADH) + "?"
				lyr_fl += tmp

				lyr_fl += "]|"

				fl += lyr_fl
			} else {
				NL := "5th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BC%d", i))
				LD3 := ""
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BE%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BF%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BG%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BH%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BY%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BZ%d", i)), 64)
				ADH = math.Round(ADH*100) / 100
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)
			}

			//Layer6
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BI%d", i))
			if LD1 != "" {
				lyr_fl := "|["

				NL := "6th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BJ%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BK%d", i))
				LD4_dbl, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BL%d", i)), 64)
				LD4_dbl = math.Round(LD4_dbl*100) / 100
				LD4 := fmt.Sprintf("%.2f", LD4_dbl)

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BM%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BN%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BO%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("CA%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				tmp = "?" + NL + "?"
				lyr_fl += tmp

				tmp = "?" + LD1 + "?"
				lyr_fl += tmp

				tmp = "?" + LD2 + "?"
				lyr_fl += tmp

				tmp = "?" + LD3 + "?"
				lyr_fl += tmp

				tmp = "?" + LD4 + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", WD) + "?"
				lyr_fl += tmp

				tmp = "?" + strconv.Itoa(RM) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", DIFF) + "?"
				lyr_fl += tmp

				tmp = "?" + fmt.Sprintf("%f", LYR) + "?"
				lyr_fl += tmp

				lyr_fl += "]|"

				fl += lyr_fl

			} else {
				NL := "6th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BJ%d", i))
				LD3 := ""
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BL%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BM%d", i)), 64)
				WD = math.Round(WD*100) / 100
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BN%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BO%d", i)), 64)
				DIFF = math.Round(DIFF*100) / 100
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("CA%d", i)), 64)
				LYR = math.Round(LYR*100) / 100
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)
			}

			fl += "}"

			Array_R_NDL = append(Array_R_NDL, R_NDL)

			arr_fl = append(arr_fl, fl)

		} else {
			code = 1
		}
		i++
	}

	_ = os.Remove("./uploads/Read.xlsx")

	tools.CreateFile("./uploads/rd.txt")
	tools.WriteFile("./uploads/rd.txt", arr_fl)

	res.Status = http.StatusOK
	res.Message = "Masuk"
	res.Data = Array_R_NDL
	return res, nil
}

func Input_NDL(stat string) (Response, error) {
	var res Response
	var I_NDL str.Input_NDL
	if stat == "ok" {

		con := db.CreateCon()

		var path = "./uploads/rd.txt"

		var data2 = []string{}

		by := tools.ReadFile(path)
		by2 := byte(0)
		by = append(by, by2)

		var new string = ""
		var i int = 0
		for by[i] != 0 {
			var co int = 0
			new = ""

			if by[i] == 124 {
				co++
				i++
				for co < 2 {
					if by[i] == 124 {
						co++
						i++
						data2 = append(data2, new)
					} else {
						new += string(by[i])
						i++
					}
				}
			} else if by[i] == 125 {
				fmt.Println(data2[15])
				fmt.Println(data2[25])
				fmt.Println(data2[26])
				fmt.Println(len(data2))

				data2[15] = strings.Replace(data2[15], "[", "", -1)
				data2[15] = strings.Replace(data2[15], "]", "", -1)
				data2[15] = strings.Replace(data2[15], "?", "|", -1)
				fmt.Println(data2[15])

				ln := len(data2)

				k := 25

				for k < ln {
					data2[k] = strings.Replace(data2[k], "[", "", -1)
					data2[k] = strings.Replace(data2[k], "]", "", -1)
					data2[k] = strings.Replace(data2[k], "?", "|", -1)
					fmt.Println(data2[k])
					k++
				}

				if data2[8] != "" {
					var ws str.Ws_no

					date, _ := time.Parse("2-Jan-06", data2[0])
					date_sql := date.Format("2006-01-02")

					date2, _ := time.Parse("2-Jan-06", data2[1])
					date_sql2 := date2.Format("2006-01-02")

					date3, _ := time.Parse("2-Jan-06", data2[2])
					date_sql3 := date3.Format("2006-01-02")

					sqlStatement := "SELECT ws_no FROM ndl_table WHERE ws_no=?"

					_ = con.QueryRow(sqlStatement, data2[8]).Scan(&ws.Ws_no)

					if ws.Ws_no == "" {

						sqlStatement := "INSERT INTO ndl_table (ws_no,tambah_data_tanggal,customer_delivery_date,job_done,durasi,analyzer_version,order_status,cylinder_status,gol,cust,item_name,model,up,repeat_ndl,toleransi,order_ndl,w_s_order,width,lenght_ndl,gusset,prod_size,w,c_ndl,color,total) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

						stmt, err := con.Prepare(sqlStatement)

						if err != nil {
							return res, err
						}

						_, err = stmt.Exec(data2[8], date_sql, date_sql2, date_sql3, data2[3], data2[4], data2[5], data2[6], data2[7], data2[9], data2[10], data2[11], data2[12], data2[13], data2[14], data2[15], data2[16], data2[17], data2[18], data2[19], data2[20], data2[21], data2[22], data2[23], data2[24])

						//Rekap
						gl_rkp := ""
						if data2[7] == "S" {
							gl_rkp = "BBK"
						} else {
							gl_rkp = "LTU"
						}

						sqlStatement = "INSERT INTO rekap (ws_no,date_rekap,customer_name,item_name,order_rekap,ws_meter,plant,delivery_period,status_rekap,comment_note) values(?,?,?,?,?,?,?,?,?,?)"

						stmt, err = con.Prepare(sqlStatement)

						if err != nil {
							return res, err
						}

						ord := String_Separator_To_Int(data2[15])

						_, err = stmt.Exec(data2[8], date_sql, data2[9], data2[10], ord[1], data2[17], gl_rkp, "", 0, "")

						if err != nil {
							return res, err
						}

						for j := 25; j < ln; j++ {

							fl_lyr := String_Separator_To_String(data2[j])

							db_lyr := "layer" + string(fl_lyr[0][0])

							if fl_lyr[0][0] == '1' {

								sqlStatement = "INSERT INTO " + db_lyr + " (ws_no,nama_layer,layer_detail,width,rm,diff,lyr,ink,adh,meter,kg,diff_price) values(?,?,?,?,?,?,?,?,?,?,?,?)"

								stmt, err = con.Prepare(sqlStatement)

								if err != nil {
									return res, err
								}

								ld := "|" + fl_lyr[1] + "|" + "|" + fl_lyr[2] + "|" + "|" + fl_lyr[3] + "|" + "|" + fl_lyr[4] + "|"

								_, err = stmt.Exec(data2[8], fl_lyr[0], ld, fl_lyr[5], fl_lyr[6], fl_lyr[7], fl_lyr[8], fl_lyr[9], fl_lyr[10], 0.0, 0.0, 0.0)

							} else if fl_lyr[0][0] == '6' {

								sqlStatement = "INSERT INTO " + db_lyr + " (ws_no,nama_layer,layer_detail,width,rm,diff,lyr,meter,kg,diff_price) values(?,?,?,?,?,?,?,?,?,?)"

								stmt, err = con.Prepare(sqlStatement)

								if err != nil {
									return res, err
								}

								ld := "|" + fl_lyr[1] + "|" + "|" + fl_lyr[2] + "|" + "|" + fl_lyr[3] + "|" + "|" + fl_lyr[4] + "|"

								_, err = stmt.Exec(data2[8], fl_lyr[0], ld, fl_lyr[5], fl_lyr[6], fl_lyr[7], fl_lyr[8], 0.0, 0.0, 0.0)

							} else {

								sqlStatement = "INSERT INTO " + db_lyr + " (ws_no,nama_layer,layer_detail,width,rm,diff,lyr,adh,meter,kg,diff_price) values(?,?,?,?,?,?,?,?,?,?,?)"

								stmt, err = con.Prepare(sqlStatement)

								if err != nil {
									return res, err
								}

								ld := "|" + fl_lyr[1] + "|" + "|" + fl_lyr[2] + "|" + "|" + fl_lyr[3] + "|" + "|" + fl_lyr[4] + "|"

								_, err = stmt.Exec(data2[8], fl_lyr[0], ld, fl_lyr[5], fl_lyr[6], fl_lyr[7], fl_lyr[8], fl_lyr[9], 0.0, 0.0, 0.0)

							}
						}

						//template
						sqlStatement = "INSERT INTO template (ws_no,date_template,internal_instruction_number,customer_name,item_name,material,quantity,quantity_status,delivery_period,ld1,ld3,meter,kg,price_kg,lyr) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

						stmt, err = con.Prepare(sqlStatement)

						if err != nil {
							return res, err
						}

						dlr := ""

						cnt := 0
						ld1 := ""
						ld3 := ""
						kg := ""
						prc := ""
						lyr_str := ""
						for x := 1; x <= 6; x++ {
							var dl str.Detail_layer

							ly := "layer" + strconv.Itoa(x)

							sqlStatement := "SELECT layer_detail,width,lyr FROM " + ly + " WHERE ws_no=?"

							_ = con.QueryRow(sqlStatement, data2[8]).Scan(&dl.Detail_layer, &dl.Width, &dl.Lyr)

							if dl.Detail_layer != "" {
								lyr_str = lyr_str + "|" + strconv.Itoa(x) + "|"
								spt := String_Separator_To_String(dl.Detail_layer)

								if cnt < 1 {
									dlr += spt[0]
									cnt++
								} else {
									dlr = dlr + "/" + spt[0]
								}

								fix := "|" + spt[0] + "|"
								ld1 += fix
								tmp := dl.Width * 1000.0
								fx := ""

								if x == 1 {
									fx = "|('W) " + strconv.FormatFloat(tmp, 'f', -1, 64) + "|"
								} else {
									fx = "|" + strconv.FormatFloat(tmp, 'f', -1, 64) + " x " + spt[1] + "|"
								}

								lyr_d := "|" + strconv.FormatFloat(dl.Lyr, 'f', -1, 64) + "|"
								kg += lyr_d

								ld3 += fx

								prc += "|0|"

							}

						}

						_, err = stmt.Exec(data2[8], date_sql, 0, data2[9], data2[10], dlr, ord[1], "roll/pcs", "", ld1, ld3, data2[17], kg, prc, lyr_str)

						stmt.Close()
					}

				}

				data2 = []string{}
				i++
			} else {
				i++
			}

		}

	}

	_ = os.Remove("./uploads/rd.txt")

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = I_NDL

	return res, nil
}

func Read_NDL(page int) (Response, error) {
	var res Response
	var arr_Read_NDL []str.Read_NDL

	con := db.CreateCon()

	limit := 50

	offset := (limit * page) - limit

	fmt.Println("offset:", offset)

	sqlStatement := "SELECT ws_no,tambah_data_tanggal,customer_delivery_date,job_done,durasi,analyzer_version,order_status,cylinder_status,gol,cust,item_name,model,up,repeat_ndl,toleransi,order_ndl,w_s_order,width,lenght_ndl,gusset,prod_size,w,c_ndl,color,total FROM ndl_table ORDER BY co ASC LIMIT ? OFFSET ?"

	rows, err := con.Query(sqlStatement, limit, offset)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	var temp str.Order

	for rows.Next() {
		fmt.Println("masuk")
		var Rd_NDL str.Read_NDL
		err = rows.Scan(&Rd_NDL.Ws_no, &Rd_NDL.Tambah_data_tanggal, &Rd_NDL.Customer_delivery_date,
			&Rd_NDL.Job_done, &Rd_NDL.Durasi, &Rd_NDL.Analyzer_version, &Rd_NDL.Order_status,
			&Rd_NDL.Cylider_status, &Rd_NDL.Gol, &Rd_NDL.Cust, &Rd_NDL.Item_name, &Rd_NDL.Model, &Rd_NDL.Up,
			&Rd_NDL.Repeat_ndl, &Rd_NDL.Toleransi, &temp.Order, &Rd_NDL.W_s_order, &Rd_NDL.Width,
			&Rd_NDL.Lenght, &Rd_NDL.Gusset, &Rd_NDL.Prod_size, &Rd_NDL.W, &Rd_NDL.C, &Rd_NDL.Color,
			&Rd_NDL.Total_layer)
		if err != nil {
			return res, err
		}

		fmt.Println(Rd_NDL)

		date, _ := time.Parse("2006-01-02", Rd_NDL.Tambah_data_tanggal)
		date_sql := date.Format("02-01-2006")
		date2, _ := time.Parse("2006-01-02", Rd_NDL.Customer_delivery_date)
		date_sql2 := date2.Format("02-01-2006")
		date3, _ := time.Parse("2006-01-02", Rd_NDL.Job_done)
		date_sql3 := date3.Format("02-01-2006")

		Rd_NDL.Tambah_data_tanggal = date_sql
		Rd_NDL.Customer_delivery_date = date_sql2
		Rd_NDL.Job_done = date_sql3

		Rd_NDL.Order_ndl = String_Separator_To_Int(temp.Order)

		fmt.Println(Rd_NDL.Order_ndl)
		for i := 1; i <= 6; i++ {
			var lyr str.Layer

			ly := "layer" + strconv.Itoa(i)

			if i == 1 {

				sqlStatement := "SELECT nama_layer,layer_detail,width,rm,diff,lyr,ink,adh FROM " + ly + " WHERE ws_no=?"

				_ = con.QueryRow(sqlStatement, Rd_NDL.Ws_no).Scan(&lyr.Nama_layer, &lyr.Layer_datail, &lyr.Width_layer, &lyr.Rm_layer, &lyr.Diff_layer,
					&lyr.Lyr_layer, &Rd_NDL.Ink_layer, &lyr.Adh_layer)

				if lyr.Nama_layer != "" {

					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, lyr.Nama_layer)

					ld := String_Separator_To_String(lyr.Layer_datail)
					fmt.Println(string(ld[2][len(ld[2])-1]))

					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, ld[0])
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, ld[1])
					fmt.Println(ld[0], ld[1])
					ld[2] = string(ld[2][len(ld[2])-1])
					fmt.Println("miu:", ld[2])

					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, ld[2])
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, ld[3])

					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
					Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

				} else {
					nl := ""
					nm := strconv.Itoa(i) + "nd Layer"
					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
					Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)
				}

			} else if i == 6 {
				sqlStatement := "SELECT nama_layer,layer_detail,width,rm,diff,lyr FROM " + ly + " WHERE ws_no=?"

				_ = con.QueryRow(sqlStatement, Rd_NDL.Ws_no).Scan(&lyr.Nama_layer, &lyr.Layer_datail, &lyr.Width_layer, &lyr.Rm_layer, &lyr.Diff_layer,
					&lyr.Lyr_layer)

				if lyr.Nama_layer != "" {
					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, lyr.Nama_layer)

					ld := String_Separator_To_String(lyr.Layer_datail)

					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, ld[0])
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, ld[1])

					ld[2] = string(ld[2][len(ld[2])-1])
					fmt.Println(ld[2])

					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, ld[2])
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, ld[3])

					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
				} else {
					nl := ""
					nm := strconv.Itoa(i) + "th Layer"
					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
				}

			} else {
				sqlStatement := "SELECT nama_layer,layer_detail,width,rm,diff,lyr,adh FROM " + ly + " WHERE ws_no=?"

				_ = con.QueryRow(sqlStatement, Rd_NDL.Ws_no).Scan(&lyr.Nama_layer, &lyr.Layer_datail, &lyr.Width_layer, &lyr.Rm_layer, &lyr.Diff_layer,
					&lyr.Lyr_layer, &lyr.Adh_layer)

				fmt.Println("ly ", ly)
				fmt.Println("nama: ", lyr.Nama_layer)

				if lyr.Nama_layer != "" {

					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, lyr.Nama_layer)

					ld := String_Separator_To_String(lyr.Layer_datail)

					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, ld[0])
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, ld[1])

					ld[2] = string(ld[2][len(ld[2])-1])
					fmt.Println(ld[2])

					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, ld[2])
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, ld[3])

					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
					Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)
				} else {
					nl := ""
					if i == 2 {

						nm := strconv.Itoa(i) + "nd Layer"
						Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
						Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
						Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
						Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
						Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
						Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
						Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
						Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
						Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
						Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

					} else if i == 3 {

						nm := strconv.Itoa(i) + "rd Layer"
						Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
						Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
						Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
						Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
						Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
						Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
						Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
						Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
						Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
						Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

					} else {

						nm := strconv.Itoa(i) + "th Layer"
						Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
						Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
						Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
						Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
						Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
						Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
						Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
						Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
						Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
						Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

					}
				}
			}
		}
		arr_Read_NDL = append(arr_Read_NDL, Rd_NDL)
	}

	if arr_Read_NDL == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Read_NDL
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Read_NDL
	}

	return res, nil
}

func Update_NDL(WS_no string, tambah_data_tanggal string, customer_delivery_date string, job_done string,
	order_status string, cylinder_status string, gol string, cust string, item_name string,
	model string, up int, repeat_ndl int, toleransi int, order_ndl int, layer string, detail_layer string) (Response, error) {

	var res Response
	con := db.CreateCon()
	var dur str.Duration

	var WLC str.WLC

	sqlstatement2 := "SELECT width,lenght_ndl,color FROM ndl_table WHERE ws_no=?"

	_ = con.QueryRow(sqlstatement2, WS_no).Scan(&WLC.Width, &WLC.Lenght_ndl, &WLC.Color)

	date, _ := time.Parse("02-01-06", tambah_data_tanggal)
	date_sql := date.Format("2006-01-02")

	date2, _ := time.Parse("02-01-06", customer_delivery_date)
	date_sql2 := date2.Format("2006-01-02")

	date3, _ := time.Parse("02-01-06", job_done)
	date_sql3 := date3.Format("2006-01-02")

	sqlstatement := "SELECT datediff(" + "\"" + date_sql3 + "\" , " + "\"" + date_sql + "\")"

	_ = con.QueryRow(sqlstatement).Scan(&dur.Duration)

	order2 := order_ndl * ((100 + toleransi) / 100)

	order_full := "|" + strconv.Itoa(order_ndl) + "|" + "|" + strconv.Itoa(order2) + "|"

	w_s_order := WLC.Lenght_ndl * float64(order2) / float64(up)
	w_s_order = math.Round(w_s_order*100) / 100

	prod_size := WLC.Width * float64(up) * 2.0
	prod_size = math.Round(prod_size*100) / 100

	lyr_arr := String_Separator_To_String(layer)

	detail_layer_arr := String_Separator_To_String(detail_layer)

	co := 0
	df := 0.0

	var miu str.Miu

	sqlstatement = "SELECT layer_detail FROM layer1 WHERE ws_no = ?"

	_ = con.QueryRow(sqlstatement, WS_no).Scan(&miu.Miu)

	ld_arr := String_Separator_To_String(miu.Miu)
	miu.Miu = ld_arr[2]

	total_all := 0.0

	for i := 0; i < len(lyr_arr); i++ {

		nLYR := "layer" + string(lyr_arr[i][len(lyr_arr[i])-1])

		if i == 0 {

			ld2, _ := strconv.ParseFloat(detail_layer_arr[co+1], 64)
			ld2 = math.Round(ld2*100) / 100
			ld4, _ := strconv.ParseFloat(detail_layer_arr[co+2], 64)
			ld4 = math.Round(ld4*100) / 100
			WDTH, _ := strconv.ParseFloat(detail_layer_arr[co+3], 64)
			WDTH = math.Round(WDTH*100) / 100
			RM, _ := strconv.ParseFloat(detail_layer_arr[co+4], 64)
			RM = math.Round(RM*100) / 100
			WDTH_after, _ := strconv.ParseFloat(detail_layer_arr[co+3+5], 64)
			WDTH_after = math.Round(WDTH_after*100) / 100
			df = RM - WDTH

			lyr := ld2 * ld4 * WDTH * w_s_order / 1000
			lyr = math.Round(lyr*100) / 100

			ink := 0.5 * float64(WLC.Color) * WDTH * w_s_order / 1000
			ink = math.Round(ink*100) / 100

			adh := 3.5 * w_s_order * WDTH_after / 1000
			adh = math.Round(adh*100) / 100

			total_all = total_all + lyr + ink + adh

			ldet := "|" + detail_layer_arr[co] + "|" + "|" + detail_layer_arr[co+1] + "|" + "|" + miu.Miu + "|" + "|" + detail_layer_arr[co+2] + "|"

			sqlstatement = "UPDATE " + nLYR + " SET layer_detail=?,width=?,rm=?,diff=?,lyr=?,ink=?,adh=? WHERE ws_no=?"

			stmt, err := con.Prepare(sqlstatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(ldet, WDTH, int(RM), df, lyr, ink, adh, WS_no)

			if err != nil {
				return res, err
			}

		} else if i == 5 {

			ld2, _ := strconv.ParseFloat(detail_layer_arr[co+1], 64)
			ld2 = math.Round(ld2*100) / 100
			ld4, _ := strconv.ParseFloat(detail_layer_arr[co+2], 64)
			ld4 = math.Round(ld4*100) / 100
			WDTH, _ := strconv.ParseFloat(detail_layer_arr[co+3], 64)
			WDTH = math.Round(WDTH*100) / 100
			RM, _ := strconv.ParseFloat(detail_layer_arr[co+4], 64)
			RM = math.Round(RM*100) / 100

			df = RM - WDTH

			lyr := ld2 * ld4 * WDTH * w_s_order / 1000
			lyr = math.Round(lyr*100) / 100

			total_all = total_all + lyr

			ldet := "|" + detail_layer_arr[co] + "|" + "|" + detail_layer_arr[co+1] + "|" + "|" + miu.Miu + "|" + "|" + detail_layer_arr[co+2] + "|"

			sqlstatement = "UPDATE " + nLYR + " SET layer_detail=?,width=?,rm=?,diff=?,lyr=? WHERE ws_no=?"

			stmt, err := con.Prepare(sqlstatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(ldet, WDTH, int(RM), df, lyr, WS_no)

			if err != nil {
				return res, err
			}

		} else {

			ld2, _ := strconv.ParseFloat(detail_layer_arr[co+1], 64)
			ld2 = math.Round(ld2*100) / 100
			ld4, _ := strconv.ParseFloat(detail_layer_arr[co+2], 64)
			ld4 = math.Round(ld4*100) / 100
			WDTH, _ := strconv.ParseFloat(detail_layer_arr[co+3], 64)
			WDTH = math.Round(WDTH*100) / 100
			RM, _ := strconv.ParseFloat(detail_layer_arr[co+4], 64)
			RM = math.Round(RM*100) / 100
			WDTH_after, _ := strconv.ParseFloat(detail_layer_arr[co+3+5], 64)
			WDTH_after = math.Round(WDTH_after*100) / 100

			df = RM - WDTH

			lyr := ld2 * ld4 * WDTH * w_s_order / 1000
			lyr = math.Round(lyr*100) / 100

			adh := 3.5 * w_s_order * WDTH_after / 1000
			adh = math.Round(adh*100) / 100

			total_all = total_all + lyr + adh

			ldet := "|" + detail_layer_arr[co] + "|" + "|" + detail_layer_arr[co+1] + "|" + "|" + miu.Miu + "|" + "|" + detail_layer_arr[co+2] + "|"

			sqlstatement = "UPDATE " + nLYR + " SET layer_detail=?,width=?,rm=?,diff=?,lyr=?,adh=? WHERE ws_no=?"

			stmt, err := con.Prepare(sqlstatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(ldet, WDTH, int(RM), df, lyr, adh, WS_no)

			if err != nil {
				return res, err
			}

		}

		co = co + 5
	}

	total_all = math.Round(total_all*100) / 100

	sqlstatement = "UPDATE ndl_table SET tambah_data_tanggal=?,customer_delivery_date=?,job_done=?,durasi=?," +
		"order_status=?,cylinder_status=?,gol=?,cust=?,item_name=?,model=?,up=?,repeat_ndl=?," +
		"toleransi=?,order_ndl=?,w_s_order=?,prod_size=?,total=? WHERE ws_no=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(date_sql, date_sql2, date_sql3, dur.Duration, order_status, cylinder_status,
		gol, cust, item_name, model, up, repeat_ndl, toleransi, order_full, w_s_order, prod_size,
		total_all, WS_no)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	//Rekap
	gl_rkp := ""
	if gol == "S" {
		gl_rkp = "BBK"
	} else {
		gl_rkp = "LTU"
	}

	sqlstatement = "UPDATE rekap SET date_rekap=?,customer_name=?,item_name=?,order_rekap=?,ws_meter=?,plant=? WHERE ws_no=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(date_sql, cust, item_name, order2, w_s_order, gl_rkp, WS_no)

	if err != nil {
		return res, err
	}

	//template

	co = 0
	ld3 := ""
	ld1 := ""
	material := ""
	lyr := ""
	cnt := 0
	meter := ""
	kg := ""

	for i := 0; i < len(lyr_arr); i++ {

		condt := string(lyr_arr[i][len(lyr_arr[i])-1])

		condt_int, _ := strconv.Atoi(condt)
		lyr = lyr + "|" + string(lyr_arr[i][len(lyr_arr[i])-1]) + "|"

		ld1 = ld1 + "|" + detail_layer_arr[co] + "|"

		if condt_int == 1 {
			WDTH, _ := strconv.ParseFloat(detail_layer_arr[co+3], 64)
			WDTH = math.Round(WDTH*100) / 100
			WDTH = WDTH * 1000
			ld3 = ld3 + "|('W) " + strconv.FormatFloat(WDTH, 'f', -1, 64) + "|"
		} else {
			WDTH, _ := strconv.ParseFloat(detail_layer_arr[co+3], 64)
			WDTH = math.Round(WDTH*100) / 100
			WDTH = WDTH * 1000
			ld3 = ld3 + "|" + strconv.FormatFloat(WDTH, 'f', -1, 64) + " x " + detail_layer_arr[co+1] + "|"
		}

		if cnt < 1 {
			material += detail_layer_arr[co]
			cnt++
		} else {
			material = material + "/" + detail_layer_arr[co]
		}

		ld2, _ := strconv.ParseFloat(detail_layer_arr[co+1], 64)
		ld2 = math.Round(ld2*100) / 100
		ld4, _ := strconv.ParseFloat(detail_layer_arr[co+2], 64)
		ld4 = math.Round(ld4*100) / 100
		WDTH, _ := strconv.ParseFloat(detail_layer_arr[co+3], 64)
		WDTH = math.Round(WDTH*100) / 100

		lyr := ld2 * ld4 * WDTH * w_s_order / 1000
		lyr = math.Round(lyr*100) / 100

		meter = strconv.FormatFloat(w_s_order, 'f', -1, 64)
		kg = kg + "|" + strconv.FormatFloat(lyr, 'f', -1, 64) + "|"
		co = co + 5
	}

	sqlstatement = "UPDATE template SET date_template=?,customer_name=?,item_name=?,material=?,quantity=?,ld1=?,ld3=?,meter=?,kg=?,lyr=? WHERE ws_no=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(date_sql, cust, item_name, material, order2, ld1, ld3, meter, kg, lyr, WS_no)

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

func Page() (Response, error) {
	var res Response
	var arr str.Page_no

	con := db.CreateCon()

	sqlStatement := "SELECT COUNT(ws_no) FROM ndl_table"

	err := con.QueryRow(sqlStatement).Scan(&arr.Page)

	if err != nil {
		return res, err
	}

	sisa := arr.Page % 50
	page := 0

	if sisa > 0 {
		page = ((arr.Page - sisa) / 50) + 1
	} else {
		page = ((arr.Page - sisa) / 50)
	}

	arr.Page = page

	if arr.Page == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr
	}

	return res, nil
}

func Read_NDL_WSNO(ws_no string) (Response, error) {
	var res Response
	var arr_Read_NDL []str.Read_NDL

	con := db.CreateCon()

	sqlStatement := "SELECT ws_no,tambah_data_tanggal,customer_delivery_date,job_done,durasi,analyzer_version,order_status,cylinder_status,gol,cust,item_name,model,up,repeat_ndl,toleransi,order_ndl,w_s_order,width,lenght_ndl,gusset,prod_size,w,c_ndl,color,total FROM ndl_table WHERE ws_no=?"

	rows, err := con.Query(sqlStatement, ws_no)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	var temp str.Order

	for rows.Next() {
		var Rd_NDL str.Read_NDL
		err = rows.Scan(&Rd_NDL.Ws_no, &Rd_NDL.Tambah_data_tanggal, &Rd_NDL.Customer_delivery_date,
			&Rd_NDL.Job_done, &Rd_NDL.Durasi, &Rd_NDL.Analyzer_version, &Rd_NDL.Order_status,
			&Rd_NDL.Cylider_status, &Rd_NDL.Gol, &Rd_NDL.Cust, &Rd_NDL.Item_name, &Rd_NDL.Model, &Rd_NDL.Up,
			&Rd_NDL.Repeat_ndl, &Rd_NDL.Toleransi, &temp.Order, &Rd_NDL.W_s_order, &Rd_NDL.Width,
			&Rd_NDL.Lenght, &Rd_NDL.Gusset, &Rd_NDL.Prod_size, &Rd_NDL.W, &Rd_NDL.C, &Rd_NDL.Color,
			&Rd_NDL.Total_layer)
		if err != nil {
			return res, err
		}

		date, _ := time.Parse("2006-01-02", Rd_NDL.Tambah_data_tanggal)
		date_sql := date.Format("02-01-2006")
		date2, _ := time.Parse("2006-01-02", Rd_NDL.Customer_delivery_date)
		date_sql2 := date2.Format("02-01-2006")
		date3, _ := time.Parse("2006-01-02", Rd_NDL.Job_done)
		date_sql3 := date3.Format("02-01-2006")

		Rd_NDL.Tambah_data_tanggal = date_sql
		Rd_NDL.Customer_delivery_date = date_sql2
		Rd_NDL.Job_done = date_sql3

		Rd_NDL.Order_ndl = String_Separator_To_Int(temp.Order)

		for i := 1; i <= 6; i++ {
			var lyr str.Layer

			ly := "layer" + strconv.Itoa(i)

			if i == 1 {

				sqlStatement := "SELECT nama_layer,layer_detail,width,rm,diff,lyr,ink,adh FROM " + ly + " WHERE ws_no=?"

				_ = con.QueryRow(sqlStatement, Rd_NDL.Ws_no).Scan(&lyr.Nama_layer, &lyr.Layer_datail, &lyr.Width_layer, &lyr.Rm_layer, &lyr.Diff_layer,
					&lyr.Lyr_layer, &Rd_NDL.Ink_layer, &lyr.Adh_layer)

				if lyr.Nama_layer != "" {

					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, lyr.Nama_layer)

					ld := String_Separator_To_String(lyr.Layer_datail)

					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, ld[0])
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, ld[1])

					ld[2] = string(ld[2][len(ld[2])-1])
					fmt.Println(ld[2])

					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, ld[2])
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, ld[3])

					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
					Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

				} else {
					nl := ""
					nm := strconv.Itoa(i) + "nd Layer"
					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
					Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)
				}

			} else if i == 6 {
				sqlStatement := "SELECT nama_layer,layer_detail,width,rm,diff,lyr FROM " + ly + " WHERE ws_no=?"

				_ = con.QueryRow(sqlStatement, Rd_NDL.Ws_no).Scan(&lyr.Nama_layer, &lyr.Layer_datail, &lyr.Width_layer, &lyr.Rm_layer, &lyr.Diff_layer,
					&lyr.Lyr_layer)

				if lyr.Nama_layer != "" {
					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, lyr.Nama_layer)

					ld := String_Separator_To_String(lyr.Layer_datail)

					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, ld[0])
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, ld[1])

					ld[2] = string(ld[2][len(ld[2])-1])
					fmt.Println(ld[2])

					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, ld[2])
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, ld[3])

					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
				} else {
					nl := ""
					nm := strconv.Itoa(i) + "th Layer"
					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
				}

			} else {
				sqlStatement := "SELECT nama_layer,layer_detail,width,rm,diff,lyr,adh FROM " + ly + " WHERE ws_no=?"

				_ = con.QueryRow(sqlStatement, Rd_NDL.Ws_no).Scan(&lyr.Nama_layer, &lyr.Layer_datail, &lyr.Width_layer, &lyr.Rm_layer, &lyr.Diff_layer,
					&lyr.Lyr_layer, &lyr.Adh_layer)

				fmt.Println("ly ", ly)
				fmt.Println("nama: ", lyr.Nama_layer)

				if lyr.Nama_layer != "" {

					Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, lyr.Nama_layer)

					ld := String_Separator_To_String(lyr.Layer_datail)

					Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, ld[0])
					Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, ld[1])

					ld[2] = string(ld[2][len(ld[2])-1])
					fmt.Println(ld[2])

					Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, ld[2])
					Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, ld[3])

					Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
					Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
					Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
					Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
					Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)
				} else {
					nl := ""
					if i == 2 {

						nm := strconv.Itoa(i) + "nd Layer"
						Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
						Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
						Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
						Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
						Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
						Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
						Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
						Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
						Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
						Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

					} else if i == 3 {

						nm := strconv.Itoa(i) + "rd Layer"
						Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
						Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
						Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
						Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
						Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
						Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
						Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
						Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
						Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
						Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

					} else {

						nm := strconv.Itoa(i) + "th Layer"
						Rd_NDL.Nama_layer = append(Rd_NDL.Nama_layer, nm)
						Rd_NDL.Layer_datail_1 = append(Rd_NDL.Layer_datail_1, nl)
						Rd_NDL.Layer_datail_2 = append(Rd_NDL.Layer_datail_2, nl)
						Rd_NDL.Layer_datail_3 = append(Rd_NDL.Layer_datail_3, nl)
						Rd_NDL.Layer_datail_4 = append(Rd_NDL.Layer_datail_4, nl)
						Rd_NDL.Width_layer = append(Rd_NDL.Width_layer, lyr.Width_layer)
						Rd_NDL.Rm_layer = append(Rd_NDL.Rm_layer, lyr.Rm_layer)
						Rd_NDL.Diff_layer = append(Rd_NDL.Diff_layer, lyr.Diff_layer)
						Rd_NDL.Lyr_layer = append(Rd_NDL.Lyr_layer, lyr.Lyr_layer)
						Rd_NDL.Adh_layer = append(Rd_NDL.Adh_layer, lyr.Adh_layer)

					}
				}
			}
		}

		arr_Read_NDL = append(arr_Read_NDL, Rd_NDL)
	}

	if arr_Read_NDL == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Read_NDL
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Read_NDL
	}

	return res, nil
}
