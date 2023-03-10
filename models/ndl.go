package models

import (
	"Backend-Project-NDL/db"
	str "Backend-Project-NDL/struct-all-ndl"
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

	for code != 1 {
		var R_NDL str.Read_NDL
		WS_NO := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", i))
		if WS_NO != "" {

			vc := "A" + strconv.Itoa(i)
			hc := "C" + strconv.Itoa(i)

			style, _ := xlsx.NewStyle(`{"number_format":15}`)
			xlsx.SetCellStyle(sheet1Name, vc, hc, style)

			R_NDL.Tambah_data_tanggal = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i))
			R_NDL.Customer_delivery_date = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i))
			R_NDL.Job_done = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i))
			R_NDL.Durasi, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)))
			R_NDL.Analyzer_version = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i))
			R_NDL.Order_status = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", i))
			R_NDL.Cylider_status = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i))

			R_NDL.Gol = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", i))
			R_NDL.Ws_no = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", i))
			R_NDL.Cust = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", i))
			R_NDL.Item_name = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", i))
			R_NDL.Model = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", i))
			R_NDL.Up, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("M%d", i)))
			R_NDL.Repeat_ndl, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("N%d", i)))
			R_NDL.Toleransi, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("O%d", i)))
			ORDER1, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("P%d", i)))
			ORDER2, _ := strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Q%d", i)))
			R_NDL.Order_ndl = append(R_NDL.Order_ndl, ORDER1)
			R_NDL.Order_ndl = append(R_NDL.Order_ndl, ORDER2)
			R_NDL.W_s_order, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("R%d", i)), 64)
			R_NDL.Width, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", i)), 64)
			R_NDL.Lenght, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("T%d", i)), 64)
			R_NDL.Gusset, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("U%d", i)), 64)
			R_NDL.Prod_size, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("V%d", i)), 64)
			R_NDL.W, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("W%d", i)), 64)
			R_NDL.C, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("X%d", i)), 64)
			R_NDL.Color, _ = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Y%d", i)))
			R_NDL.Total_layer, _ = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("CB%d", i)), 64)

			//Layer1
			LD1 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Z%d", i))
			if LD1 != "" {
				NL := "1st Layer"
				R_NDL.Nama_layer = append(R_NDL.Nama_layer, NL)

				LD2 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AA%d", i))
				LD3 := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AB%d", i))
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

			}

			//Layer3
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AN%d", i))
			if LD1 != "" {
				NL := "1st Layer"
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

			}

			//Layer4
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AU%d", i))
			if LD1 != "" {
				NL := "1st Layer"
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

			}

			//Layer5
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BB%d", i))
			if LD1 != "" {
				NL := "1st Layer"
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

			}

			//Layer6
			LD1 = xlsx.GetCellValue(sheet1Name, fmt.Sprintf("BI%d", i))
			if LD1 != "" {
				NL := "1st Layer"
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

			}

			Array_R_NDL = append(Array_R_NDL, R_NDL)
		} else {
			code = 1
		}
		i++
	}

	_ = os.Remove("./uploads/Read.xlsx")

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

func Input_NDL(Ws_no string, Tambah_data_tanggal string, Customer_delivery_date string, Job_done string, Durasi int,
	Analyzer_version string, Order_status string, Cylider_status string, Gol string, Cust string, Item_name string, Model string,
	Up int, Repeat_ndl int, Toleransi int, Order_ndl string, W_s_order float64, Width float64, Lenght float64, Gusset float64,
	Prod_size float64, W float64, C float64, Color int, Layer_name string, layer_detail string,
	width_layer string, rm string, diff string, lyr string, ink float64, adh string, total string) (Response, error) {
	var res Response
	var I_NDL str.Input_NDL
	var ws str.Ws_no

	con := db.CreateCon()

	sqlStatement := "SELECT ws_no FROM ndl_table WHERE ws_no=?"

	_ = con.QueryRow(sqlStatement, Ws_no).Scan(&ws.Ws_no)

	if ws.Ws_no == "" {

		date, _ := time.Parse("2-Jan-06", Tambah_data_tanggal)
		date_sql := date.Format("2006-01-02")

		date2, _ := time.Parse("2-Jan-06", Customer_delivery_date)
		date_sql2 := date2.Format("2006-01-02")

		date3, _ := time.Parse("2-Jan-06", Job_done)
		date_sql3 := date3.Format("2006-01-02")

		sqlStatement := "INSERT INTO ndl_table (ws_no,tambah_data_tanggal,customer_delivery_date,job_done,durasi,analyzer_version,order_status,cylinder_status,gol,cust,item_name,model,up,repeat_ndl,toleransi,order_ndl,w_s_order,width,length_ndl,gusset,prod_size,w,c_ndl,color) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(Ws_no, date_sql, date_sql2, date_sql3, Durasi, Analyzer_version, Order_status, Cylider_status, Gol, Cust, Item_name, Model, Up, Repeat_ndl, Toleransi, Order_ndl, W_s_order, Width, Lenght, Gusset, Prod_size, W, C, Color)

		nl := String_Separator_To_String(Layer_name)
		ld := String_Separator_To_String(layer_detail)
		wl := String_Separator_To_float64(width_layer)
		rm_l := String_Separator_To_Int(rm)
		diff_l := String_Separator_To_float64(diff)
		lyr_l := String_Separator_To_float64(lyr)
		adh_l := String_Separator_To_float64(adh)
		total_l := String_Separator_To_float64(total)

		for i := 0; i < len(nl); i++ {

			str := "layer" + string(nl[i][0])

			if nl[i][0] == '1' {
				L_cat := string(nl[i][0])

				nm := Generate_Id_layer(L_cat)

				nm_str := strconv.Itoa(nm)

				id := str + "-" + nm_str

				sqlStatement = "INSERT INTO " + str + " (ws_no,id_layer,nama_layer,layer_detail,width,rm,diff,lyr,ink,adh,total) values(?,?,?,?,?,?,?,?,?,?,?)"

				stmt, err = con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(Ws_no, id, nl[i], ld[i], wl[i], rm_l[i], diff_l[i], lyr_l[i], ink, adh_l[i], total_l[i])

			} else {
				L_cat := string(nl[i][0])

				nm := Generate_Id_layer(L_cat)

				nm_str := strconv.Itoa(nm)

				id := str + "-" + nm_str

				sqlStatement = "INSERT INTO " + str + " (ws_no,id_layer,nama_layer,layer_detail,width,rm,diff,lyr,ink,adh,total) values(?,?,?,?,?,?,?,?,?,?,?)"

				stmt, err = con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(Ws_no, id, nl[i], ld[i], wl[i], rm_l[i], diff_l[i], lyr_l[i], ink, adh_l[i], total_l[i])

			}
		}

		stmt.Close()

	}

	sqlStatement = "SELECT ws_no,tambah_data_tanggal,customer_delivery_date,job_done,durasi,analyzer_version,order_status,cylinder_status,gol,cust,item_name,model,up,repeat_ndl,toleransi,order_ndl,w_s_order,width,length_ndl,gusset,prod_size,w,c_ndl,color FROM ndl_table WHERE kode_transaksi=?"

	_ = con.QueryRow(sqlStatement, Ws_no).Scan(&I_NDL.Ws_no, &I_NDL.Tambah_data_tanggal, &I_NDL.Customer_delivery_date,
		&I_NDL.Job_done, &I_NDL.Durasi, &I_NDL.Analyzer_version, &I_NDL.Order_status, &I_NDL.Cylider_status, &I_NDL.Gol,
		&I_NDL.Cust, &I_NDL.Item_name, &I_NDL.Model, &I_NDL.Up, &I_NDL.Repeat_ndl, &I_NDL.Toleransi,
		&I_NDL.Order_ndl, &I_NDL.W_s_order, &I_NDL.Width, &I_NDL.Lenght, &I_NDL.Gusset, &I_NDL.Prod_size,
		&I_NDL.W, &I_NDL.C, &I_NDL.Color)

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
