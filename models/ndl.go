package models

import (
	"Backend-Project-NDL/db"
	str "Backend-Project-NDL/struct-all-ndl"
	"Backend-Project-NDL/tools"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
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
			tmp = "|" + fmt.Sprintf("%f", R_NDL.W_s_order) + "|"
			fl += tmp

			R_NDL.Width, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", i)), 64)
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Width) + "|"
			fl += tmp

			R_NDL.Lenght, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("T%d", i)), 64)
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Lenght) + "|"
			fl += tmp

			R_NDL.Gusset, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("U%d", i)), 64)
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Gusset) + "|"
			fl += tmp

			R_NDL.Prod_size, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("V%d", i)), 64)
			tmp = "|" + fmt.Sprintf("%f", R_NDL.Prod_size) + "|"
			fl += tmp

			R_NDL.W, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("W%d", i)), 64)
			tmp = "|" + fmt.Sprintf("%f", R_NDL.W) + "|"
			fl += tmp

			R_NDL.C, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("X%d", i)), 64)
			tmp = "|" + fmt.Sprintf("%f", R_NDL.C) + "|"
			fl += tmp

			R_NDL.Color, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Y%d", i)))
			tmp = "|" + strconv.Itoa(R_NDL.Color) + "|"
			fl += tmp

			R_NDL.Total_layer, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("CB%d", i)), 64)
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
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AC%d", i))

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
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)
				tmp = "?" + fmt.Sprintf("%f", WD) + "?"
				lyr_fl += tmp

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AE%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)
				tmp = "?" + strconv.Itoa(RM) + "?"
				lyr_fl += tmp

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AF%d", i)), 64)
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)
				tmp = "?" + fmt.Sprintf("%f", DIFF) + "?"
				lyr_fl += tmp

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BP%d", i)), 64)
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)
				tmp = "?" + fmt.Sprintf("%f", LYR) + "?"
				lyr_fl += tmp

				R_NDL.Ink_layer, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BQ%d", i)), 64)
				tmp = "?" + fmt.Sprintf("%f", R_NDL.Ink_layer) + "?"
				lyr_fl += tmp

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BR%d", i)), 64)
				R_NDL.Adh_layer = append(R_NDL.Adh_layer, ADH)
				tmp = "?" + fmt.Sprintf("%f", ADH) + "?"
				lyr_fl += tmp

				lyr_fl += "]|"

				fl += lyr_fl
			}

			//Layer2
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AG%d", i))
			if LD1 != "" {
				lyr_fl := "|["

				NL := "2nd Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AH%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AI%d", i))
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AJ%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AK%d", i)), 64)
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AL%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AM%d", i)), 64)
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BS%d", i)), 64)
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BT%d", i)), 64)
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
			}

			//Layer3
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AN%d", i))
			if LD1 != "" {
				lyr_fl := "|["
				NL := "3rd Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AO%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AP%d", i))
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AQ%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AR%d", i)), 64)
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AS%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AT%d", i)), 64)
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BU%d", i)), 64)
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BV%d", i)), 64)
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

			}

			//Layer4
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AU%d", i))
			if LD1 != "" {
				lyr_fl := "|["
				NL := "4th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AV%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AW%d", i))
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AX%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AY%d", i)), 64)
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AZ%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BA%d", i)), 64)
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BW%d", i)), 64)
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BX%d", i)), 64)
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
			}

			//Layer5
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BB%d", i))
			if LD1 != "" {
				lyr_fl := "|["

				NL := "5th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BC%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BD%d", i))
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BE%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BF%d", i)), 64)
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BG%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BH%d", i)), 64)
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BY%d", i)), 64)
				R_NDL.Lyr_layer = append(R_NDL.Lyr_layer, LYR)

				ADH, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BZ%d", i)), 64)
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
			}

			//Layer6
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BI%d", i))
			if LD1 != "" {
				lyr_fl := "|["

				NL := "6th Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BJ%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BK%d", i))
				LD4 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BL%d", i))

				R_NDL.Layer_datail_1 = append(R_NDL.Layer_datail_1, LD1)
				R_NDL.Layer_datail_2 = append(R_NDL.Layer_datail_2, LD2)
				R_NDL.Layer_datail_3 = append(R_NDL.Layer_datail_3, LD3)
				R_NDL.Layer_datail_4 = append(R_NDL.Layer_datail_4, LD4)

				WD, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BM%d", i)), 64)
				R_NDL.Width_layer = append(R_NDL.Width_layer, WD)

				RM, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BN%d", i)))
				R_NDL.Rm_layer = append(R_NDL.Rm_layer, RM)

				DIFF, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BO%d", i)), 64)
				R_NDL.Diff_layer = append(R_NDL.Diff_layer, DIFF)

				LYR, _ := strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("CA%d", i)), 64)
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

func Generate_Id_layer(L_cat string) int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT layer" + L_cat + " FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET layer" + L_cat + "=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
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

					date, _ := time.Parse("2-Jan-06", data2[0])
					date_sql := date.Format("2006-01-02")

					date2, _ := time.Parse("2-Jan-06", data2[1])
					date_sql2 := date2.Format("2006-01-02")

					date3, _ := time.Parse("2-Jan-06", data2[2])
					date_sql3 := date3.Format("2006-01-02")

					sqlStatement := "INSERT INTO ndl_table (ws_no,tambah_data_tanggal,customer_delivery_date,job_done,durasi,analyzer_version,order_status,cylinder_status,gol,cust,item_name,model,up,repeat_ndl,toleransi,order_ndl,w_s_order,width,lenght_ndl,gusset,prod_size,w,c_ndl,color,total) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					_, err = stmt.Exec(data2[8], date_sql, date_sql2, date_sql3, data2[3], data2[4], data2[5], data2[6], data2[7], data2[9], data2[10], data2[11], data2[12], data2[13], data2[14], data2[15], data2[16], data2[17], data2[18], data2[19], data2[20], data2[21], data2[22], data2[23], data2[24])

					for j := 25; j < ln; j++ {

						fl_lyr := String_Separator_To_String(data2[j])

						db_lyr := "layer" + string(fl_lyr[0][0])

						if fl_lyr[0][0] == '1' {

							sqlStatement = "INSERT INTO " + db_lyr + " (ws_no,nama_layer,layer_detail,width,rm,diff,lyr,ink,adh) values(?,?,?,?,?,?,?,?,?)"

							stmt, err = con.Prepare(sqlStatement)

							if err != nil {
								return res, err
							}

							ld := "|" + fl_lyr[1] + "|" + "|" + fl_lyr[2] + "|" + "|" + fl_lyr[3] + "|" + "|" + fl_lyr[4] + "|"

							_, err = stmt.Exec(data2[8], fl_lyr[0], ld, fl_lyr[5], fl_lyr[6], fl_lyr[7], fl_lyr[8], fl_lyr[9], fl_lyr[10])

						} else if fl_lyr[0][0] == '6' {

							sqlStatement = "INSERT INTO " + db_lyr + " (ws_no,nama_layer,layer_detail,width,rm,diff,lyr) values(?,?,?,?,?,?,?)"

							stmt, err = con.Prepare(sqlStatement)

							if err != nil {
								return res, err
							}

							ld := "|" + fl_lyr[1] + "|" + "|" + fl_lyr[2] + "|" + "|" + fl_lyr[3] + "|" + "|" + fl_lyr[4] + "|"

							_, err = stmt.Exec(data2[8], fl_lyr[0], ld, fl_lyr[5], fl_lyr[6], fl_lyr[7], fl_lyr[8])

						} else {

							sqlStatement = "INSERT INTO " + db_lyr + " (ws_no,nama_layer,layer_detail,width,rm,diff,lyr,adh) values(?,?,?,?,?,?,?,?)"

							stmt, err = con.Prepare(sqlStatement)

							if err != nil {
								return res, err
							}

							ld := "|" + fl_lyr[1] + "|" + "|" + fl_lyr[2] + "|" + "|" + fl_lyr[3] + "|" + "|" + fl_lyr[4] + "|"

							_, err = stmt.Exec(data2[8], fl_lyr[0], ld, fl_lyr[5], fl_lyr[6], fl_lyr[7], fl_lyr[8], fl_lyr[9])

						}
					}

					stmt.Close()

				}

				data2 = []string{}
				i++
			} else {
				i++
			}

		}

	}

	_ = os.Remove("./uploads/rd.xlsx")

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = I_NDL

	return res, nil
}

/*func Read_NDL() (Response, error) {
	var res Response
	var arr_Read_NDL []str.Read_NDL
	var Read_NDL str.Read_NDL

	con := db.CreateCon()

	sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,satuan_barang,harga_barang FROM stock ORDER BY co ASC"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Kode_stock, &invent.Nama_barang, &invent.Jumlah_barang, &invent.Satuan_barang, &invent.Harga_barang)
		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}*/

func Update_Stock(kode_inventory string, nama_barang string, jumlah_barang float64, harga_barang int, satuan_barang string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlstatement := "UPDATE stock SET nama_barang=?,jumlah_barang=?,harga_barang=?,satuan_barang=? WHERE kode_stock=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_barang, jumlah_barang, harga_barang, satuan_barang, kode_inventory)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

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
