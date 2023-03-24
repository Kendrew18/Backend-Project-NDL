package models

import (
	"Backend-Project-NDL/db"
	str "Backend-Project-NDL/struct-all-ndl"
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

				sqlstatement := "UPDATE " + u_layer + " SET meter=?,kg=?,diff_price=? WHERE kode_stock=?"

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

				pos_tot_kg = pos_tot_kg + "|" + strconv.FormatFloat(lyr_tot_kg[co], 'f', -1, 64) + "|" + "|" + strconv.FormatFloat(lyr_tot_m[co], 'f', -1, 64) + "|" + "|" + strconv.FormatInt(lyr_tot_price[co], 10) + "|"

				sqlstatement := "UPDATE " + u_layer + " SET meter=?,kg=?,diff_price=? WHERE kode_stock=?"

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
}
