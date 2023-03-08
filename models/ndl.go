package models

import (
	"net/http"
	"project-NDL/db"
	STR "project-NDL/struct_all_ndl"
	"strconv"
	"time"
)

func Generate_Id_layer() int {
	var obj STR.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT layer FROM generate_id"
	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET layer=?"

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
	var invent STR.Input_NDL
	var ws STR.Ws_no

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

		for i := 0; i < len(nl); i++ {
			str := "layer" + string(nl[i][0])
			if nl[i][0] == '1' {

				nm := Generate_Id_layer()

				nm_str := strconv.Itoa(nm)

				id := str + "-" + nm_str

				sqlStatement = "INSERT INTO " + str + " (ws_no,id_layer,nama_layer,layer_detail,width,rm,diff,lyr,ink,adh,total) values(?,?,?,?,?,?,?,?,?,?,?)"

				stmt, err = con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(Ws_no, id, nl[i], ld[i], wl[i])
			} else {

			}
		}

		sqlStatement = "INSERT INTO ndl_table (ws_no,tambah_data_tanggal,customer_delivery_date,job_done,durasi,analyzer_version,order_status,cylinder_status,gol,cust,item_name,model,up,repeat_ndl,toleransi,order_ndl,w_s_order,width,length_ndl,gusset,prod_size,w,c_ndl,color) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

		stmt, err = con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(Ws_no, date_sql, date_sql2, date_sql3, Durasi, Analyzer_version, Order_status, Cylider_status, Gol, Cust, Item_name, Model, Up, Repeat_ndl, Toleransi, Order_ndl, W_s_order, width, Lenght, Gusset, Prod_size, W, C, Color)

		stmt.Close()

	}

	return res, nil
}

func Read_Stock() (Response, error) {
	var res Response
	var arr_invent []str.Read_Stock
	var invent str.Read_Stock

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
}

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
