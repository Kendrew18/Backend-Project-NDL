package models

import (
	"Backend-Project-NDL/db"
	str "Backend-Project-NDL/struct-all-ndl"
	"net/http"
	"strconv"
)

func Read_Rekap(page int) (Response, error) {
	var res Response
	var arr_Read_Rekap []str.Read_rekap

	con := db.CreateCon()

	limit := 50

	offset := (limit * page) - limit

	sqlStatement := "SELECT ws_no,date_rekap,customer_name,item_name,order_rekap,ws_meter,plant,delivery_period,status_rekap,comment_note FROM rekap ORDER BY id_rekap ASC LIMIT ? OFFSET ?"

	rows, err := con.Query(sqlStatement, limit, offset)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var Rd_RKP str.Read_rekap
		err = rows.Scan(&Rd_RKP.Ws_no, &Rd_RKP.Date_rekap, &Rd_RKP.Customer_name, &Rd_RKP.Item_name,
			&Rd_RKP.Order_rekap, &Rd_RKP.Ws_meter, &Rd_RKP.Plant, &Rd_RKP.Delivery_period,
			&Rd_RKP.Status_rekap, &Rd_RKP.Comment_note)
		if err != nil {
			return res, err
		}

		for i := 1; i <= 6; i++ {
			var lyr str.Rekap_layer
			var wsno str.Ws_no

			ly := "layer" + strconv.Itoa(i)

			sqlStatement := "SELECT meter,kg,diff_price,ws_no FROM " + ly + " WHERE ws_no=?"

			_ = con.QueryRow(sqlStatement, Rd_RKP.Ws_no).Scan(&lyr.Meter, &lyr.KG, &lyr.Diff_price, &wsno.Ws_no)

			if wsno.Ws_no != "" {

				Rd_RKP.Meter = append(Rd_RKP.Meter, lyr.Meter)
				Rd_RKP.KG = append(Rd_RKP.KG, lyr.KG)
				Rd_RKP.Diff_price = append(Rd_RKP.Diff_price, lyr.Diff_price)

			} else {
				Rd_RKP.Meter = append(Rd_RKP.Meter, lyr.Meter)
				Rd_RKP.KG = append(Rd_RKP.KG, lyr.KG)
				Rd_RKP.Diff_price = append(Rd_RKP.Diff_price, lyr.Diff_price)
			}

		}

		arr_Read_Rekap = append(arr_Read_Rekap, Rd_RKP)
	}

	if arr_Read_Rekap == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Read_Rekap
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Read_Rekap
	}

	return res, nil
}

func Update_Status_Rekap(ws_no string, status_rekap int, delivery_period string, comment string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlstatement := "UPDATE rekap SET status_rekap=?, delivery_period=?, comment_note=? WHERE ws_no=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(status_rekap, delivery_period, comment, ws_no)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	sqlstatement = "UPDATE template SET delivery_period=? WHERE ws_no=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(delivery_period, ws_no)

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
