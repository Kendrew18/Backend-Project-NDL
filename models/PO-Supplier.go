package models

import (
	"Backend-Project-NDL/db"
	str "Backend-Project-NDL/struct-all-ndl"
	"fmt"
	"net/http"
	"strconv"
)

func Input_PO(ws_no string, layer string, nama_po_supplier string, tanggal_po string, meter string, kg string,
	price string) (Response, error) {
	var res Response
	var wsno str.Ws_no
	var I_PO str.Input_PO_Supplier

	con := db.CreateCon()

	sqlStatement := "SELECT ws_no FROM PO-supplier WHERE ws_no=?"

	_ = con.QueryRow(sqlStatement, ws_no).Scan(&wsno.Ws_no)

	if wsno.Ws_no == "" {

		var pos str.Outstanding

		sqlStatement := "INSERT INTO PO-supplier (ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding) values(?,?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		lyr := String_Separator_To_Int(layer)
		mtr := String_Separator_To_float64(meter)
		kg_f := String_Separator_To_float64(kg)
		prc := String_Separator_To_Int64(price)

		lyr_tot_kg := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		lyr_tot_m := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		lyr_tot_price := []int64{0, 0, 0, 0, 0, 0}

		for i := 0; i < len(lyr); i++ {
			x := lyr[i] - 1
			lyr_tot_kg[x] += kg_f[i]
			lyr_tot_m[x] += mtr[i]
			lyr_tot_price[x] += prc[i]
		}

		sqlStatement = "SELECT lyr,price_kg,kg,meter FROM template WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&pos.Layer, &pos.Price_kg, &pos.Kg, &pos.Meter)

		kode := String_Separator_To_Int(pos.Layer)
		kg_temp := String_Separator_To_float64(pos.Kg)
		prc_temp := String_Separator_To_Int64(pos.Price_kg)

		pos_tot_kg := ""

		pos_ot_kg := ""

		for i := 0; i < len(kode); i++ {
			co := kode[i] - 1
			if lyr_tot_kg[co] != 0.0 {

				u_layer := "layer" + strconv.Itoa(kode[i])

				pos_tot_kg = pos_tot_kg + "|" + strconv.FormatFloat(lyr_tot_kg[co], 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(lyr_tot_m[co], 'f', -1, 64) + "|" + "|" + strconv.FormatInt(lyr_tot_price[co], 10) + "|"

				sqlstatement := "UPDATE " + u_layer + " SET meter=?,kg=?,diff_price=? WHERE ws_no=?"

				stmt2, err := con.Prepare(sqlstatement)

				if err != nil {
					return res, err
				}

				_, err = stmt2.Exec(lyr_tot_m[co], lyr_tot_kg[co], lyr_tot_price[co])

				if err != nil {
					return res, err
				}

				stmt2.Close()

				temp_1 := lyr_tot_kg[co] - kg_temp[i]
				temp_2 := lyr_tot_m[co] - pos.Meter
				temp_3 := prc_temp[i] - lyr_tot_price[co]

				pos_ot_kg = pos_ot_kg + "|" + strconv.FormatFloat(temp_1, 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(temp_2, 'f', -1, 64) + "|" + "|" + strconv.FormatInt(temp_3, 10) + "|"
			}
		}

		_, err = stmt.Exec(ws_no, layer, nama_po_supplier, tanggal_po, meter, kg, price, pos_tot_kg, pos_ot_kg)

		stmt.Close()

		sqlStatement = "SELECT ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding FROM PO-supplier WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&I_PO.Ws_no, &I_PO.Layer, &I_PO.Nama_PO_Supplier,
			&I_PO.Tanggal_PO, &I_PO.Meter, &I_PO.Kg, &I_PO.Diff_Price, &I_PO.Total, &I_PO.Outstanding)

		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = I_PO

	} else {

		sqlStatement = "SELECT ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding FROM PO-supplier WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&I_PO.Ws_no, &I_PO.Layer, &I_PO.Nama_PO_Supplier,
			&I_PO.Tanggal_PO, &I_PO.Meter, &I_PO.Kg, &I_PO.Diff_Price, &I_PO.Total, &I_PO.Outstanding)

		var pos str.Outstanding

		lyr := String_Separator_To_Int(layer)
		mtr := String_Separator_To_float64(meter)
		kg_f := String_Separator_To_float64(kg)
		prc := String_Separator_To_Int64(price)

		lyr_tot_kg := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		lyr_tot_m := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		lyr_tot_price := []int64{0, 0, 0, 0, 0, 0}

		lyr_sql := String_Separator_To_Int(I_PO.Layer)
		lyr_sql_tot := String_Separator_To_String(I_PO.Total)

		I_PO.Layer += layer
		I_PO.Meter += meter
		I_PO.Kg += kg
		I_PO.Diff_Price += price
		I_PO.Tanggal_PO += tanggal_po
		I_PO.Nama_PO_Supplier += nama_po_supplier

		for i := 0; i < len(lyr); i++ {
			x := lyr[i] - 1
			lyr_tot_kg[x] += kg_f[i]
			lyr_tot_m[x] += mtr[i]
			lyr_tot_price[x] += prc[i]
		}

		x := 0

		for i := 0; i < len(lyr_sql); i++ {
			tmp1, _ := strconv.ParseFloat(lyr_sql_tot[x], 64)
			lyr_tot_kg[lyr_sql[i]-1] = lyr_tot_kg[lyr_sql[i]-1] + tmp1

			tmp2, _ := strconv.ParseFloat(lyr_sql_tot[x+1], 64)
			lyr_tot_m[lyr_sql[i]-1] = lyr_tot_m[lyr_sql[i]-1] + tmp2

			tmp3, _ := strconv.ParseInt(lyr_sql_tot[x+2], 10, 64)
			lyr_tot_price[lyr_sql[i]-1] = lyr_tot_price[lyr_sql[i]-1] + tmp3

			x = x + 3
		}

		sqlStatement = "SELECT lyr,price_kg,kg,meter FROM template WHERE ws_no=?"

		_ = con.QueryRow(sqlStatement, ws_no).Scan(&pos.Layer, &pos.Price_kg, &pos.Kg, &pos.Meter)

		kode := String_Separator_To_Int(pos.Layer)
		kg_temp := String_Separator_To_float64(pos.Kg)
		prc_temp := String_Separator_To_Int64(pos.Price_kg)

		pos_tot_kg := ""

		pos_ot_kg := ""

		for i := 0; i < len(kode); i++ {
			co := kode[i] - 1
			if lyr_tot_kg[co] != 0.0 {

				u_layer := "layer" + strconv.Itoa(kode[i])

				pos_tot_kg = pos_tot_kg + "|" + fmt.Sprintf("%.2f", lyr_tot_kg[co]) + "|" + "|" + fmt.Sprintf("%.2f", lyr_tot_m[co]) + "|" + "|" + strconv.FormatInt(lyr_tot_price[co], 10) + "|"

				sqlstatement := "UPDATE " + u_layer + " SET meter=?,kg=?,diff_price=? WHERE ws_no=?"

				stmt2, err := con.Prepare(sqlstatement)

				if err != nil {
					return res, err
				}

				_, err = stmt2.Exec(lyr_tot_m[co], lyr_tot_kg[co], lyr_tot_price[co])

				if err != nil {
					return res, err
				}

				stmt2.Close()

				pos_tot_kg = pos_tot_kg + "|" + strconv.FormatFloat(lyr_tot_kg[co], 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(lyr_tot_m[co], 'f', -1, 64) + "|" + "|" + strconv.FormatInt(lyr_tot_price[co], 10) + "|"

				temp_1 := lyr_tot_kg[co] - kg_temp[i]
				temp_2 := lyr_tot_m[co] - pos.Meter
				temp_3 := prc_temp[i] - lyr_tot_price[co]

				pos_ot_kg = pos_ot_kg + "|" + strconv.FormatFloat(temp_1, 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(temp_2, 'f', -1, 64) + "|" + "|" + strconv.FormatInt(temp_3, 10) + "|"
			}
		}

		sqlstatement := "UPDATE PO-supplier SET layer=?,nama_po_supplier=?,tanggal_PO=?,meter=?,kg=?,diff_price=?,total=?,outstanding=? WHERE kode_stock=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(I_PO.Layer, I_PO.Nama_PO_Supplier, I_PO.Tanggal_PO, I_PO.Meter, I_PO.Kg,
			I_PO.Diff_Price, pos_tot_kg, pos_ot_kg, I_PO.Ws_no)

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
} //(dibenerin)

func Read_PO(ws_no string, lyr int, lyr_fx string) (Response, error) {
	var res Response
	var arr_str str.Read_PO_supplier_str
	var arr_Read_po str.Read_PO_supplier_fix
	var arr_Read_po_apnd []str.Read_PO_supplier_fix

	con := db.CreateCon()

	sqlStatement := "SELECT ws_no,layer,nama_po_supplier,tanggal_PO,meter,kg,diff_price,total,outstanding FROM template WHERE ws_no=?"

	err := con.QueryRow(sqlStatement, ws_no).Scan(&arr_str.Ws_no, &arr_str.Layer, &arr_str.Nama_po_supplier,
		&arr_str.Tanggal_PO, &arr_str.Tanggal_PO, &arr_str.Meter, &arr_str.Kg, &arr_str.Diff_Price,
		&arr_str.Total, &arr_str.Outstanding)

	if err != nil {
		return res, err
	}

	ly := String_Separator_To_Int(arr_str.Layer)
	nps := String_Separator_To_String(arr_str.Nama_po_supplier)
	tpo := String_Separator_To_String(arr_str.Tanggal_PO)
	mtr := String_Separator_To_float64(arr_str.Meter)
	kg := String_Separator_To_float64(arr_str.Kg)
	dp := String_Separator_To_Int64(arr_str.Diff_Price)
	tt := String_Separator_To_String(arr_str.Total)
	ot := String_Separator_To_String(arr_str.Outstanding)

	for i := 0; i < len(ly); i++ {

		if ly[i] == lyr {

			arr_Read_po.Nama_po_supplier = append(arr_Read_po.Nama_po_supplier, nps[i])
			arr_Read_po.Tanggal_PO = append(arr_Read_po.Tanggal_PO, tpo[i])
			arr_Read_po.Meter = append(arr_Read_po.Meter, mtr[i])
			arr_Read_po.Kg = append(arr_Read_po.Kg, kg[i])
			arr_Read_po.Diff_Price = append(arr_Read_po.Diff_Price, dp[i])

		}

	}

	arr_Read_po.Ws_no = ws_no
	ln := String_Separator_To_Int(lyr_fx)

	co := 0

	for i := 0; i < len(ln); i++ {
		if ln[i] == lyr {
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
