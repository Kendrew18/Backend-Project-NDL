package models

import (
	"Backend-Project-NDL/db"
	str "Backend-Project-NDL/struct-all-ndl"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Input_PO(ws_no string, layer string, nama_po_supplier string, tanggal_po string, meter string, kg string,
	price string) (Response, error) {
	var res Response
	var wsno str.Ws_no
	var I_PO str.Input_PO_Supplier

	con := db.CreateCon()

	sqlStatement := "SELECT ws_no FROM PO-supplier WHERE ws_no=?"

	_ = con.QueryRow(sqlStatement, ws_no).Scan(&wsno.Ws_no)

	date, _ := time.Parse("02-01-06", tanggal_po)
	date_sql := date.Format("2006-01-02")

	if wsno.Ws_no == "" {

		var pos str.Outstanding

		sqlStatement := "INSERT INTO PO-supplier (ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding,lyr_tot_our) values(?,?,?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		lyr, _ := strconv.Atoi(string(layer[len(layer)-1]))
		mtr, _ := strconv.ParseFloat(meter, 64)
		kg_f, _ := strconv.ParseFloat(kg, 64)
		prc, _ := strconv.ParseInt(price, 10, 64)

		var lyr_tot_kg float64
		var lyr_tot_m float64
		var lyr_tot_price int64

		lyr_tot_kg += kg_f
		lyr_tot_m += mtr
		lyr_tot_price += prc

		sqlStatement = "SELECT lyr,price_kg,kg,meter FROM template WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&pos.Layer, &pos.Price_kg, &pos.Kg, &pos.Meter)

		kode := String_Separator_To_Int(pos.Layer)
		kg_temp := String_Separator_To_float64(pos.Kg)
		prc_temp := String_Separator_To_Int64(pos.Price_kg)

		pos_tot_kg := ""

		pos_ot_kg := ""

		for i := 0; i < len(kode); i++ {
			if kode[i] == lyr {

				u_layer := "layer" + strconv.Itoa(kode[i])

				pos_tot_kg = pos_tot_kg + "|" + strconv.FormatFloat(lyr_tot_kg, 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(lyr_tot_m, 'f', -1, 64) + "|" + "|" + strconv.FormatInt(lyr_tot_price, 10) + "|"

				sqlstatement := "UPDATE " + u_layer + " SET meter=?,kg=?,diff_price=? WHERE ws_no=?"

				stmt2, err := con.Prepare(sqlstatement)

				if err != nil {
					return res, err
				}

				_, err = stmt2.Exec(lyr_tot_m, lyr_tot_kg, lyr_tot_price)

				if err != nil {
					return res, err
				}

				stmt2.Close()

				temp_1 := lyr_tot_kg - kg_temp[i]
				temp_2 := lyr_tot_m - pos.Meter
				temp_3 := prc_temp[i] - lyr_tot_price

				pos_ot_kg = pos_ot_kg + "|" + strconv.FormatFloat(temp_1, 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(temp_2, 'f', -1, 64) + "|" + "|" + strconv.FormatInt(temp_3, 10) + "|"
			}
		}

		layer_fix := "|" + layer + "|"
		nama_po_supplier_fix := "|" + nama_po_supplier + "|"
		meter = "|" + meter + "|"
		kg = "|" + kg + "|"
		price = "|" + price + "|"

		_, err = stmt.Exec(ws_no, layer_fix, nama_po_supplier_fix, date_sql, meter, kg, price, pos_tot_kg, pos_ot_kg, layer_fix)

		stmt.Close()

		sqlStatement = "SELECT ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding FROM PO-supplier WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&I_PO.Ws_no, &I_PO.Layer, &I_PO.Nama_PO_Supplier,
			&I_PO.Tanggal_PO, &I_PO.Meter, &I_PO.Kg, &I_PO.Diff_Price, &I_PO.Total, &I_PO.Outstanding)

		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = I_PO

	} else {

		sqlStatement = "SELECT ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding,lyr_tot_out FROM PO-supplier WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&I_PO.Ws_no, &I_PO.Layer, &I_PO.Nama_PO_Supplier,
			&I_PO.Tanggal_PO, &I_PO.Meter, &I_PO.Kg, &I_PO.Diff_Price, &I_PO.Total, &I_PO.Outstanding, I_PO.Lyr_tot_out)

		var pos str.Outstanding

		lyr_sql_out := String_Separator_To_String(I_PO.Outstanding)

		lyr_sql_tot := String_Separator_To_String(I_PO.Total)

		sqlStatement = "SELECT lyr,price_kg,kg,meter FROM template WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&pos.Layer, &pos.Price_kg, &pos.Kg, &pos.Meter)

		x := 0

		kode := String_Separator_To_Int(pos.Layer)

		pos_tot_kg := ""

		pos_ot_kg := ""

		lyr_tot_out_sep := String_Separator_To_String(I_PO.Lyr_tot_out)

		code_kode := 0

		code_lyr_tot := 0

		for j := 0; j < len(kode); j++ {
			lyr, _ := strconv.Atoi(string(layer[len(layer)-1]))
			if kode[j] == lyr {
				code_kode = 1
			} else {
				code_kode = 0
			}
		}

		for i := 0; i < len(lyr_tot_out_sep); i++ {
			if lyr_tot_out_sep[i][len(lyr_tot_out_sep[i])-1] == layer[len(layer)-1] {
				code_lyr_tot = 1
			} else {
				code_lyr_tot = 0
			}

		}

		if code_kode == 1 && code_lyr_tot == 1 {

			for i := 0; i < len(lyr_tot_out_sep); i++ {

				if lyr_tot_out_sep[i][len(lyr_tot_out_sep[i])-1] == layer[len(layer)-1] {

					tmp1, _ := strconv.ParseFloat(lyr_sql_tot[x], 64)
					kg_f, _ := strconv.ParseFloat(kg, 64)
					lyr_tot_kg := tmp1 + kg_f

					tmp2, _ := strconv.ParseFloat(lyr_sql_tot[x+1], 64)
					meter_f, _ := strconv.ParseFloat(meter, 64)
					lyr_tot_m := meter_f + tmp2

					tmp3, _ := strconv.ParseInt(lyr_sql_tot[x+2], 10, 64)
					price_f, _ := strconv.ParseInt(price, 10, 64)
					lyr_tot_price := price_f + tmp3

					kg_out_f, _ := strconv.ParseFloat(lyr_sql_out[x], 64)
					lyr_out_kg := tmp1 + kg_out_f

					meter_out_f, _ := strconv.ParseFloat(lyr_sql_out[x+1], 64)
					lyr_out_m := meter_out_f + tmp2

					price_out_f, _ := strconv.ParseInt(lyr_sql_out[x+2], 10, 64)
					lyr_out_price := price_out_f + tmp3

					pos_tot_kg = pos_tot_kg + "|" + fmt.Sprintf("%.2f", lyr_tot_kg) + "|" + "|" + fmt.Sprintf("%.2f", lyr_tot_m) + "|" + "|" + strconv.FormatInt(lyr_tot_price, 10) + "|"

					pos_ot_kg = pos_ot_kg + "|" + fmt.Sprintf("%.2f", lyr_out_kg) + "|" + "|" + fmt.Sprintf("%.2f", lyr_out_m) + "|" + "|" + strconv.FormatInt(lyr_out_price, 10) + "|"

					u_layer := "layer" + strconv.Itoa(kode[i])

					sqlstatement := "UPDATE " + u_layer + " SET meter=?,kg=?,diff_price=? WHERE ws_no=?"

					stmt2, err := con.Prepare(sqlstatement)

					if err != nil {
						return res, err
					}

					_, err = stmt2.Exec(lyr_tot_m, lyr_tot_kg, lyr_tot_price)

					if err != nil {
						return res, err
					}

					stmt2.Close()
					code_lyr_tot = 1

				}

				x = x + 3

			}
		} else if code_kode == 1 && code_lyr_tot == 0 {

			mtr, _ := strconv.ParseFloat(meter, 64)
			kg_f, _ := strconv.ParseFloat(kg, 64)
			prc, _ := strconv.ParseInt(price, 10, 64)

			var lyr_tot_kg float64
			var lyr_tot_m float64
			var lyr_tot_price int64

			lyr_tot_kg += kg_f
			lyr_tot_m += mtr
			lyr_tot_price += prc

			kode := String_Separator_To_Int(pos.Layer)
			kg_temp := String_Separator_To_float64(pos.Kg)
			prc_temp := String_Separator_To_Int64(pos.Price_kg)
			lyr, _ := strconv.Atoi(string(layer[len(layer)-1]))

			for i := 0; i < len(kode); i++ {
				if kode[i] == lyr {

					I_PO.Lyr_tot_out = I_PO.Lyr_tot_out + "|" + layer + "|"

					u_layer := "layer" + strconv.Itoa(kode[i])

					pos_tot_kg = pos_tot_kg + "|" + strconv.FormatFloat(lyr_tot_kg, 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(lyr_tot_m, 'f', -1, 64) + "|" + "|" + strconv.FormatInt(lyr_tot_price, 10) + "|"

					sqlstatement := "UPDATE " + u_layer + " SET meter=?,kg=?,diff_price=? WHERE ws_no=?"

					stmt2, err := con.Prepare(sqlstatement)

					if err != nil {
						return res, err
					}

					_, err = stmt2.Exec(lyr_tot_m, lyr_tot_kg, lyr_tot_price)

					if err != nil {
						return res, err
					}

					stmt2.Close()

					temp_1 := lyr_tot_kg - kg_temp[i]
					temp_2 := lyr_tot_m - pos.Meter
					temp_3 := prc_temp[i] - lyr_tot_price

					pos_ot_kg = pos_ot_kg + "|" + strconv.FormatFloat(temp_1, 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(temp_2, 'f', -1, 64) + "|" + "|" + strconv.FormatInt(temp_3, 10) + "|"
				}
			}
		}

		I_PO.Layer = I_PO.Layer + "|" + layer + "|"
		I_PO.Meter = I_PO.Meter + "|" + meter + "|"
		I_PO.Kg = I_PO.Kg + "|" + kg + "|"
		I_PO.Diff_Price = I_PO.Diff_Price + "|" + price + "|"
		I_PO.Tanggal_PO = I_PO.Tanggal_PO + "|" + tanggal_po + "|"
		I_PO.Nama_PO_Supplier = I_PO.Nama_PO_Supplier + "|" + nama_po_supplier + "|"

		I_PO.Total += pos_tot_kg
		I_PO.Outstanding += pos_ot_kg

		sqlstatement := "UPDATE PO-supplier SET layer=?,nama_po_supplier=?,tanggal_PO=?,meter=?,kg=?,diff_price=?,total=?,outstanding=? WHERE kode_stock=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(I_PO.Layer, I_PO.Nama_PO_Supplier, I_PO.Tanggal_PO, I_PO.Meter, I_PO.Kg,
			I_PO.Diff_Price, I_PO.Total, I_PO.Outstanding, I_PO.Lyr_tot_out, I_PO.Ws_no)

		if err != nil {
			return res, err
		}

		stmt.Close()

		sqlStatement = "SELECT ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding FROM PO-supplier WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&I_PO.Ws_no, &I_PO.Layer, &I_PO.Nama_PO_Supplier,
			&I_PO.Tanggal_PO, &I_PO.Meter, &I_PO.Kg, &I_PO.Diff_Price, &I_PO.Total, &I_PO.Outstanding)

		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = I_PO

	}

	return res, nil
}

func Read_PO(ws_no string, lyr string) (Response, error) {
	var res Response
	var arr_str str.Read_PO_supplier_str
	var arr_Read_po str.Read_PO_supplier_fix
	var arr_Read_po_apnd []str.Read_PO_supplier_fix

	con := db.CreateCon()

	sqlStatement := "SELECT ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding,lyr_tot_out FROM PO-supplier WHERE ws_no=?"

	err := con.QueryRow(sqlStatement, ws_no).Scan(&arr_str.Ws_no, &arr_str.Layer, &arr_str.Nama_po_supplier,
		&arr_str.Tanggal_PO, &arr_str.Tanggal_PO, &arr_str.Meter, &arr_str.Kg, &arr_str.Diff_Price,
		&arr_str.Total, &arr_str.Outstanding, &arr_str.Lyr_tot_out)

	if err != nil {
		return res, err
	}

	ly := String_Separator_To_String(arr_str.Layer)
	nps := String_Separator_To_String(arr_str.Nama_po_supplier)
	tpo := String_Separator_To_String(arr_str.Tanggal_PO)
	mtr := String_Separator_To_float64(arr_str.Meter)
	kg := String_Separator_To_float64(arr_str.Kg)
	dp := String_Separator_To_Int64(arr_str.Diff_Price)
	tt := String_Separator_To_String(arr_str.Total)
	ot := String_Separator_To_String(arr_str.Outstanding)

	for i := 0; i < len(ly); i++ {

		if string(ly[i][len(ly[i])-1]) == lyr {

			arr_Read_po.Nama_po_supplier = append(arr_Read_po.Nama_po_supplier, nps[i])
			arr_Read_po.Tanggal_PO = append(arr_Read_po.Tanggal_PO, tpo[i])
			arr_Read_po.Meter = append(arr_Read_po.Meter, mtr[i])
			arr_Read_po.Kg = append(arr_Read_po.Kg, kg[i])
			arr_Read_po.Diff_Price = append(arr_Read_po.Diff_Price, dp[i])

		}

	}

	arr_Read_po.Ws_no = ws_no
	ln := String_Separator_To_String(arr_str.Lyr_tot_out)

	co := 0

	for i := 0; i < len(ln); i++ {
		if string(ln[i][len(ln[i])-1]) == lyr {
			arr_Read_po.Total_meter, _ = strconv.ParseFloat(tt[co], 64)
			arr_Read_po.Total_kg, _ = strconv.ParseFloat(tt[co+1], 64)
			arr_Read_po.Total_price, _ = strconv.ParseInt(tt[co+2], 10, 64)

			arr_Read_po.Outstanding_meter, _ = strconv.ParseFloat(ot[co], 64)
			arr_Read_po.Outstanding_kg, _ = strconv.ParseFloat(ot[co+1], 64)
			arr_Read_po.Outstanding_price, _ = strconv.ParseInt(ot[co+2], 10, 64)
		}
		co += 3
	}

	arr_Read_po_apnd = append(arr_Read_po_apnd, arr_Read_po)

	if arr_Read_po_apnd == nil {
		arr_Read_po_apnd = append(arr_Read_po_apnd, arr_Read_po)
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Read_po_apnd
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Read_po_apnd
	}

	return res, nil
}

func Lyr_PO(ws_no string) (Response, error) {
	var res Response
	var arr str.Layer_PO
	var arr_str str.Layer_PO_Str

	con := db.CreateCon()

	sqlStatement := "SELECT lyr FROM template WHERE ws_no=?"

	err := con.QueryRow(sqlStatement, ws_no).Scan(&arr_str.Lyr_PO)

	if err != nil {
		return res, err
	}

	arr.Lyr_PO = String_Separator_To_String(arr_str.Lyr_PO)

	if arr.Lyr_PO == nil {
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
